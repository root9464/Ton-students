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
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	accessToken, refreshToken, err := c.jwtModule.GenerateKeyPair(jwt_dto.UserData{
		ID:       user.ID,
		Username: user.VisibleName,
		Role:     string(user.Role),
	})

	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "Authorized successfully",
		"data":    user,
		"token": fiber.Map{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func (c *authController) RefreshAccessToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh_token")
	if refreshToken == "" {
		return ctx.Status(400).JSON(&fiber.Map{
			"status":  "failed",
			"message": "Refresh token is missing",
		})
	}

	accessToken := ctx.Cookies("access_token")
	if accessToken != "" {
		isValid, err := c.jwtHelpers.CheckTokenExpiration(accessToken, c.publicKey)
		if err != nil {
			return ctx.Status(401).JSON(&fiber.Map{
				"status":  "failed",
				"message": "Failed to validate access token",
			})
		}

		if isValid {
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  "failed",
				"message": "Access token is still valid",
			})
		}
	}

	newAccessToken, err := c.jwtModule.RefreshAccessToken(refreshToken, c.publicKey, c.privateKey)
	if err != nil {
		return ctx.Status(401).JSON(&fiber.Map{
			"status":  "failed",
			"message": "Invalid or expired refresh token",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *newAccessToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "Access token refreshed successfully",
		"token": fiber.Map{
			"accessToken": newAccessToken,
		},
	})
}

func (c *authController) JwtPing(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": c.jwtModule.Ping(),
	})
}
