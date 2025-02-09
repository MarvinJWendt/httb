package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
)

func (s Service) GetPing(w http.ResponseWriter, r *http.Request, _ api.GetPingParams) {
	sendFormattedResponse(w, r, "pong", "message")
}
