package service

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/marvinjwendt/httb/internal/pkg/config"
)

type Service struct {
	config     *config.Config
	validator  *validator.Validate
	translator ut.Translator
}

func NewService(config *config.Config) (*Service, error) {
	val, trans, err := initValidator()
	if err != nil {
		return nil, err
	}

	return &Service{
		config:     config,
		validator:  val,
		translator: trans,
	}, nil
}

func initValidator() (*validator.Validate, ut.Translator, error) {
	validate := validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	translator, found := uni.GetTranslator("en")
	if !found {
		return nil, nil, fmt.Errorf("could not find en translation")
	}
	validate = validator.New()
	err := entranslations.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to register default translation: %w", err)
	}
	return validate, translator, nil
}
