package utils

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UUID uuid.UUID `json:"uuid"`
	jwt.RegisteredClaims
}

func Sign(UUID uuid.UUID, prikey *rsa.PrivateKey) (string, error) {
	now := time.Now()
	claims := &Claims{
		UUID: UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "SDCRAFT-OAuth2-Server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	out, err := token.SignedString(prikey)
	if err != nil {
		return "", err
	}
	return out, nil
}
