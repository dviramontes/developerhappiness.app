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

	t.Run("slack event handle: url_verification", func(t *testing.T) {
		var e slack.Event
		challenge := "challenge-token"
		testEvent := &slack.Event{
			Token:     "123",
			Challenge: challenge,
			Type:      "url_verification",
		}
		jsonPayload, err := json.Marshal(testEvent)
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
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			t.Error(err)
		}
		if e.Challenge != challenge {
			t.Error("challenge token does not match!")
		}
	})

	t.Cleanup(func() {
		ts := httptest.NewServer(router)
		defer ts.Close()
	})
}
