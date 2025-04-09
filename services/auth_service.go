package services

import (
	"errors"

	"github.com/AVtheking/ticketo/dto"
	"github.com/AVtheking/ticketo/models"
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

func (s *AuthService) RegisterUser(request *dto.User) error {
	var user models.User

	//check if the user with same email exists

	result := s.db.Find(&user, "email = ?", request.Email)
	if result.Error != nil {
		return result.Error
	}

	if user.Email != "" {
		return errors.New("user with same email already exists")
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password

}
