package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShowTime struct {
	gorm.Model
	MovieID        uuid.UUID `json:"movie_id" gorm:"column:movie_id"`
	Movie          Movie     `json:"movie" gorm:"foreignKey:MovieID"`
	StartTime      time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime        time.Time `json:"end_time" gorm:"column:end_time"`
	AvailableSeats int       `json:"available_seats" gorm:"column:available_seats"`
	Price          float64   `json:"price" gorm:"column:price"`
}
