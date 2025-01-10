package jwt_helpers

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (h *jwtHelper) CreateJwt(claims jwt.Claims, key ed25519.PrivateKey) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
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

func (h *jwtHelper) CheckTokenExpiration(tokenString string, publicKey ed25519.PublicKey) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return false, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp, ok := claims["exp"].(float64)
		if !ok {
			return false, fmt.Errorf("exp field not found in token")
		}

		expirationTime := time.Unix(int64(exp), 0)
		if expirationTime.Before(time.Now()) {
			return false, nil
		}
		return true, nil
	}

	return false, fmt.Errorf("invalid token claims")
}
