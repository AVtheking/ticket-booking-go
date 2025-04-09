package models

import "gorm.io/gorm"

type Role string

const (
	Admin    Role = "admin"
	UserRole Role = "user"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     Role   `json:"role"`
}
