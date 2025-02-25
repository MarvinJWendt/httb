package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
)

func (s Service) DeleteReturn(w http.ResponseWriter, r *http.Request, _ api.DeleteReturnParams) {
	sendRequestData(w, r)
}

func (s Service) GetReturn(w http.ResponseWriter, r *http.Request, _ api.GetReturnParams) {
	sendRequestData(w, r)
}

func (s Service) PatchReturn(w http.ResponseWriter, r *http.Request, _ api.PatchReturnParams) {
	sendRequestData(w, r)
}

func (s Service) PostReturn(w http.ResponseWriter, r *http.Request, _ api.PostReturnParams) {
	sendRequestData(w, r)
}

func (s Service) PutReturn(w http.ResponseWriter, r *http.Request, _ api.PutReturnParams) {
	sendRequestData(w, r)
}

func sendRequestData(w http.ResponseWriter, r *http.Request) {
	var body []byte
	if r.Body != nil {
		body = make([]byte, r.ContentLength)
		_, _ = r.Body.Read(body)
	}

	sendJSON(w, http.StatusOK, map[string]interface{}{
		"method":     r.Method,
		"path":       r.URL.Path,
		"query":      r.URL.Query(),
		"headers":    r.Header,
		"uri":        r.RequestURI,
		"data":       string(body),
		"remoteAddr": r.RemoteAddr,
	})
}
