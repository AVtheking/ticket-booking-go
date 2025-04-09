package routes

import (
	"github.com/AVtheking/ticketo/controllers"
	"github.com/AVtheking/ticketo/middlewares"
	"github.com/AVtheking/ticketo/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShowTimeRoutes struct {
	db *gorm.DB
}

func NewShowTimeRoutes(db *gorm.DB) *ShowTimeRoutes {
	return &ShowTimeRoutes{db: db}
}

func (s *ShowTimeRoutes) RegisterRoutes(routes *gin.RouterGroup) {
	showtimeService := services.NewShowTimeService(s.db)
	showtimeController := controllers.NewShowtimeController(showtimeService)
	showtimeRoutes := routes.Group("/showtimes")
	showtimeRoutes.Use(middlewares.AuthMiddleware())
	{
		showtimeRoutes.GET("/:movieID", showtimeController.GetShowtimesByMovieID)

	}
	showtimeRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	{
		showtimeRoutes.POST("", showtimeController.CreateShowtime)
		showtimeRoutes.PUT("/:showtimeID", showtimeController.UpdateShowtime)
		showtimeRoutes.DELETE("/:showtimeID", showtimeController.DeleteShowtime)
	}
}
