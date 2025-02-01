package main

import (
	_ "embed"
	"fmt"
	"github.com/marvinjwendt/httb/internal/pkg/service"
	"log/slog"
	"os"

	"github.com/marvinjwendt/httb/internal/pkg/config"
)

var cfg *config.Config

func init() {
	// Init config
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Init logger
	logger, err := cfg.SetupLogger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to setup logger: %v\n", err)
		os.Exit(1)
	}
	slog.SetDefault(logger)

	// Print config in debug mode
	slog.Debug("configuration", "environment", cfg)
}

func main() {
	if err := service.NewService(cfg).Start(); err != nil {
		slog.Error("failed to start service", "error", err)
	}
}
