package service

import (
	"log/slog"
	"net/http"
)

func sendError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}

func sendAndLogError(w http.ResponseWriter, status int, msg string) {
	slog.Error(msg)
	http.Error(w, msg, status)
}
