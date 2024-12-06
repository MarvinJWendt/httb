package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ResponseFormat int

const (
	ResponseFormatInvalid ResponseFormat = iota
	ResponseFormatJSON
	ResponseFormatText
)

// getFormat decides which format to use for the response.
// It first checks the query parameter `format` and then the `Accept` header.
// If no format is specified, it defaults to JSON.
func getFormat(r *http.Request) ResponseFormat {
	// Read format from query parameter
	formatParam := r.URL.Query().Get("format")
	switch formatParam {
	case "json":
		return ResponseFormatJSON
	case "text":
		return ResponseFormatText
	}

	// Read format from request accept header
	acceptHeader := r.Header.Get("Accept")
	switch acceptHeader {
	case "application/json":
		return ResponseFormatJSON
	case "text/plain":
		return ResponseFormatText
	}

	return ResponseFormatJSON
}

func formatResponse(ctx echo.Context, r *http.Request, text, keyName string) error {
	format := getFormat(r)

	switch format {
	case ResponseFormatJSON:
		return ctx.JSON(http.StatusOK, map[string]string{keyName: text})
	case ResponseFormatText:
		return ctx.String(http.StatusOK, text)
	default:
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid format"})
	}
}
