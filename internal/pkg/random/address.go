package random

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"log/slog"
)

// Address generates a random address.
func Address() api.Address {
	var address api.Address
	if err := gofakeit.Struct(&address); err != nil {
		slog.Error("failed to generate random address", "error", err)
	}

	return address
}
