package main

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/marvinjwendt/httb/internal/pkg/service"

	"github.com/marvinjwendt/httb/internal/pkg/config"
	_ "go.uber.org/automaxprocs"
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
	svc, err := service.NewService(cfg)
	if err != nil {
		slog.Error("failed to init service", "error", err)
		os.Exit(1)
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start service in a goroutine
	go func() {
		if err := svc.Start(ctx); err != nil {
			slog.Error("failed to start service", "error", err)
			cancel()
		}
	}()

	// Wait for shutdown signal
	<-sigChan
	slog.Info("received shutdown signal, starting graceful shutdown")

	// Give the server time to shutdown gracefully based on config
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer shutdownCancel()

	if err := svc.Shutdown(shutdownCtx); err != nil {
		slog.Error("failed to shutdown service gracefully", "error", err)
		os.Exit(1)
	}

	slog.Info("service shutdown completed")
}
