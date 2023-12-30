package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	connect()
	Migrate()
}

func connect() {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=photo port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbPort)
	conn, err := gorm.Open(postgres.Open(dbURI))

	if err != nil {
		panic(err.Error())
	}

	DB = conn
}
