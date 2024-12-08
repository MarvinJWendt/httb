package service

import (
	"github.com/labstack/echo/v4"
	"github.com/marvinjwendt/httb/internal/pkg/api"
)

func (s Service) GetPing(ctx echo.Context, _ api.GetPingParams) error {
	return formatResponse(ctx, "pong", "message")
}
