// +build integration

package api

import (
	"bytes"
	"encoding/json"
	"github.com/dviramontes/developerhappiness.app/internal/config"
	"github.com/dviramontes/developerhappiness.app/pkg/db"
	"github.com/dviramontes/developerhappiness.app/pkg/slack"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestRouter(api *API) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/webhook/slack", api.SlackHandler)
	return r
}

func Test_API(t *testing.T) {
	config := config.Read("../../config.yaml", nil)
	db, err := db.Connect("postgres://postgres:postgres@10.254.254.254:5432/happydev_test?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	testAPI := New(config, db)
	router := NewTestRouter(testAPI)

	t.Run("handle: url_verification", func(t *testing.T) {
		var n slack.Notification
		challenge := "challenge-token"
		testNotification := &slack.Notification{
			Token:     "123",
			Challenge: challenge,
			Type:      "url_verification",
		}
		jsonPayload, err := json.Marshal(testNotification)
		if err != nil {
			t.Fatal(err)
		}
		res, err := http.NewRequest(
			"POST",
			"/api/webhook/slack",
			bytes.NewBuffer(jsonPayload),
		)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.NewDecoder(res.Body).Decode(&n); err != nil {
			t.Error(err)
		}
		if n.Challenge != challenge {
			t.Error("challenge token does not match!")
		}
	})

	t.Run("handle: team_join", func(t *testing.T) {
		var n slack.Notification
		testNotification := &slack.Notification{
			Event: slack.Event{
				Type: "team_join",
				User: slack.User{
					ID:      "xyz",
					Name:    "david",
					Deleted: false,
					Tz:      "America/Denver",
					Profile: slack.Profile{
						Email:   "me@mail.com",
						Image32: "https://avatars.slack-edge.com/profile.png",
					},
					IsAdmin: false,
					IsOwner: false,
					IsBot:   false,
				},
			},
			Type: "event_callback",
		}
		jsonPayload, err := json.Marshal(testNotification)
		if err != nil {
			t.Fatal(err)
		}
		res, err := http.NewRequest(
			"POST",
			"/api/webhook/slack",
			bytes.NewBuffer(jsonPayload),
		)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.NewDecoder(res.Body).Decode(&n); err != nil {
			t.Error(err)
		}
	})

	t.Cleanup(func() {
		ts := httptest.NewServer(router)
		defer ts.Close()
	})
}
