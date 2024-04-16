package db

import (
	"log"
	"os"

	"github.com/detivenc/jsoncrudapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("POSTGRESQL_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Fail to Connect to DB")
	}

	DB.AutoMigrate(&models.Post{})
}
