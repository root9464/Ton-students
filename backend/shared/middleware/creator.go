package middleware

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
	jwt_helpers "github.com/root9464/Ton-students/module/jwt/helpers"
	user_model "github.com/root9464/Ton-students/module/user/model"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

type Middleware struct {
	logger     *logger.Logger
	userRepo   user_repository.IUserRepository
	jwtHelpers jwt_helpers.IJwtHelper

	publicKey string
}

func NewMiddleware(
	logger *logger.Logger,
	userRepo user_repository.IUserRepository,
	jwtHelpers jwt_helpers.IJwtHelper,
	publicKey string,
) *Middleware {
	return &Middleware{
		logger:     logger,
		userRepo:   userRepo,
		jwtHelpers: jwtHelpers,
		publicKey:  publicKey,
	}
}

var rolePriority = map[user_model.Role]int{
	user_model.UserRole:    1,
	user_model.CreatorRole: 2,
	user_model.ModerRole:   3,
	user_model.AdminRole:   4,
}

var whitelist = map[string]*regexp.Regexp{
	"/api/creator/service/all-services": nil,                                                            // прямой доступ
	"/api/creator/service/get/":         regexp.MustCompile(`^/api/creator/service/get/[0-9a-fA-F-]+$`), //выражение валидатор
	"/api/creator/service/get-service":  nil,
	"/api/creator/service/feed":         nil,
}

func (rm *Middleware) CreatorOnly() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		path := ctx.Path()
		for route, re := range whitelist {
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
		if rolePriority[userPayload.Role] < rolePriority[user_model.CreatorRole] {
			return ctx.Status(403).JSON(fiber.Map{
				"error": "Only creator role is allowed",
			})
		}
		return ctx.Next()
	}
}
