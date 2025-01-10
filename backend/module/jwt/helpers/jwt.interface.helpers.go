package jwt_helpers

import (
	"crypto/ed25519"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IJwtHelper = (*jwtHelper)(nil)

type IJwtHelper interface {
	CreateJwt(claims jwt.Claims, key ed25519.PrivateKey) (*string, error)
	VerifyJwt(tokenString string, key ed25519.PublicKey) (*jwt.Token, error)
}

type jwtHelper struct {
	logger    *logger.Logger
	validator *validator.Validate
}

func NewJwtHelper(logger *logger.Logger, validator *validator.Validate) *jwtHelper {
	return &jwtHelper{logger: logger, validator: validator}
}
