package repository

import (
	"errors"
	"guardpost-gin/config"
	"guardpost-gin/internal/models"
	requests "guardpost-gin/internal/requests/auth"
	"guardpost-gin/internal/utils"
)

func CreateUser(request *requests.RegisterRequest) error {

	// Check if user exists
	_, err := FindUserByEmail(request.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash Password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	// Create User Instance
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	res := config.DB.Create(user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
