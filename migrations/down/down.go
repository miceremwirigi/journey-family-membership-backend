package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/database"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	fmt.Println("Dropping tables...")
	db.Migrator().DropTable(&models.Event{}, &models.Message{}, &models.Visitor{}, &models.SmallGroup{}, &models.Family{}, &models.Member{})
	fmt.Println("Tables dropped.")
}
