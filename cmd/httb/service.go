package main

import (
	"github.com/labstack/echo/v4"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"math/rand"
	"net/http"
)

type Format int

const (
	FormatInvalid Format = iota
	FormatJSON
	FormatText
)

// getFormat decides which format to use for the response.
// It first checks the query parameter `format` and then the `Accept` header.
// If no format is specified, it defaults to JSON.
func getFormat(r *http.Request) Format {
	// Read format from query parameter
	formatParam := r.URL.Query().Get("format")
	switch formatParam {
	case "json":
		return FormatJSON
	case "text":
		return FormatText
	}

	// Read format from request accept header
	acceptHeader := r.Header.Get("Accept")
	switch acceptHeader {
	case "application/json":
		return FormatJSON
	case "text/plain":
		return FormatText
	}

	return FormatJSON
}

func formatResponse(ctx echo.Context, r *http.Request, text, keyName string) error {
	format := getFormat(r)

	switch format {
	case FormatJSON:
		return ctx.JSON(http.StatusOK, map[string]string{keyName: text})
	case FormatText:
		return ctx.String(http.StatusOK, text)
	default:
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid format"})
	}
}

// Verify that Service fits the api.ServerInterface interface
var _ api.ServerInterface = (*Service)(nil)

type Service struct {
	Greetings []string
}

func (s Service) GetJsonRandom(ctx echo.Context, params api.GetJsonRandomParams) error {
	// TODO implement me
	panic("implement me")
}

func (s Service) GetGreeting(ctx echo.Context, _ api.GetGreetingParams) error {
	randomGreeting := s.Greetings[rand.Intn(len(s.Greetings))] //nolint:gosec // not security sensitive
	return formatResponse(ctx, ctx.Request(), randomGreeting, "greeting")
}

func (s Service) GetPing(ctx echo.Context, _ api.GetPingParams) error {
	return formatResponse(ctx, ctx.Request(), "pong", "message")
}
