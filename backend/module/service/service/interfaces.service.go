package service_serv

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/ent"
	service_dto "github.com/root9464/Ton-students/module/service/dto"
	service_repository "github.com/root9464/Ton-students/module/service/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IServService = (*servService)(nil)

type IServService interface {
	CreateService(ctx context.Context, dto *service_dto.CreateServiceDto) (*ent.Service, error)
}

type servService struct {
	logger    *logger.Logger
	validator *validator.Validate
	db        *ent.Client
	repo      service_repository.IServiceRepository
}

func NewServService(
	logger *logger.Logger,
	validator *validator.Validate,
	db *ent.Client,
	repo service_repository.IServiceRepository,
) *servService {
	return &servService{
		logger:    logger,
		validator: validator,
		db:        db,
		repo:      repo,
	}
}
