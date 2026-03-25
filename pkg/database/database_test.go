package database

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

func TestDatabaseIntegration(t *testing.T) {
	// Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set env to test
	os.Setenv("DB_NAME", "journey_family_test")

	// Connect to the database
	db, err := Connect()
	if err != nil {
		t.Fatalf("could not connect to test database: %v", err)
	}

	// Run migrations
	db.AutoMigrate(&models.Family{}, &models.SmallGroup{}, &models.Member{}, &models.Visitor{}, &models.Message{}, &models.Event{})

	// Create a new small group
	smallGroup := &models.SmallGroup{
		Name: "Test Small Group",
	}
	result := db.Create(smallGroup)
	if result.Error != nil {
		t.Fatalf("could not create small group: %v", result.Error)
	}

	// Create a new family
	family := &models.Family{
		Name: "Test Family",
	}
	result = db.Create(family)
	if result.Error != nil {
		t.Fatalf("could not create family: %v", result.Error)
	}

	// Create a new member
	member := &models.Member{
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@example.com",
		FamilyID:  family.ID,
		SmallGroupID: smallGroup.ID,
	}
	result = db.Create(member)
	if result.Error != nil {
		t.Fatalf("could not create member: %v", result.Error)
	}

	// Update family with head id
	family.HeadID = &member.ID
	result = db.Save(family)
	if result.Error != nil {
		t.Fatalf("could not update family: %v", result.Error)
	}

	// Update small group with leader id
	smallGroup.LeaderID = &member.ID
	result = db.Save(smallGroup)
	if result.Error != nil {
		t.Fatalf("could not update small group: %v", result.Error)
	}

	// Read the member
	var retrievedMember models.Member
	result = db.First(&retrievedMember, "email = ?", "test@example.com")
	if result.Error != nil {
		t.Fatalf("could not retrieve member: %v", result.Error)
	}
	if retrievedMember.FirstName != "Test" {
		t.Errorf("expected first name to be 'Test', but got '%s'", retrievedMember.FirstName)
	}

	// Delete the member
	result = db.Delete(&retrievedMember)
	if result.Error != nil {
		t.Fatalf("could not delete member: %v", result.Error)
	}

	// Delete the family
	result = db.Delete(&family)
	if result.Error != nil {
		t.Fatalf("could not delete family: %v", result.Error)
	}

	// Delete the small group
	result = db.Delete(&smallGroup)
	if result.Error != nil {
		t.Fatalf("could not delete small group: %v", result.Error)
	}

	// Drop tables
	db.Migrator().DropTable(&models.Event{}, &models.Message{}, &models.Visitor{}, &models.SmallGroup{}, &models.Family{}, &models.Member{})
}
