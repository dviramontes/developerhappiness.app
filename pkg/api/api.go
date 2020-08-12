package api

import (
	"encoding/json"
	"github.com/spf13/viper"
	"net/http"
)

type API struct {
	config *viper.Viper
}

type SlackURLVerifyPayload struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type SlackURLVerifyResponse struct {
	Challenge string `json:"challenge"`
}

func New(c *viper.Viper) *API {
	return &API{
		config: c,
	}
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
