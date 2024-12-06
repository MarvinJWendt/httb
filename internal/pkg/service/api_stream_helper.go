package service

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
	"time"
)

// stream handles the raw streaming of data to the client.
// It writes data at specified intervals using the generator function.
func (s Service) stream(ctx echo.Context, interval *api.StreamInterval, generator func() (any, error)) error {
	// Determine the tick duration based on the provided interval
	var tickDuration time.Duration
	if interval == nil {
		tickDuration = 0
	} else {
		tickDuration = time.Duration(*interval) * time.Millisecond
	}

	// Retrieve the timeout from the service configuration
	timeout := s.config.Timeout

	// Channel to signal when streaming should stop
	done := make(chan struct{})

	// Handle client disconnect
	notify := ctx.Request().Context().Done()

	// Note: Removed SSE-specific headers to enable raw streaming
	// The Content-Type is expected to be set by the caller (e.g., streamJSON)

	// Flush the headers to ensure the client starts receiving data
	if f, ok := ctx.Response().Writer.(http.Flusher); ok {
		f.Flush()
	}

	// Initialize a ticker if a tick duration is specified
	var ticker *time.Ticker
	if tickDuration > 0 {
		ticker = time.NewTicker(tickDuration)
		defer ticker.Stop()
	}

	// Initialize a timer for the overall timeout
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case <-notify:
			// Client disconnected
			close(done)
			return nil
		case <-timer.C:
			// Timeout reached
			close(done)
			return errors.New("streaming timed out")
		case <-done:
			return nil
		default:
			// Generate data using the provided generator function
			data, err := generator()
			if err != nil {
				// Handle generator error by sending an error message to the client
				errorMsg := map[string]string{"error": err.Error()}
				jsonData, _ := json.Marshal(errorMsg)
				_, writeErr := ctx.Response().Write(append(jsonData, '\n'))
				if writeErr != nil {
					return writeErr
				}
				if f, ok := ctx.Response().Writer.(http.Flusher); ok {
					f.Flush()
				}
				return err
			}

			// Encode the generated data as JSON
			jsonData, err := json.Marshal(data)
			if err != nil {
				// Handle JSON encoding error
				errorMsg := map[string]string{"error": err.Error()}
				jsonData, _ = json.Marshal(errorMsg)
				_, writeErr := ctx.Response().Write(append(jsonData, '\n'))
				if writeErr != nil {
					return writeErr
				}
				if f, ok := ctx.Response().Writer.(http.Flusher); ok {
					f.Flush()
				}
				return err
			}

			// Write the JSON data followed by a newline to delimit JSON objects
			_, err = ctx.Response().Write(append(jsonData, '\n'))
			if err != nil {
				return err
			}

			// Flush the response to ensure the client receives the data immediately
			if f, ok := ctx.Response().Writer.(http.Flusher); ok {
				f.Flush()
			}

			// If a ticker is set, wait for the next tick before sending the next data chunk
			if ticker != nil {
				select {
				case <-ticker.C:
					continue
				case <-notify:
					close(done)
					return nil
				case <-timer.C:
					close(done)
					return errors.New("streaming timed out")
				}
			} else {
				// If no interval is set, send data continuously without delay
				continue
			}
		}
	}
}

// streamJSON sets the Content-Type to application/json and streams JSON-encoded data.
func (s Service) streamJSON(ctx echo.Context, interval *api.StreamInterval, generator func() (any, error)) error {
	// Set the Content-Type header for JSON streaming
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().Header().Set("Cache-Control", "no-cache")
	ctx.Response().Header().Set("Connection", "keep-alive")

	// Flush the headers to ensure the client starts receiving data
	if f, ok := ctx.Response().Writer.(http.Flusher); ok {
		f.Flush()
	}

	// Call the generic stream function with the provided generator
	return s.stream(ctx, interval, generator)
}
