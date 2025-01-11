package jwt_helpers

import (
	"crypto/ed25519"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	jwt_dto "github.com/root9464/Ton-students/module/jwt/dto"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IJwtHelper = (*jwtHelper)(nil)

type IJwtHelper interface {
	CreateJwt(claims jwt.Claims, key ed25519.PrivateKey) (*string, error)
	VerifyJwt(tokenString string, key ed25519.PublicKey) (*jwt.Token, error)
	CheckTokenExpiration(token string, publicKey ed25519.PublicKey) (bool, error)
	ParseJwt(tokenString string, key ed25519.PublicKey) (*jwt_dto.UserJwtPayload, error)
}

type jwtHelper struct {
	logger    *logger.Logger
	validator *validator.Validate
}

func NewJwtHelper(logger *logger.Logger, validator *validator.Validate) *jwtHelper {
	return &jwtHelper{logger: logger, validator: validator}
}
