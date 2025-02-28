package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
)

func (s Service) DeleteStatusCode(w http.ResponseWriter, _ *http.Request, code int, _ api.DeleteStatusCodeParams) {
	sendStatus(w, code)
}

func (s Service) GetStatusCode(w http.ResponseWriter, _ *http.Request, code int, _ api.GetStatusCodeParams) {
	sendStatus(w, code)
}

func (s Service) PatchStatusCode(w http.ResponseWriter, _ *http.Request, code int, _ api.PatchStatusCodeParams) {
	sendStatus(w, code)
}

func (s Service) PostStatusCode(w http.ResponseWriter, _ *http.Request, code int, _ api.PostStatusCodeParams) {
	sendStatus(w, code)
}

func (s Service) PutStatusCode(w http.ResponseWriter, _ *http.Request, code int, _ api.PutStatusCodeParams) {
	sendStatus(w, code)
}

func sendStatus(w http.ResponseWriter, code int) {
	sendJSON(w, code, map[string]any{
		"status":  code,
		"message": http.StatusText(code),
	})
}
