package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(username string) (string, error) {
	config := GetConfig()
	claims := customClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10000).Unix(),
			Issuer:    "mine",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Jwt.Secret))
}
