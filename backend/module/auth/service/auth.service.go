package auth_service

import (
	"time"

	"github.com/gofiber/fiber/v2"
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
