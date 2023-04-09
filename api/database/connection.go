package database

import (
	"api/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() {
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		db       = os.Getenv("DB_NAME")
		port     = os.Getenv("DB_PORT")
	)

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, db, port)

	connection, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
