package initializers

import (
	"fmt"
	"os"

	"github.com/detivenc/restapigo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializerDB() {
	dsn := os.Getenv("POSTGRES_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}
	DB.AutoMigrate(&model.User{})
}
