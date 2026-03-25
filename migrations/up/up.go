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

	fmt.Println("Running migrations...")
	db.AutoMigrate(&models.Family{}, &models.SmallGroup{}, &models.Member{}, &models.Visitor{}, &models.Message{}, &models.Event{})

	// Create Indexes
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_members_email ON members(email);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_members_zone ON members(zone);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_members_family_id ON members(family_id);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_members_small_group_id ON members(small_group_id);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_members_deleted_at ON members(deleted_at);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_families_head_id ON families(head_id);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_families_deleted_at ON families(deleted_at);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_small_groups_zone ON small_groups(zone);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_small_groups_deleted_at ON small_groups(deleted_at);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_visitors_email ON visitors(email);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_visitors_deleted_at ON visitors(deleted_at);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_messages_date ON messages(date);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_messages_deleted_at ON messages(deleted_at);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_events_deleted_at ON events(deleted_at);`)

	fmt.Println("Migrations completed.")
}
