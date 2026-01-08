package repository

import (
	"guardpost-gin/config"
	"guardpost-gin/internal/models"
)

func CreateUser(user *models.User) error {
	res := config.DB.Create(user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func FindUserByEmail(email string) *models.User {
	var user models.User

	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil
	}

	return &user
}
