package api

import (
	"encoding/json"
	"fmt"
	"github.com/dviramontes/developerhappiness.app/pkg/db"
	"github.com/dviramontes/developerhappiness.app/pkg/slack"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type API struct {
	config *viper.Viper
	db     *db.DB
}

func New(c *viper.Viper, db *db.DB) *API {
	return &API{
		config: c,
		db:     db,
	}
}

func (a *API) SlackHandler(w http.ResponseWriter, r *http.Request) {
	var e slack.Event

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		log.Printf("failed to decode json payload from slack event webhook, err: %v", err)
		http.Error(w, "error decoding json payload from slack event webhook", http.StatusInternalServerError)
		return
	}

	res, err := a.Route(&e)
	if err != nil {
		log.Printf("failed to process event from slack API, err: %v", err)
		http.Error(w, "error processing event from slack API", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (a *API) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.db.GetUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to revtrieve user list, err: %v", err), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (a *API) Route(event *slack.Event) (slack.Response, error) {
	if event.Type == "url_verification" {
		// respond with JSON challenge token
		return slack.Response{Challenge: event.Challenge}, nil
	}

	if event.Type == "event_callback" {
		incoming := event.Event.User
		active := !incoming.Deleted

		switch event.Event.Type {
		case "user_change":
			var u db.User
			if err := a.db.Conn.
				Where("slack_id LIKE ?", fmt.Sprintf("%%%s%%", incoming.ID)).
				First(&u).Error; err != nil {
				return slack.Response{}, err
			}

			u.Name = incoming.Name
			u.Active = active
			u.IsBot = incoming.IsBot
			u.Email = incoming.Profile.Email
			u.Timezone = incoming.Tz
			u.ImgUrl = incoming.Profile.Image32
			u.IsAdmin = incoming.IsAdmin
			u.IsOwner = incoming.IsOwner

			a.db.Conn.Save(&u)

			return slack.Response{}, nil

		case "team_join":
			newUser := db.User{
				Name:     incoming.Name,
				Active:   active,
				IsBot:    incoming.IsBot,
				Email:    incoming.Profile.Email,
				Timezone: incoming.Tz,
				ImgUrl:   incoming.Profile.Image32,
				IsAdmin:  incoming.IsAdmin,
				IsOwner:  incoming.IsOwner,
				SlackId:  incoming.ID,
				IsNew:    true,
			}

			if err := a.db.CreateUser(&newUser); err != nil {
				return slack.Response{}, err
			}

			return slack.Response{}, nil
		case "member_joined_channel":
			// TODO: handle invitations ?
			fallthrough
		default:
			return slack.Response{}, nil
		}
	}

	return slack.Response{}, nil
}
