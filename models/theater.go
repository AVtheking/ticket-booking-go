package models

import (
	"time"

	"github.com/google/uuid"
)

type SeatType string

const (
	Regular SeatType = "regular"
	Premium SeatType = "premium"
)

type Theater struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid" `
	Name      string          `json:"name" gorm:"column:name"`
	Location  string          `json:"location" gorm:"column:location"`
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at"`
	Screens   []TheaterScreen `json:"screens" `
}

type TheaterScreen struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid" `
	Name      string    `json:"name" gorm:"column:name"`
	TheaterID uuid.UUID `json:"theater_id" gorm:"column:theater_id"`
	Seats     []Seat    `json:"seats" gorm:"foreignKey:TheaterScreenID"`
	Capacity  int       `json:"capacity" gorm:"column:capacity"`
}

type Seat struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid" `
	TheaterID       uuid.UUID `json:"theater_id" gorm:"column:theater_id"`
	TheaterScreenID uuid.UUID `json:"theater_screen_id" gorm:"column:theater_screen_id"`
	Row             int       `json:"row" gorm:"column:row"`
	Price           float64   `json:"price" gorm:"column:price"`
	Number          int       `json:"number" gorm:"column:number"`
	SeatType        SeatType  `json:"seat_type" gorm:"column:seat_type"`
}
