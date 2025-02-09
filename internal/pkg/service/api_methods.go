package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
)

func (s Service) DeleteMethodDelete(w http.ResponseWriter, r *http.Request, params api.DeleteMethodDeleteParams) {
	sendFormattedResponse(w, r, "DELETE", "method")
}

func (s Service) GetMethodGet(w http.ResponseWriter, r *http.Request, params api.GetMethodGetParams) {
	sendFormattedResponse(w, r, "GET", "method")
}

func (s Service) PatchMethodPatch(w http.ResponseWriter, r *http.Request, params api.PatchMethodPatchParams) {
	sendFormattedResponse(w, r, "PATCH", "method")
}

func (s Service) PostMethodPost(w http.ResponseWriter, r *http.Request, params api.PostMethodPostParams) {
	sendFormattedResponse(w, r, "POST", "method")
}

func (s Service) PutMethodPut(w http.ResponseWriter, r *http.Request, params api.PutMethodPutParams) {
	sendFormattedResponse(w, r, "PUT", "method")
}
