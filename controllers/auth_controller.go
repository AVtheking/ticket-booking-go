package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		db: db,
	}
}

func (c *AuthController) RegisterUser(ctx *gin.Context) {

}
