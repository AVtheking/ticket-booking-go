package models

import "gorm.io/gorm"

type Genre string

const (
	Action Genre = "action"
	Comedy Genre = "comedy"
	Drama  Genre = "drama"
	SciFi  Genre = "sci-fi"
)

type Movie struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PosterImage string   `json:"poster_image" gorm:"not null" column:"poster_image"`
	Cast        []string `json:"cast"`
	Genre       Genre    `json:"genre"`
	Year        int      `json:"year"`
	Rating      float64  `json:"rating"`
	Duration    int      `json:"duration"`
	PosterUrl   string   `json:"poster_url" gorm:"not null" column:"poster_url"`
}
