package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
)

func (s Service) DeleteStatusStatus(w http.ResponseWriter, r *http.Request, status int, params api.DeleteStatusStatusParams) {
	sendStatus(w, status)
}

func (s Service) GetStatusStatus(w http.ResponseWriter, r *http.Request, status int, params api.GetStatusStatusParams) {
	sendStatus(w, status)
}

func (s Service) PatchStatusStatus(w http.ResponseWriter, r *http.Request, status int, params api.PatchStatusStatusParams) {
	sendStatus(w, status)
}

func (s Service) PostStatusStatus(w http.ResponseWriter, r *http.Request, status int, params api.PostStatusStatusParams) {
	sendStatus(w, status)
}

func (s Service) PutStatusStatus(w http.ResponseWriter, r *http.Request, status int, params api.PutStatusStatusParams) {
	sendStatus(w, status)
}

func sendStatus(w http.ResponseWriter, status int) {
	sendJSON(w, status, map[string]any{
		"status":  status,
		"message": http.StatusText(status),
	})
}
