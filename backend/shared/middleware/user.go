package middleware

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (rm *Middleware) UserOnly() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		publicKeyBytes, err := hex.DecodeString(rm.publicKey)
		if err != nil {
			rm.logger.Errorf("Failed to decode public key: %v", err)
			return ctx.Status(500).JSON(fiber.Map{
				"error": "Failed to decode public key",
			})
		}

		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			rm.logger.Warn("Missing Authorization header")
			return ctx.Status(401).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		isValid, err := rm.jwtHelpers.CheckTokenExpiration(tokenString, ed25519.PublicKey(publicKeyBytes))
		if err != nil {
			return fmt.Errorf("token validation error: %v", err)
		}

		if !isValid {
			return fmt.Errorf("token has expired")
		}

		userPayload, err := rm.jwtHelpers.ParseJwt(tokenString, ed25519.PublicKey(publicKeyBytes))
		if err != nil {
			rm.logger.Errorf("Failed to parse JWT: %v", err)
			return ctx.Status(401).JSON(fiber.Map{
				"error": "Invalid or expired token",
				"msg":   err.Error(),
				"token": tokenString,
			})
		}

		rm.logger.Infof("User %d has role %s", userPayload.Sub, userPayload.Role)

		return ctx.Next()
	}
}
