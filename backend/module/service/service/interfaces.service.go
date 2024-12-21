package service_serv

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/ent"
	service_repository "github.com/root9464/Ton-students/module/service/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

type IServService interface {
	CreateService(ctx context.Context, service *ent.Service) (*ent.Service, error)
}

type servService struct {
	logger    *logger.Logger
	validator *validator.Validate

	repo service_repository.IServiceRepository
}

func NewServService(
	logger *logger.Logger,
	validator *validator.Validate,
	repo service_repository.IServiceRepository,
) *servService {
	return &servService{
		logger:    logger,
		validator: validator,
		repo:      repo,
	}
}
