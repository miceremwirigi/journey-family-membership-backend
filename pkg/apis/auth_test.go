package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/common/dto"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/config"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// Mock MemberRepository
type mockMemberRepository struct {
	members map[string]*models.Member
}

func (m *mockMemberRepository) FindByEmail(email string) (*models.Member, error) {
	member, ok := m.members[email]
	if !ok {
		return nil, errors.New("member not found")
	}
	return member, nil
}

func newMockMemberRepository() *mockMemberRepository {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	return &mockMemberRepository{
		members: map[string]*models.Member{
			"test@example.com": {
				Email:    "test@example.com",
				Password: string(hashedPassword),
				Role:     "admin",
			},
		},
	}
}
func (m *mockMemberRepository) GetMembers() ([]models.Member, error) {
	return nil, nil
}
func (m *mockMemberRepository) GetMember(id string) (models.Member, error) {
	return models.Member{}, nil
}
func (m *mockMemberRepository) CreateMember(member *models.Member) error {
	return nil
}
func (m *mockMemberRepository) UpdateMember(id string, member *models.Member) error {
	return nil
}
func (m *mockMemberRepository) DeleteMember(id string) error {
	return nil
}


func TestLogin(t *testing.T) {
	app := fiber.New()
	mockRepo := newMockMemberRepository()
	cfg := &config.Config{JWTSecret: "test_secret"}

	app.Post("/login", func(c *fiber.Ctx) error {
		return Login(c, mockRepo, cfg)
	})

	t.Run("successful login", func(t *testing.T) {
		loginReq := dto.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(loginReq)
		req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		loginReq := dto.LoginRequest{
			Email:    "wrong@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(loginReq)
		req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected status %d, got %d", fiber.StatusUnauthorized, resp.StatusCode)
		}
	})

	t.Run("invalid password", func(t *testing.T) {
		loginReq := dto.LoginRequest{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}
		body, _ := json.Marshal(loginReq)
		req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected status %d, got %d", fiber.StatusUnauthorized, resp.StatusCode)
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})
}
