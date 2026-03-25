package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

// Mock MemberService
type mockMemberService struct {
	members map[string]models.Member
	nextID  uint
}

func newMockMemberService() *mockMemberService {
	return &mockMemberService{
		members: make(map[string]models.Member),
		nextID:  1,
	}
}

func (m *mockMemberService) GetMembers() ([]models.Member, error) {
	members := make([]models.Member, 0, len(m.members))
	for _, member := range m.members {
		members = append(members, member)
	}
	return members, nil
}

func (m *mockMemberService) GetMember(id string) (models.Member, error) {
	member, ok := m.members[id]
	if !ok {
		return models.Member{}, errors.New("member not found")
	}
	return member, nil
}

func (m *mockMemberService) CreateMember(member *models.Member) error {
	member.ID = m.nextID
	m.members[strconv.Itoa(int(m.nextID))] = *member
	m.nextID++
	return nil
}

func (m *mockMemberService) UpdateMember(id string, member *models.Member) error {
	if _, ok := m.members[id]; !ok {
		return errors.New("member not found")
	}
	m.members[id] = *member
	return nil
}

func (m *mockMemberService) DeleteMember(id string) error {
	if _, ok := m.members[id]; !ok {
		return errors.New("member not found")
	}
	delete(m.members, id)
	return nil
}

// Mock MemberService that returns errors
type mockErrorMemberService struct{}

func (m *mockErrorMemberService) GetMembers() ([]models.Member, error) {
	return nil, errors.New("GetMembers error")
}
func (m *mockErrorMemberService) GetMember(id string) (models.Member, error) {
	return models.Member{}, errors.New("GetMember error")
}
func (m *mockErrorMemberService) CreateMember(member *models.Member) error {
	return errors.New("CreateMember error")
}
func (m *mockErrorMemberService) UpdateMember(id string, member *models.Member) error {
	return errors.New("UpdateMember error")
}
func (m *mockErrorMemberService) DeleteMember(id string) error {
	return errors.New("DeleteMember error")
}

func TestMemberHandler_ErrorPaths(t *testing.T) {
	app := fiber.New()
	mockService := &mockErrorMemberService{}
	handler := NewMemberHandler(mockService)

	app.Get("/members", handler.GetMembers)
	app.Get("/members/:id", handler.GetMember)
	app.Post("/members", handler.CreateMember)
	app.Put("/members/:id", handler.UpdateMember)
	app.Delete("/members/:id", handler.DeleteMember)

	t.Run("GetMembers error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/members", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("GetMember error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/members/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateMember error", func(t *testing.T) {
		member := &models.Member{Email: "test@example.com"}
		body, _ := json.Marshal(member)
		req := httptest.NewRequest("POST", "/members", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateMember bad request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/members", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("UpdateMember error", func(t *testing.T) {
		member := &models.Member{Email: "test@example.com"}
		body, _ := json.Marshal(member)
		req := httptest.NewRequest("PUT", "/members/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("UpdateMember bad request", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/members/1", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("DeleteMember error", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/members/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})
}

