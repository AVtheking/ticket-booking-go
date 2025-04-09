package controllers

import (
	"net/http"

	"github.com/AVtheking/ticketo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TheaterController struct {
	db *gorm.DB
}

func NewTheaterController(db *gorm.DB) *TheaterController {
	return &TheaterController{
		db: db,
	}
}

func (config *TheaterController) GetTheaterByID(ctx *gin.Context) {
	theaterId := ctx.Param("id")
	var theater models.Theater

	err := config.db.First(&theater, theaterId).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, theater)
}

func (c *TheaterController) GetTheaters(ctx *gin.Context) {
	var theaters []models.Theater
	err := c.db.Find(&theaters).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, theaters)
}

func (c *TheaterController) CreateTheater(ctx *gin.Context) {
	var theater models.Theater

	if err := ctx.ShouldBindJSON(&theater); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.db.Create(&theater).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
