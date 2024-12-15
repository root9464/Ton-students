package service

import (
	"github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/shared/logger"
)

type IAuthService interface {
	Hello() string
}

var _ IAuthService = (*service)(nil)

type service struct {
	config *config.Config
	logger *logger.Logger
}

func NewService(config *config.Config, logger *logger.Logger) *service {
	return &service{
		config: config,
		logger: logger,
	}
}
