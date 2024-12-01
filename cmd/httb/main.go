package main

import (
	"github.com/marvinjwendt/httb/internal/pkg/config"
	"log/slog"
)

var (
	cfg *config.Config
)

func init() {
	// Init config
	env := config.ReadEnv()
	cfg = config.New(env)

	// Init logger
	slog.SetDefault(cfg.Logger)

	// Print config in debug mode
	slog.Debug("configuration", "environment", env)
}

func main() {
	slog.Info("starting httb server", "address", cfg.Listen)
}
