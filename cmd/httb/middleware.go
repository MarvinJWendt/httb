package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// DelayMiddleware is a custom Echo middleware that reads the `delay` query parameter
// and sleeps for the specified duration in milliseconds.
func DelayMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		delayParam := c.QueryParam("delay")

		if delayParam != "" {
			delayMs, err := strconv.Atoi(delayParam)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid delay parameter")
			}

			if delayMs < 0 || delayMs > 10000 {
				return echo.NewHTTPError(http.StatusBadRequest, "Delay must be between 0 and 10000 milliseconds")
			}

			if delayMs > 0 {
				time.Sleep(time.Duration(delayMs) * time.Millisecond)
			}
		}

		return next(c)
	}
}
