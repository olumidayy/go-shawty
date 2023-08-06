package database

import (
	"fmt"
	"os"

	"github.com/olumidayy/go-shawty/src/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func CreateConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_DB_PORT"),
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database.")
	}

	fmt.Println("Connected to Database successfully.")

	database.AutoMigrate(&models.User{})

	return database
}