package service

import (
	"errors"
	"fmt"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
	"strings"
)

func (s Service) Validate(w http.ResponseWriter, data any) bool {
	err := defaults.Set(data)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, api.ValidationError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error setting default value: %s", err.Error()),
		})
		return false
	}

	err = s.validator.Struct(data)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			sendJSON(w, http.StatusBadRequest, api.ValidationError{
				Message:    "Validation failed: illegal argument was passed",
				StatusCode: http.StatusBadRequest,
			})
		}

		var result api.ValidationMessages
		var errorsMap validator.ValidationErrors
		errors.As(err, &errorsMap)
		translatedErrors := errorsMap.Translate(s.translator)
		for k, v := range translatedErrors {
			keySplit := strings.Split(k, ".")
			result = append(result, api.ValidationMessage{
				Field:   strings.ToLower(keySplit[len(keySplit)-1]),
				Message: v,
			})
		}
		sendJSON(w, http.StatusBadRequest, api.ValidationError{
			Message:    "Validation failed",
			StatusCode: http.StatusBadRequest,
			Errors:     &result,
		})
		return false
	}

	return true
}
