package main

import (
	_ "embed"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marvinjwendt/httb/assets"
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
	slog.Info("starting httb server", "address", cfg.Addr)

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
	e.Use(middleware.CORS())

	// Custom middleware
	e.Use(DelayMiddleware)

	e.GET("/openapi.yaml", func(c echo.Context) error {
		return c.String(200, assets.OpenAPISpec)
	})

	e.StaticFS("/docs", echo.MustSubFS(assets.Swagger, "swagger-ui"))

	api.RegisterHandlers(e, service)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		slog.Error("server stopped", "error", err)
	}
}
