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

	// Get JWT Secret
	key = []byte(os.Getenv("JWT_SECRET"))

	// Create JWT Token
	t = jwt.New(jwt.SigningMethodHS256)

	// Sign JWT Token
	s, err := t.SignedString(key)

	// Guard Clause
	if err != nil {
		return "", err
	}

	return s, nil
}
