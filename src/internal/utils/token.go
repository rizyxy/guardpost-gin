package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT() (string, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)
	key = []byte(os.Getenv("JWT_SECRET"))

	t = jwt.New(jwt.SigningMethodHS256)

	s, err := t.SignedString(key)

	if err != nil {
		return "", err
	}

	return s, nil
}

func GenerateAccessAndRefreshToken() (string, string, error) {
	accessToken, err := GenerateJWT()

	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateJWT()

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func VerifyJWT(token string) bool {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false
	}

	return jwtToken.Valid
}
