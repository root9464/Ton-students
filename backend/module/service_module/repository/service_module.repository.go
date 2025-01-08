package serv_repository

import (
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

type IServiceModuleRepository interface{}

type serviceRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewServiceModuleRepository(db *gorm.DB, logger *logger.Logger) *serviceRepository {
	return &serviceRepository{
		db:     db,
		logger: logger,
	}
}
