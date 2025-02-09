package service

import (
	"encoding/json"
	"net/http"
)

type ResponseFormat int

const (
	ResponseFormatInvalid ResponseFormat = iota
	ResponseFormatJSON
	ResponseFormatText
)

func prepareJSON(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
}

func sendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	prepareJSON(w, statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

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

func sendFormattedResponse(w http.ResponseWriter, r *http.Request, text, keyName string) {
	format := getFormat(r)

	switch format {
	case ResponseFormatJSON:
		sendJSON(w, http.StatusOK, map[string]string{keyName: text})
	case ResponseFormatText:
		_, _ = w.Write([]byte(text))
	default:
		sendError(w, http.StatusBadRequest, "invalid format")
	}
}
