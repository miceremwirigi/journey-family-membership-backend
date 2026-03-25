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

// Mock FamilyService
type mockFamilyService struct {
    families map[string]models.Family
    nextID   uint
}

func newMockFamilyService() *mockFamilyService {
    return &mockFamilyService{
        families: make(map[string]models.Family),
        nextID:   1,
    }
}

func (m *mockFamilyService) GetFamilies() ([]models.Family, error) {
    families := make([]models.Family, 0, len(m.families))
    for _, family := range m.families {
        families = append(families, family)
    }
    return families, nil
}

func (m *mockFamilyService) GetFamily(id string) (models.Family, error) {
    family, ok := m.families[id]
    if !ok {
        return models.Family{}, errors.New("family not found")
    }
    return family, nil
}

func (m *mockFamilyService) CreateFamily(family *models.Family) error {
    family.ID = m.nextID
    m.families[strconv.Itoa(int(m.nextID))] = *family
    m.nextID++
    return nil
}

func (m *mockFamilyService) UpdateFamily(id string, family *models.Family) error {
    if _, ok := m.families[id]; !ok {
        return errors.New("family not found")
    }
    m.families[id] = *family
    return nil
}

func (m *mockFamilyService) DeleteFamily(id string) error {
    if _, ok := m.families[id]; !ok {
        return errors.New("family not found")
    }
    delete(m.families, id)
    return nil
}

// Mock FamilyService that returns errors
type mockErrorFamilyService struct{}

func (m *mockErrorFamilyService) GetFamilies() ([]models.Family, error) {
	return nil, errors.New("GetFamilies error")
}
func (m *mockErrorFamilyService) GetFamily(id string) (models.Family, error) {
	return models.Family{}, errors.New("GetFamily error")
}
func (m *mockErrorFamilyService) CreateFamily(family *models.Family) error {
	return errors.New("CreateFamily error")
}
func (m *mockErrorFamilyService) UpdateFamily(id string, family *models.Family) error {
	return errors.New("UpdateFamily error")
}
func (m *mockErrorFamilyService) DeleteFamily(id string) error {
	return errors.New("DeleteFamily error")
}

func TestFamilyHandler_ErrorPaths(t *testing.T) {
	app := fiber.New()
	mockService := &mockErrorFamilyService{}
	handler := NewFamilyHandler(mockService)

	app.Get("/families", handler.GetFamilies)
	app.Get("/families/:id", handler.GetFamily)
	app.Post("/families", handler.CreateFamily)
	app.Put("/families/:id", handler.UpdateFamily)
	app.Delete("/families/:id", handler.DeleteFamily)

	t.Run("GetFamilies error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/families", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("GetFamily error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/families/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateFamily error", func(t *testing.T) {
		family := &models.Family{Name: "Test Family"}
		body, _ := json.Marshal(family)
		req := httptest.NewRequest("POST", "/families", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateFamily bad request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/families", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("UpdateFamily error", func(t *testing.T) {
		family := &models.Family{Name: "Test Family"}
		body, _ := json.Marshal(family)
		req := httptest.NewRequest("PUT", "/families/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("UpdateFamily bad request", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/families/1", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("DeleteFamily error", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/families/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})
}
