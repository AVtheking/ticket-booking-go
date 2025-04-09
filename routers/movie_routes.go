package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieRoutes struct {
	db *gorm.DB
}

func NewMovieRoutes(db *gorm.DB) *MovieRoutes {
	return &MovieRoutes{db: db}
}

func (r *MovieRoutes) SetupRoutes(router *gin.Engine) {
	movieRoutes := router.Group("/movies")
	{
		movieRoutes.GET("", r.GetMovies)
	}
}
