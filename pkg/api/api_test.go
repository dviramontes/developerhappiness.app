package api

import (
	"bytes"
	"encoding/json"
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
	testAPI := New()
	router := NewTestRouter(testAPI)

	t.Run("SlackHandler", func(t *testing.T) {
		var sURLVerify SlackURLVerifyPayload
		challenge := "super-challenging-token"
		testEvent := &SlackURLVerifyPayload{
			Token:     "123",
			Challenge: challenge,
			Type:      "event_callback",
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

		if err := json.NewDecoder(res.Body).Decode(&sURLVerify); err != nil {
			t.Error(err)
		}

		if sURLVerify.Challenge != challenge {
			t.Error("challenge token does not match!")
		}
	})

	t.Cleanup(func() {
		ts := httptest.NewServer(router)
		defer ts.Close()
	})
}
