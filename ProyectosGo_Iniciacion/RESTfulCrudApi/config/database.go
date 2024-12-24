package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "mel.db.elephantsql.com"
	port     = 5432
	user     = "hgqoklyk"
	password = "zgSBVXOBS5hvk5aRyr8EZeeIg2JvTfPG"
	dbName   = "hgqoklyk"
)

func DatabaseConnection() *gorm.DB {
	sqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlinfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
