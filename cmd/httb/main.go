package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"github.com/marvinjwendt/httb/internal/pkg/config"
	slogecho "github.com/samber/slog-echo"
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

	service := Service{
		Greetings: []string{
			"Hello, World!",
			"Hello there!",
			"Hi!",
			"Hey!",
			"What's up?",
			"Hello, friend!",
		},
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Echo middlewares
	e.Use(slogecho.New(cfg.Logger))
	e.Use(middleware.Recover())

	api.RegisterHandlers(e, service)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		slog.Error("server stopped", "error", err)
	}
}
