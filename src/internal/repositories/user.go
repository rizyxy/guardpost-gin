package repositories

import (
	"errors"
	"guardpost-gin/config"
	"guardpost-gin/src/internal/requests"
	"guardpost-gin/src/internal/utils"
	"guardpost-gin/src/models"
)

func CreateUser(request *requests.RegisterRequest) error {

	_, err := FindUserByEmail(request.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

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
