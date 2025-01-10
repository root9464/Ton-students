package jwt_helpers

import (
	"crypto/ed25519"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func (h *jwtHelper) CreateJwt(claims jwt.Claims, key ed25519.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (h *jwtHelper) VerifyJwt(tokenString string, key ed25519.PublicKey) (*jwt.Token, error) {
	h.logger.Info("Verifying JWT token...")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	h.logger.Info("JWT token verified successfully")
	return token, nil
}
