package auth_controller

import (
	"github.com/gofiber/fiber/v2"
	auth_dto "github.com/root9464/Ton-students/module/auth/dto"
	jwt_dto "github.com/root9464/Ton-students/module/jwt/dto"
	"github.com/root9464/Ton-students/shared/utils"
)

func (c *authController) Authorize(ctx *fiber.Ctx) error {
	data := new(auth_dto.AutorizeDto)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	user, err := c.authService.Authorize(ctx.Context(), data)
	if err != nil {
		if errorResponse := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(400).JSON(errorResponse)
		}
	}

	accessToken, refreshToken, err := c.jwtModule.GenerateKeyPair(jwt_dto.UserData{
		ID:       user.Data.ID,
		Username: user.Data.Username,
	})

	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User authorized",
		"data":    user,
		"token": fiber.Map{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func (c *authController) JwtPing(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": c.jwtModule.Ping(),
	})
}
