package routes

import (
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
	reservationRoutes := routes.Group("/reservations")

}
