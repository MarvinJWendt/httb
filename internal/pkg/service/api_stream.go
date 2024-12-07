package service

import (
	"github.com/labstack/echo/v4"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"github.com/marvinjwendt/httb/internal/pkg/random"
)

func (s Service) GetStreamJsonUser(ctx echo.Context, params api.GetStreamJsonUserParams) error {
	return s.streamJSON(ctx, params.Interval, func() (any, error) {
		return random.User(), nil
	})
}

func (s Service) GetStreamJsonLogs(ctx echo.Context, params api.GetStreamJsonLogsParams) error {
	return s.streamJSON(ctx, params.Interval, func() (any, error) {
		return random.NewLog(1, nil)[0], nil // TODO: add log levels
	})
}
