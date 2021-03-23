package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ince01/note-server/internal/orm/models"
	"github.com/ince01/note-server/pkg/server"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectionString() string {
	var dbHost = os.Getenv("POSTGRES_HOST")
	var dbPort = os.Getenv("POSTGRES_PORT")
	var dbUserName = os.Getenv("POSTGRES_USER")
	var dbPassword = os.Getenv("POSTGRES_PASSWORD")
	var dbName = os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password='%s' sslmode=disable", dbHost, dbPort, dbUserName, dbName, dbPassword)
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&models.User{},
		&models.Note{},
	)

	server.Run(db)
}
