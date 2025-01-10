package jwt_funcs

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	jwt_dto "github.com/root9464/Ton-students/module/jwt/dto"
)

func (f *JwtFuncs) GenerateKeyPair(userData jwt_dto.UserData) (*string, *string, error) {
	f.logger.Info("Validating user data...")
	if err := f.validator.Struct(userData); err != nil {
		f.logger.Warnf("validate error: %s", err.Error())
		return nil, nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}
	f.logger.Info("Validating success...")
	userRaw := new(string)

	if userData.LastName == nil {
		*userRaw = fmt.Sprintf("%s%s%d", userData.Username, userData.FirstName, userData.ID)
	} else {
		*userRaw = fmt.Sprintf("%s%s%d%s", userData.Username, userData.FirstName, userData.ID, *userData.LastName)
	}

	hash := sha256.New()
	hash.Write([]byte(*userRaw))
	refinedHash := hex.EncodeToString(hash.Sum(nil))

	accessClaims := jwt.MapClaims{
		"iss":       "Ton-students::admin",
		"sub":       userData.ID,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(15 * time.Minute).Unix(),
		"user_hash": refinedHash,
	}

	refreshClaims := jwt.MapClaims{
		"iss":       "Ton-students::admin",
		"sub":       userData.ID,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"user_hash": refinedHash,
	}

	accessToken, err := f.helpers.CreateJwt(accessClaims, f.privateKey)

	if err != nil {
		f.logger.Warnf("create access token error: %s", err.Error())
		return nil, nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	refreshToken, err := f.helpers.CreateJwt(refreshClaims, f.privateKey)

	if err != nil {
		f.logger.Warnf("create refresh token error: %s", err.Error())
		return nil, nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return &accessToken, &refreshToken, nil
}

func (f *JwtFuncs) Ping() string {
	return "pong"
}
