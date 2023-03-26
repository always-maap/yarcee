package database

import (
	"api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin1234 dbname=yarcee port=5432"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
