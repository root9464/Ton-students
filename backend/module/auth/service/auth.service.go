package auth_service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/ent"
	auth_repository "github.com/root9464/Ton-students/backend/module/auth/repository"
	"github.com/root9464/Ton-students/backend/shared/utils"
	tma "github.com/telegram-mini-apps/init-data-golang"
)

func ValidateInitData(initDataRaw string, botToken string, expIn time.Duration, log *utils.Logger) (*tma.InitData, error) {
	if err := tma.Validate(initDataRaw, botToken, expIn); err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	parseUser, err := tma.Parse(initDataRaw)
	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return &parseUser, nil
}

func GetUserInDb(db *ent.Client, id int64) (*ent.User, error) {
	getUserInDb, err := auth_repository.GetUserByID(db, id)
	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return getUserInDb, nil
}

func CreateUser(db *ent.Client, user *tma.InitData) (*ent.User, error) {
	createdUser, err := auth_repository.CreateUser(db, user)

	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return createdUser, nil
}
