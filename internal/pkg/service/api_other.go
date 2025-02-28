package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
)

func (s Service) GetPing(w http.ResponseWriter, r *http.Request, params api.GetPingParams) {
	if ok := s.Validate(w, &params); !ok {
		return
	}

	sendFormattedResponse(w, r, "pong", "message")
}
