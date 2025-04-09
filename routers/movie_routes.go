package routes

import (
	"github.com/AVtheking/ticketo/controllers"
	"github.com/AVtheking/ticketo/middlewares"
	"github.com/AVtheking/ticketo/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieRoutes struct {
	db *gorm.DB
}

func NewMovieRoutes(db *gorm.DB) *MovieRoutes {
	return &MovieRoutes{db: db}
}

func (r *MovieRoutes) RegisterRoutes(router *gin.Engine) {
	movieService := services.NewMovieService(r.db)
	movieController := controllers.NewMovieController(movieService)
	protectedMovieRoutes := router.Group("/movies")

	protectedMovieRoutes.Use(middlewares.AuthMiddleware())
	{
		protectedMovieRoutes.GET("", movieController.GetMovies)
		protectedMovieRoutes.GET("/:id", movieController.GetMovieById)

	}

	adminMovieRoutes := router.Group("/admin/movies")
	adminMovieRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	{
		adminMovieRoutes.POST("", movieController.CreateMovie)
		adminMovieRoutes.PUT("/:id", movieController.UpdateMovie)
		adminMovieRoutes.DELETE("/:id", movieController.DeleteMovie)
	}
}
