package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"github.com/marvinjwendt/httb/internal/pkg/random"
	"net/http"
)

func (s Service) GetJsonRandomLog(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomLogParams) {
	if params.Count == nil {
		params.Count = new(int)
		*params.Count = 1
	}

	if params.LogLevels == nil {
		params.LogLevels = new([]string)
		*params.LogLevels = []string{"debug", "info", "warn", "error"}
	}

	if params.LogLevelWeights == nil {
		params.LogLevelWeights = new([]float32)
		*params.LogLevelWeights = []float32{1, 5, 2, 1}
	}

	if len(*params.LogLevels) != len(*params.LogLevelWeights) {
		sendError(w, http.StatusBadRequest, "log levels and weights must have the same length")
	}

	logLevels := make(map[string]float32)
	for i, level := range *params.LogLevels {
		logLevels[level] = (*params.LogLevelWeights)[i]
	}

	prepareJSON(w, http.StatusOK)
	_, _ = w.Write([]byte(random.NewLog(*params.Count, logLevels).String()))
}

func (s Service) GetJsonRandomAddress(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomAddressParams) {
	sendJSON(w, http.StatusOK, random.Address())
}

func (s Service) GetJsonRandomAddresses(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomAddressesParams) {
	if params.Count == nil {
		params.Count = new(int)
		*params.Count = 10
	}

	addresses := make([]api.Address, *params.Count)
	for i := range addresses {
		addresses[i] = random.Address()
	}

	sendJSON(w, http.StatusOK, addresses)
}

func (s Service) GetJsonRandomContact(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomContactParams) {
	sendJSON(w, http.StatusOK, random.Contact())
}

func (s Service) GetJsonRandomContacts(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomContactsParams) {
	if params.Count == nil {
		params.Count = new(int)
		*params.Count = 10
	}

	contacts := make([]api.Contact, *params.Count)
	for i := range contacts {
		contacts[i] = random.Contact()
	}

	sendJSON(w, http.StatusOK, contacts)
}

func (s Service) GetJsonRandom(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomParams) {
	panic("not implemented")
}

func (s Service) GetJsonRandomUser(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomUserParams) {
	sendJSON(w, http.StatusOK, random.User())
}

func (s Service) GetJsonRandomUsers(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomUsersParams) {
	if params.Count == nil {
		params.Count = new(int)
		*params.Count = 10
	}

	users := make([]api.User, *params.Count)
	for i := range users {
		users[i] = random.User()
	}

	sendJSON(w, http.StatusOK, users)
}
