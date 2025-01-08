package middleware

import (
	"github.com/gofiber/fiber/v2"
	user_model "github.com/root9464/Ton-students/module/user/model"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

type RoleMiddleware struct {
	logger   *logger.Logger
	userRepo user_repository.IUserRepository
}

func NewRoleMiddleware(logger *logger.Logger, userRepo user_repository.IUserRepository) *RoleMiddleware {
	return &RoleMiddleware{logger: logger, userRepo: userRepo}
}

var rolePriority = map[user_model.Role]int{
	user_model.UserRole:    1,
	user_model.CreatorRole: 2,
	user_model.ModerRole:   3,
	user_model.AdminRole:   4,
}

func (rm *RoleMiddleware) CreatorOnly() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userHash := ctx.Get("user_hash")
		if userHash == "" {
			rm.logger.Warn("Missing user_hash header")
			return ctx.Status(401).JSON(fiber.Map{
				"error": "Missing user_hash header",
			})
		}

		user, err := rm.userRepo.GetByHash(ctx.Context(), userHash)
		if err != nil {
			rm.logger.Errorf("Failed to retrieve user: %v", err)
			return ctx.Status(401).JSON(fiber.Map{
				"error": "User not found or unauthorized",
			})
		}

		userRolePriority, userRoleExists := rolePriority[user.Role]
		requiredRolePriority := rolePriority[user_model.CreatorRole]

		if !userRoleExists || userRolePriority < requiredRolePriority {
			rm.logger.Warnf("Access denied for user: %s, role: %s", userHash, user.Role)
			return ctx.Status(403).JSON(fiber.Map{
				"error": "Access denied. Insufficient role privileges.",
			})
		}

		return ctx.Next()
	}
}
