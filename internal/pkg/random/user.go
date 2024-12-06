package random

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"log/slog"
)

// User generates a random user.
func User() api.User {
	var user api.User
	if err := gofakeit.Struct(&user); err != nil {
		slog.Error("failed to generate random user", "error", err)
	}

	return user
}
