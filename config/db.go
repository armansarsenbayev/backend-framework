package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dsn := fmt.Sprintf("host=%s user=postgres password=1234 dbname=food_delivery_v2 port=5432 sslmode=disable", dbHost)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB = database
	log.Println("Database connection successful!")
}

