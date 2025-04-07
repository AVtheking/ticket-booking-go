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
	ID        int             `json:"id" gorm:"column:id"`
	Name      string          `json:"name" gorm:"column:name"`
	Location  string          `json:"location" gorm:"column:location"`
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at"`
	Screens   []TheaterScreen `json:"screens" gorm:"foreignKey:TheaterID"`
}

type TheaterScreen struct {
	ID        uuid.UUID `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	TheaterID uuid.UUID `json:"theater_id" gorm:"column:theater_id"`
	Seats     []Seat    `json:"seats" gorm:"column:seats"`
	Capacity  int       `json:"capacity" gorm:"column:capacity"`
}

type Seat struct {
	ID              uuid.UUID `json:"id" gorm:"column:id"`
	TheaterID       uuid.UUID `json:"theater_id" gorm:"column:theater_id"`
	TheaterScreenID uuid.UUID `json:"theater_screen_id" gorm:"column:theater_screen_id"`
	Row             int       `json:"row" gorm:"column:row"`
	Number          int       `json:"number" gorm:"column:number"`
	SeatType        SeatType  `json:"seat_type" gorm:"column:seat_type"`
}
