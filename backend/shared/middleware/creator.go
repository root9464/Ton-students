package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	user_model "github.com/root9464/Ton-students/module/user/model"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

type RoleMiddleware struct {
	logger      *logger.Logger
	userRepo    user_repository.IUserRepository
	redisClient *redis.Client
	cacheMutex  sync.RWMutex
}

func NewRoleMiddleware(
	logger *logger.Logger,
	userRepo user_repository.IUserRepository,
	redisClient *redis.Client,
) *RoleMiddleware {
	return &RoleMiddleware{
		logger:      logger,
		userRepo:    userRepo,
		redisClient: redisClient,
	}
}

var rolePriority = map[user_model.Role]int{
	user_model.UserRole:    1,
	user_model.CreatorRole: 2,
	user_model.ModerRole:   3,
	user_model.AdminRole:   4,
}

func (rm *RoleMiddleware) CreatorOnly() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		userHash := ctx.Get("user_hash")
		rm.logger.Infof("Checking user role... user_hash: %s", userHash)

		if userHash == "" {
			rm.logger.Warn("Missing user_hash header")
			return ctx.Status(401).JSON(fiber.Map{
				"error": "Missing user_hash header",
			})
		}

		userChan := make(chan *user_model.User, 1)
		errChan := make(chan error, 1)

		go rm.getUserWithCache(ctxWithTimeout, userHash, userChan, errChan)

		var userInDb *user_model.User
		select {
		case user := <-userChan:
			userInDb = user
		case err := <-errChan:
			rm.logger.Errorf("Failed to retrieve user: %v", err)
			return ctx.Status(401).JSON(fiber.Map{
				"error": "User not found or unauthorized",
			})
		case <-ctxWithTimeout.Done():
			rm.logger.Warn("Context timeout while retrieving user")
			return ctx.Status(408).JSON(fiber.Map{
				"error": "Request timeout",
			})
		}

		rm.logger.Infof("User: %+v", userInDb)

		userRolePriority, userRoleExists := rolePriority[userInDb.Role]
		requiredRolePriority := rolePriority[user_model.CreatorRole]

		if !userRoleExists || userRolePriority < requiredRolePriority {
			rm.logger.Warnf("Access denied for user: %s, role: %s", userHash, userInDb.Role)
			return ctx.Status(403).JSON(fiber.Map{
				"error": "Access denied. Insufficient role privileges.",
			})
		}

		return ctx.Next()
	}
}

func (rm *RoleMiddleware) getUserWithCache(
	ctx context.Context,
	userHash string,
	userChan chan *user_model.User,
	errChan chan error,
) {
	requestID := fmt.Sprintf("%x", time.Now().UnixNano())
	cacheKey := fmt.Sprintf("user:hash:%s", userHash)

	rm.cacheMutex.RLock()
	cachedUser, redisErr := rm.redisClient.Get(ctx, cacheKey).Result()
	rm.cacheMutex.RUnlock()

	if redisErr == nil {
		user := new(user_model.User)
		if json.Unmarshal([]byte(cachedUser), &user) == nil {
			rm.logger.Infof("[%s] Redis hit - User: %s, Role: %s", requestID, userHash, user.Role)
			userChan <- user
			return
		}
	}

	if redisErr != redis.Nil {
		rm.logger.Warnf("[%s] Redis error: %v", requestID, redisErr)
	}

	rm.logger.Infof("[%s] Fetching user from DB: %s", requestID, userHash)
	user, err := rm.userRepo.GetByHash(ctx, userHash)
	if err != nil {
		rm.logger.Errorf("[%s] DB fetch failed: %v", requestID, err)
		errChan <- err
		return
	}

	go func() {
		rm.cacheMutex.Lock()
		defer rm.cacheMutex.Unlock()

		if userJSON, err := json.Marshal(user); err == nil {
			if cacheErr := rm.redisClient.Set(ctx, cacheKey, userJSON, 1*time.Hour).Err(); cacheErr == nil {
				rm.logger.Infof("[%s] Cached in Redis: %s", requestID, userHash)
			}
		}
	}()

	userChan <- user
}
