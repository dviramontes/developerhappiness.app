package api

import (
	"encoding/json"
	"fmt"
	"github.com/dviramontes/developerhappiness.app/pkg/db"
	"github.com/dviramontes/developerhappiness.app/pkg/slack"
	"github.com/spf13/viper"
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
		http.Error(w, "error decoding json payload from slack e webhook", http.StatusInternalServerError)
		return
	}

	res := a.Route(&e)

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

func (a *API) Route(event *slack.Event) slack.Response {
	switch event.Type {
	case "url_verification":
		// respond with JSON challenge token
		return slack.Response{Challenge: event.Challenge}
	case "member_joined_channel":
		// TODO: handle
	case "team_join":
		// TODO: handle
	case "user_change":
		// TODO: handle
	default:
		return slack.Response{}
	}
	return slack.Response{}
}
