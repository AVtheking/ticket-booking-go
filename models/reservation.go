package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationStatus string

const (
	Pending   ReservationStatus = "pending"
	Confirmed ReservationStatus = "confirmed"
	Canceled  ReservationStatus = "canceled"
)

type Reservation struct {
	gorm.Model
	UserID            uuid.UUID         `json:"user_id" gorm:"column:user_id"`
	ShowTimeID        uuid.UUID         `json:"show_time_id" gorm:"column:show_time_id"`
	ShowTime          ShowTime          `json:"show_time" gorm:"foreignKey:ShowTimeID"`
	SeatNumbers       []int             `json:"seat_numbers" gorm:"column:seat_numbers"`
	TotalPrice        float64           `json:"total_price" gorm:"column:total_price"`
	ReservationStatus ReservationStatus `json:"reservation_status" gorm:"column:reservation_status"`
}
