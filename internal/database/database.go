package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {

	fmt.Println("Setting up database...")

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbTable := os.Getenv("DB_TABLE")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")

	connectString := fmt.Sprintf("host=%s port=%s user%s password=%s dbName=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName, dbTable)

	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})

	return db, nil
}
