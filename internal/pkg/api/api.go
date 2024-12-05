// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Defines values for FormatParam.
const (
	FormatParamJson FormatParam = "json"
	FormatParamText FormatParam = "text"
)

// Defines values for GetGreetingParamsFormat.
const (
	GetGreetingParamsFormatJson GetGreetingParamsFormat = "json"
	GetGreetingParamsFormatText GetGreetingParamsFormat = "text"
)

// Defines values for GetPingParamsFormat.
const (
	Json GetPingParamsFormat = "json"
	Text GetPingParamsFormat = "text"
)

// DelayParam defines model for DelayParam.
type DelayParam = int

// FormatParam defines model for FormatParam.
type FormatParam string

// GetGreetingParams defines parameters for GetGreeting.
type GetGreetingParams struct {
	// Format Response format (default: `json`)
	Format *GetGreetingParamsFormat `form:"format,omitempty" json:"format,omitempty"`

	// Delay Delay in milliseconds before the response is sent (min: 0; max: 10000)
	Delay *DelayParam `form:"delay,omitempty" json:"delay,omitempty"`
}

// GetGreetingParamsFormat defines parameters for GetGreeting.
type GetGreetingParamsFormat string

// GetPingParams defines parameters for GetPing.
type GetPingParams struct {
	// Format Response format (default: `json`)
	Format *GetPingParamsFormat `form:"format,omitempty" json:"format,omitempty"`

	// Delay Delay in milliseconds before the response is sent (min: 0; max: 10000)
	Delay *DelayParam `form:"delay,omitempty" json:"delay,omitempty"`
}

// GetPingParamsFormat defines parameters for GetPing.
type GetPingParamsFormat string

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns a random greeting message.
	// (GET /greeting)
	GetGreeting(ctx echo.Context, params GetGreetingParams) error
	// Returns `pong`.
	// (GET /ping)
	GetPing(ctx echo.Context, params GetPingParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetGreeting converts echo context to params.
func (w *ServerInterfaceWrapper) GetGreeting(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetGreetingParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// ------------- Optional query parameter "delay" -------------

	err = runtime.BindQueryParameter("form", true, false, "delay", ctx.QueryParams(), &params.Delay)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter delay: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGreeting(ctx, params)
	return err
}

// GetPing converts echo context to params.
func (w *ServerInterfaceWrapper) GetPing(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPingParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// ------------- Optional query parameter "delay" -------------

	err = runtime.BindQueryParameter("form", true, false, "delay", ctx.QueryParams(), &params.Delay)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter delay: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPing(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/greeting", wrapper.GetGreeting)
	router.GET(baseURL+"/ping", wrapper.GetPing)

}
