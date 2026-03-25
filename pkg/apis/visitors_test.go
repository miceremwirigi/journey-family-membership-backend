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

// Mock VisitorService
type mockVisitorService struct {
    visitors map[string]models.Visitor
    nextID   uint
}

func newMockVisitorService() *mockVisitorService {
    return &mockVisitorService{
        visitors: make(map[string]models.Visitor),
        nextID:   1,
    }
}

func (m *mockVisitorService) GetVisitors() ([]models.Visitor, error) {
    visitors := make([]models.Visitor, 0, len(m.visitors))
    for _, v := range m.visitors {
        visitors = append(visitors, v)
    }
    return visitors, nil
}

func (m *mockVisitorService) GetVisitor(id string) (models.Visitor, error) {
    v, ok := m.visitors[id]
    if !ok {
        return models.Visitor{}, errors.New("visitor not found")
    }
    return v, nil
}

func (m *mockVisitorService) CreateVisitor(v *models.Visitor) error {
    v.ID = m.nextID
    m.visitors[strconv.Itoa(int(m.nextID))] = *v
    m.nextID++
    return nil
}

func (m *mockVisitorService) UpdateVisitor(id string, v *models.Visitor) error {
    if _, ok := m.visitors[id]; !ok {
        return errors.New("visitor not found")
    }
    m.visitors[id] = *v
    return nil
}

func (m *mockVisitorService) DeleteVisitor(id string) error {
    if _, ok := m.visitors[id]; !ok {
        return errors.New("visitor not found")
    }
    delete(m.visitors, id)
    return nil
}

// Mock VisitorService that returns errors
type mockErrorVisitorService struct{}

func (m *mockErrorVisitorService) GetVisitors() ([]models.Visitor, error) {
	return nil, errors.New("GetVisitors error")
}
func (m *mockErrorVisitorService) GetVisitor(id string) (models.Visitor, error) {
	return models.Visitor{}, errors.New("GetVisitor error")
}
func (m *mockErrorVisitorService) CreateVisitor(v *models.Visitor) error {
	return errors.New("CreateVisitor error")
}
func (m *mockErrorVisitorService) UpdateVisitor(id string, v *models.Visitor) error {
	return errors.New("UpdateVisitor error")
}
func (m *mockErrorVisitorService) DeleteVisitor(id string) error {
	return errors.New("DeleteVisitor error")
}

func TestVisitorHandler_ErrorPaths(t *testing.T) {
	app := fiber.New()
	mockService := &mockErrorVisitorService{}
	handler := NewVisitorHandler(mockService)

	app.Get("/visitors", handler.GetVisitors)
	app.Get("/visitors/:id", handler.GetVisitor)
	app.Post("/visitors", handler.CreateVisitor)
	app.Put("/visitors/:id", handler.UpdateVisitor)
	app.Delete("/visitors/:id", handler.DeleteVisitor)

	t.Run("GetVisitors error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/visitors", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("GetVisitor error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/visitors/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateVisitor error", func(t *testing.T) {
		v := &models.Visitor{Email: "visitor@example.com"}
		body, _ := json.Marshal(v)
		req := httptest.NewRequest("POST", "/visitors", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateVisitor bad request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/visitors", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("UpdateVisitor error", func(t *testing.T) {
		v := &models.Visitor{Email: "visitor@example.com"}
		body, _ := json.Marshal(v)
		req := httptest.NewRequest("PUT", "/visitors/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("UpdateVisitor bad request", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/visitors/1", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("DeleteVisitor error", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/visitors/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})
}
