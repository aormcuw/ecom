package auth

import (
	"time"

	"github.com/aormcuw/ecom/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, UserID string) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTDurationinSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID": UserID,
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
