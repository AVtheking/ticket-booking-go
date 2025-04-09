package routes

import (
	"github.com/AVtheking/ticketo/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TheaterRoutes(db *gorm.DB, routes *gin.RouterGroup) {

	theaterController := controllers.NewTheaterController(db)

	theaterRoutes := routes.Group("/theaters")
	{
		theaterRoutes.GET("", theaterController.GetTheaters)
		theaterRoutes.GET("/:id", theaterController.GetTheaterByID)
		theaterRoutes.POST("", theaterController.CreateTheater)
	}
}
