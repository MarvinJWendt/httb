package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
	"strings"
)

func (s Service) Validate(w http.ResponseWriter, data any) bool {
	// Validate the struct
	if err := s.validator.Struct(data); err != nil {
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
			x := strings.Split(k, ".")
			result = append(result, api.ValidationMessage{
				Field:   strings.ToLower(x[len(x)-1]),
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
