package service

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"net/http"
	"strings"
)

var (
	validate = validator.New()
	uni      *ut.UniversalTranslator
)

func Validate[T any](w http.ResponseWriter, data T) bool {

	enLocale := en.New()
	uni = ut.New(enLocale, enLocale)

	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	_ = entranslations.RegisterDefaultTranslations(validate, trans)

	// Validate the struct
	if err := validate.Struct(data); err != nil {
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
		translatedErrors := errorsMap.Translate(trans)
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
