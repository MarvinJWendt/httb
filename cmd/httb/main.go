package main

import (
	_ "embed"
	"github.com/marvinjwendt/httb/internal/pkg/service"
	"log/slog"

	"github.com/marvinjwendt/httb/internal/pkg/config"
)

var cfg *config.Config

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
	if err := service.NewService(cfg).Start(); err != nil {
		slog.Error("failed to start service", "error", err)
	}
}
