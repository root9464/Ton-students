package serv_repository

import (
	"github.com/root9464/Ton-students/database"
	"github.com/root9464/Ton-students/shared/logger"
)

type IServiceModuleRepository interface{}

type serviceRepository struct {
	db     *database.Database
	logger *logger.Logger
}

func NewServiceModuleRepository(db *database.Database, logger *logger.Logger) *serviceRepository {
	return &serviceRepository{
		db:     db,
		logger: logger,
	}
}
