package routes

import (
	"github.com/AVtheking/ticketo/controllers"
	"github.com/AVtheking/ticketo/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRoutes struct {
	db *gorm.DB
}

func NewUserRoutes(db *gorm.DB) *UserRoutes {
	return &UserRoutes{
		db: db,
	}
}

func (r *UserRoutes) RegisterRoutes(routes *gin.RouterGroup) {
	authService := services.NewAuthService(r.db)
	authController := controllers.NewAuthController(authService)
	userRoutes := routes.Group("/users")
	{
		userRoutes.POST("/register", authController.RegisterUser)
		userRoutes.POST("/login", authController.LoginUser)
	}
}
