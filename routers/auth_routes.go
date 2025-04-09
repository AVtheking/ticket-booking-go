package routes

import (
	"github.com/AVtheking/ticketo/controllers"
	"github.com/AVtheking/ticketo/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRoutes struct {
	db *gorm.DB
}

func NewAuthRoutes(db *gorm.DB) *AuthRoutes {
	return &AuthRoutes{
		db: db,
	}
}

func (r *AuthRoutes) RegisterRoutes(routes *gin.RouterGroup) {
	authService := services.NewAuthService(r.db)
	authController := controllers.NewAuthController(authService)
	authRoutes := routes.Group("/auth")
	{
		authRoutes.POST("/register", authController.RegisterUser)
		authRoutes.POST("/login", authController.LoginUser)
	}
}
