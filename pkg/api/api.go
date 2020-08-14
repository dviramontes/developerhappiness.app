package api

import (
	"encoding/json"
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

	a.Route(&e)

	// respond with JSON challenge token
	res := slack.Verify{Challenge: e.Challenge}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (a *API) Route(event *slack.Event) {
	switch event.Type {
	case "url_verification":
		// handle
	case "member_joined_channel":
		// handle
	case "team_join":
		// handle
	case "user_change":
		// handle
	}
}
