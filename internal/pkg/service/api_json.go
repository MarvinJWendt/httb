package service

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/creasty/defaults"
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
	if params.MinDepth == nil {
		params.MinDepth = new(int)
		*params.MinDepth = 3
	}

	if params.MaxDepth == nil {
		params.MaxDepth = new(int)
		*params.MaxDepth = 5
	}

	if params.MaxElems == nil {
		params.MaxElems = new(int)
		*params.MaxElems = 3
	}

	sendJSON(w, http.StatusOK, RandomJSON(*params.MinDepth, *params.MaxDepth, *params.MaxElems))
}

func (s Service) GetJsonRandomUser(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomUserParams) {
	sendJSON(w, http.StatusOK, random.User())
}

func (s Service) GetJsonRandomUsers(w http.ResponseWriter, r *http.Request, params api.GetJsonRandomUsersParams) {
	_ = defaults.Set(&params)
	if ok := Validate[*api.GetJsonRandomUsersParams](w, &params); !ok {
		return
	}

	users := make([]api.User, *params.Count)
	for i := range users {
		users[i] = random.User()
	}

	sendJSON(w, http.StatusOK, users)
}

// RandomJSON generates a random JSON structure.
// - The root is always an object.
// - The JSON will have a minimum depth of minDepth and a maximum depth of maxDepth.
// - Each object or array will contain between 1 and maxElems elements.
func RandomJSON(minDepth, maxDepth, maxElems int) interface{} {
	// Root is always an object.
	return randomJSONObject(1, minDepth, maxDepth, maxElems)
}

// randomJSONValue generates a random JSON value based on the current depth.
// When depth < minDepth, it forces composite types (object or array) to ensure the structure is deep enough.
// When depth >= maxDepth, only primitives are allowed.
func randomJSONValue(depth, minDepth, maxDepth, maxElems int) interface{} {
	if depth >= maxDepth {
		return randomPrimitive()
	}

	// If we haven't reached the minimum depth, force composite types.
	if depth < minDepth {
		if gofakeit.Bool() {
			return randomJSONObject(depth+1, minDepth, maxDepth, maxElems)
		}
		return randomJSONArray(depth+1, minDepth, maxDepth, maxElems)
	}

	// Otherwise, choose randomly among composite and primitive types.
	choice := gofakeit.Number(0, 5)
	switch choice {
	case 0:
		return randomJSONObject(depth+1, minDepth, maxDepth, maxElems)
	case 1:
		return randomJSONArray(depth+1, minDepth, maxDepth, maxElems)
	case 2:
		return gofakeit.Sentence(3)
	case 3:
		return gofakeit.Float64()
	case 4:
		return gofakeit.Bool()
	case 5:
		return nil
	default:
		return nil
	}
}

// randomJSONObject creates a JSON object (map) with a random number of keys (between 1 and maxElems).
func randomJSONObject(depth, minDepth, maxDepth, maxElems int) map[string]interface{} {
	obj := make(map[string]interface{})
	numKeys := gofakeit.Number(1, maxElems)
	for i := 0; i < numKeys; i++ {
		key := gofakeit.Word()
		obj[key] = randomJSONValue(depth, minDepth, maxDepth, maxElems)
	}
	return obj
}

// randomJSONArray creates a JSON array with a random number of elements (between 1 and maxElems).
func randomJSONArray(depth, minDepth, maxDepth, maxElems int) []interface{} {
	numElems := gofakeit.Number(1, maxElems)
	arr := make([]interface{}, 0, numElems)
	for i := 0; i < numElems; i++ {
		arr = append(arr, randomJSONValue(depth, minDepth, maxDepth, maxElems))
	}
	return arr
}

// randomPrimitive returns a random JSON primitive: string, number, boolean, or null.
func randomPrimitive() interface{} {
	choice := gofakeit.Number(0, 3)
	switch choice {
	case 0:
		return gofakeit.Sentence(3)
	case 1:
		return gofakeit.Float64()
	case 2:
		return gofakeit.Bool()
	case 3:
		return nil
	default:
		return nil
	}
}
