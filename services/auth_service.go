package services

import (
	"errors"

	"github.com/AVtheking/ticketo/dto"
	"github.com/AVtheking/ticketo/models"
	"github.com/AVtheking/ticketo/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (s *AuthService) RegisterUser(request *dto.User) (*models.User, error) {
	var user models.User

	//check if the user with same email exists

	result := s.db.Find(&user, "email = ?", request.Email)
	if result.Error != nil {
		return nil, result.Error
	}

	if user.Email != "" {
		return nil, errors.New("user with same email already exists")
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = utils.HashPassword(request.Password)
	user.Role = models.UserRole

	result = s.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func (s *AuthService) LoginUser(request *dto.User) (*models.User, error) {
	var user models.User

	result := s.db.Find(&user, "email = ?", request.Email)
	if result.Error != nil {
		return nil, result.Error
	}

	if user.Email == "" {
		return nil, errors.New("user with same email already exists")
	}

	if !utils.CheckPassword(request.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
