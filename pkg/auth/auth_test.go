package auth

import (
	"testing"
	"time"

	"github.com/miceremwirigi/journey-family-membership-backend/pkg/config"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Create a mock member and config
	member := &models.Member{
		Email: "test@example.com",
		Role:  "admin",
	}
	cfg := &config.Config{
		JWTSecret: "test_secret",
	}

	// Test token generation
	tokenString, err := GenerateToken(member, cfg)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if tokenString == "" {
		t.Fatal("Generated token is empty")
	}

	// Test token validation
	claims, err := ValidateToken(tokenString, cfg)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Email != member.Email {
		t.Errorf("Expected email %s, but got %s", member.Email, claims.Email)
	}

	if claims.Role != member.Role {
		t.Errorf("Expected role %s, but got %s", member.Role, claims.Role)
	}

	// Test expiration
	if claims.ExpiresAt.Time.Before(time.Now()) {
		t.Error("Token should not be expired yet")
	}
}

func TestValidateInvalidToken(t *testing.T) {
	cfg := &config.Config{
		JWTSecret: "test_secret",
	}

	// Test with an invalid token string
	_, err := ValidateToken("invalid_token", cfg)
	if err == nil {
		t.Error("Expected an error for invalid token, but got nil")
	}
}
