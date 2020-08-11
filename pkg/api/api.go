package api

import (
	"encoding/json"
	"net/http"
)

type API struct{}

type SlackURLVerifyPayload struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type SlackURLVerifyResponse struct {
	Challenge string `json:"challenge"`
}

func New() *API {
	return &API{}
}

func (a *API) SlackHandler(w http.ResponseWriter, r *http.Request) {
	var sURLVerify SlackURLVerifyPayload
	if err := json.NewDecoder(r.Body).Decode(&sURLVerify); err != nil {
		http.Error(w, "error decoding json payload from slack verify webhook", http.StatusInternalServerError)
		return
	}
	// respond with JSON Challenge
	res := SlackURLVerifyResponse{sURLVerify.Challenge}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
