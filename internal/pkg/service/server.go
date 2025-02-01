package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marvinjwendt/httb/assets"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	slogecho "github.com/samber/slog-echo"
	"log/slog"
)

func (s Service) Start() error {
	slog.Info("starting httb service")

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Echo middlewares
	e.Use(slogecho.New(slog.Default()))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Custom middleware
	e.Use(DelayMiddleware)

	e.GET("/openapi.yaml", func(c echo.Context) error {
		return c.String(200, assets.OpenAPISpec)
	})

	e.StaticFS("/docs", echo.MustSubFS(assets.Swagger, "swagger-ui"))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(200, assets.LandingPage)
	})

	api.RegisterHandlers(e, s)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		slog.Error("server stopped", "error", err)
		return err
	}

	return nil
}
