package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"github.com/marvinjwendt/httb/internal/pkg/random"
	"net/http"
)

func (s Service) GetStreamJsonUser(w http.ResponseWriter, r *http.Request, params api.GetStreamJsonUserParams) {
	s.streamJSON(w, r, params.Interval, func() (any, error) {
		return random.User(), nil
	})
}

func (s Service) GetStreamJsonLogs(w http.ResponseWriter, r *http.Request, params api.GetStreamJsonLogsParams) {
	s.streamJSON(w, r, params.Interval, func() (any, error) {
		return random.NewLog(1, nil)[0], nil // TODO: add log levels
	})
}
