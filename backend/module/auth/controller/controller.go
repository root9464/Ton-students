package auth_controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/ent"
	auth_service "github.com/root9464/Ton-students/backend/module/auth/service"
	"github.com/root9464/Ton-students/backend/shared/utils"
)

type Request struct {
	InitDataRaw string `json:"init_data_raw"`
}

func Authorize(log *utils.Logger, envs *config.Config, db *ent.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := new(Request)

		if err := ctx.BodyParser(request); err != nil {
			errResp := &fiber.Map{
				"message": err.Error(),
			}
			if err := ctx.Status(500).JSON(errResp); err != nil {
				log.Error("Failed to send error response:" + err.Error())
			}
			return err
		}

		expIn := 240 * time.Hour

		user, err := auth_service.ValidateInitData(request.InitDataRaw, envs.TelegramBotToken, expIn, log)
		if err != nil {
			errResp := &fiber.Map{
				"message": err.Error(),
			}
			if err := ctx.Status(500).JSON(errResp); err != nil {
				log.Error("Failed to send error response:" + err.Error())
			}
		}

		userInDb, err := auth_service.GetUserInDb(db, user.User.ID)
		if err != nil {
			errResp := &fiber.Map{
				"message": err.Error(),
			}
			if err := ctx.Status(500).JSON(errResp); err != nil {
				log.Error("Failed to send error response:" + err.Error())
			}
		}

		if userInDb != nil {
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  "success",
				"message": "get user",
				"user":    userInDb,
			})
		}

		createdUser, err := auth_service.CreateUser(db, user)
		if err != nil {
			errResp := &fiber.Map{
				"message": err.Error(),
			}
			if err := ctx.Status(500).JSON(errResp); err != nil {
				log.Error("Failed to send error response:" + err.Error())
			}
		}

		return ctx.Status(200).JSON(&fiber.Map{
			"status":  "success",
			"message": "create user",
			"user":    createdUser,
		})
	}
}
