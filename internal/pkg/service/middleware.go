package service

import (
	"net/http"
	"strconv"
	"time"
)

// DelayMiddleware is a custom middleware that reads the `delay` query parameter
// and sleeps for the specified duration in milliseconds.
func DelayMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		delayParam := r.URL.Query().Get("delay")

		if delayParam != "" {
			delayMs, err := strconv.Atoi(delayParam)
			if err != nil {
				http.Error(w, "Invalid delay parameter", http.StatusBadRequest)
				return
			}

			if delayMs < 0 || delayMs > 10000 {
				http.Error(w, "Delay must be between 0 and 10000 milliseconds", http.StatusBadRequest)
				return
			}

			if delayMs > 0 {
				time.Sleep(time.Duration(delayMs) * time.Millisecond)
			}
		}

		next.ServeHTTP(w, r)
	})
}
