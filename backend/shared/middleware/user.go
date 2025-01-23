package middleware

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
	jwt_helpers "github.com/root9464/Ton-students/module/jwt/helpers"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

type UserMiddleware struct {
	logger     *logger.Logger
	userRepo   user_repository.IUserRepository
	jwtHelpers jwt_helpers.IJwtHelper

	publicKey string
}

func NewUserMiddleware(
	logger *logger.Logger,
	userRepo user_repository.IUserRepository,
	jwtHelpers jwt_helpers.IJwtHelper,
	publicKey string,
) *UserMiddleware {
	return &UserMiddleware{
		logger:     logger,
		userRepo:   userRepo,
		jwtHelpers: jwtHelpers,
		publicKey:  publicKey,
	}
}

var whitelistUserHandlers = map[string]*regexp.Regexp{}

func (rm *RoleMiddleware) UserOnly() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		path := ctx.Path()
		for route, re := range whitelistUserHandlers {
			if path == route || (re != nil && re.MatchString(path)) {
				return ctx.Next()
			}
		}

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
