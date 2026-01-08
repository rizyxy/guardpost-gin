package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(token string) bool {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false
	}

	return jwtToken.Valid
}
