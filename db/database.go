package database

import (
	"github.com/AVtheking/ticketo/config"
	"github.com/AVtheking/ticketo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbConfig := config.NewConfig()
	db, err := gorm.Open(postgres.Open(dbConfig.Database.PostgresConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Theater{}, &models.TheaterScreen{}, &models.Seat{})
	return db
}
