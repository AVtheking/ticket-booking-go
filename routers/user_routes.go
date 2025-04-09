package routes

import (
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
	userRoutes := routes.Group("/users")
	{
		userRoutes.POST("/register", r.RegisterUser)
		userRoutes.POST("/login", r.LoginUser)
	}
}
