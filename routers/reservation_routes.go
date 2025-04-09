package routes

import (
	"github.com/AVtheking/ticketo/controllers"
	"github.com/AVtheking/ticketo/middlewares"
	"github.com/AVtheking/ticketo/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReservationRoutes struct {
	db *gorm.DB
}

func NewReservationRoutes(db *gorm.DB) *ReservationRoutes {
	return &ReservationRoutes{
		db: db,
	}
}

func (r *ReservationRoutes) RegisterRoutes(routes *gin.RouterGroup) {
	reservationService := services.NewReservationService(r.db)
	reservationController := controllers.NewRegisteredControllers(reservationService)
	reservationRoutes := routes.Group("/reservations")
	reservationRoutes.Use(middlewares.AuthMiddleware())
	{

		reservationRoutes.POST("", reservationController.CreateReservation)
		reservationRoutes.GET("/available-seats/:showTimeID", reservationController.GetAvailableSeats)
		reservationRoutes.GET("/:reservationID", reservationController.GetReservation)
		reservationRoutes.GET("/user/:userID", reservationController.GetUserReservations)
		reservationRoutes.GET("/all", reservationController.GetAllReservations)
	}

}
