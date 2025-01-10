package jwt_module

import (
	"github.com/go-playground/validator/v10"
	jwt_funcs "github.com/root9464/Ton-students/module/jwt/functions"
	jwt_helpers "github.com/root9464/Ton-students/module/jwt/helpers"
	"github.com/root9464/Ton-students/shared/logger"
	"github.com/root9464/Ton-students/shared/utils"
	"gorm.io/gorm"
)

type JwtModule struct {
	jwtFuncs   jwt_funcs.IJwtFuncs
	jwtHelpers jwt_helpers.IJwtHelper

	privateKey string
	publicKey  string

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB
}

func (m *JwtModule) JwtHelpers() jwt_helpers.IJwtHelper {
	if m.jwtHelpers == nil {
		m.jwtHelpers = jwt_helpers.NewJwtHelper(m.logger, m.validator)
	}
	return m.jwtHelpers
}

func (m *JwtModule) JwtFuncs() jwt_funcs.IJwtFuncs {
	privateKey, publicKey, err := utils.HexToKeys(m.privateKey, m.publicKey)
	if err != nil {
		panic(err)
	}

	if m.jwtFuncs == nil {
		m.jwtFuncs = jwt_funcs.NewJwtFuncs(m.logger, m.validator, privateKey, publicKey, m.JwtHelpers())
	}
	return m.jwtFuncs
}

func NewJwtModule(logger *logger.Logger, validator *validator.Validate, db *gorm.DB, privateKey string, publicKey string) *JwtModule {
	return &JwtModule{
		logger:     logger,
		validator:  validator,
		db:         db,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}
