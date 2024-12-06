package random

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"log/slog"
)

// Contact generates a random contact.
func Contact() api.Contact {
	var contact api.Contact
	if err := gofakeit.Struct(&contact); err != nil {
		slog.Error("failed to generate random contact", "error", err)
	}

	return contact
}
