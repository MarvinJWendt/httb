package service

import (
	"context"
	"fmt"
	"net/http"

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
	server     *http.Server
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

// Shutdown gracefully shuts down the server
func (s *Service) Shutdown(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// GetHealth implements the /health endpoint for liveness probes.
func (s *Service) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"healthy"}`))
}

// GetReady implements the /ready endpoint for readiness probes.
func (s *Service) GetReady(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ready"}`))
}
