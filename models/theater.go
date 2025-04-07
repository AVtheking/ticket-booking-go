package models

import "time"

type Theater struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Movies    []Movie   `json:"movies"`
}

type TheaterScreens struct {
	
}
