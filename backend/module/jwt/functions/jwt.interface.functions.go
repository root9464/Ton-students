package jwt_funcs

import (
	"crypto/ed25519"

	"github.com/go-playground/validator/v10"
	jwt_dto "github.com/root9464/Ton-students/module/jwt/dto"
	jwt_helpers "github.com/root9464/Ton-students/module/jwt/helpers"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IJwtFuncs = (*JwtFuncs)(nil)

type IJwtFuncs interface {
	GenerateKeyPair(userData jwt_dto.UserData) (*string, *string, error)
	Ping() string
}

type JwtFuncs struct {
	logger    *logger.Logger
	validator *validator.Validate

	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey

	helpers jwt_helpers.IJwtHelper
}

func NewJwtFuncs(logger *logger.Logger, validator *validator.Validate, privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey, helpers jwt_helpers.IJwtHelper) *JwtFuncs {
	return &JwtFuncs{
		logger:     logger,
		validator:  validator,
		privateKey: privateKey,
		publicKey:  publicKey,
		helpers:    helpers,
	}
}
