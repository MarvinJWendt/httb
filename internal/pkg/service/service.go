package service

import (
	"github.com/marvinjwendt/httb/internal/pkg/config"
)

type Service struct {
	config *config.Config
}

func NewService(config *config.Config) *Service {
	return &Service{config: config}
}
