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

// Mock SmallGroupService
type mockSmallGroupService struct {
    smallGroups map[string]models.SmallGroup
    nextID      uint
}

func newMockSmallGroupService() *mockSmallGroupService {
    return &mockSmallGroupService{
        smallGroups: make(map[string]models.SmallGroup),
        nextID:      1,
    }
}

func (m *mockSmallGroupService) GetSmallGroups() ([]models.SmallGroup, error) {
    smallGroups := make([]models.SmallGroup, 0, len(m.smallGroups))
    for _, sg := range m.smallGroups {
        smallGroups = append(smallGroups, sg)
    }
    return smallGroups, nil
}

func (m *mockSmallGroupService) GetSmallGroup(id string) (models.SmallGroup, error) {
    sg, ok := m.smallGroups[id]
    if !ok {
        return models.SmallGroup{}, errors.New("small group not found")
    }
    return sg, nil
}

func (m *mockSmallGroupService) CreateSmallGroup(sg *models.SmallGroup) error {
    sg.ID = m.nextID
    m.smallGroups[strconv.Itoa(int(m.nextID))] = *sg
    m.nextID++
    return nil
}

func (m *mockSmallGroupService) UpdateSmallGroup(id string, sg *models.SmallGroup) error {
    if _, ok := m.smallGroups[id]; !ok {
        return errors.New("small group not found")
    }
    m.smallGroups[id] = *sg
    return nil
}

func (m *mockSmallGroupService) DeleteSmallGroup(id string) error {
    if _, ok := m.smallGroups[id]; !ok {
        return errors.New("small group not found")
    }
    delete(m.smallGroups, id)
    return nil
}

// Mock SmallGroupService that returns errors
type mockErrorSmallGroupService struct{}

func (m *mockErrorSmallGroupService) GetSmallGroups() ([]models.SmallGroup, error) {
	return nil, errors.New("GetSmallGroups error")
}
func (m *mockErrorSmallGroupService) GetSmallGroup(id string) (models.SmallGroup, error) {
	return models.SmallGroup{}, errors.New("GetSmallGroup error")
}
func (m *mockErrorSmallGroupService) CreateSmallGroup(sg *models.SmallGroup) error {
	return errors.New("CreateSmallGroup error")
}
func (m *mockErrorSmallGroupService) UpdateSmallGroup(id string, sg *models.SmallGroup) error {
	return errors.New("UpdateSmallGroup error")
}
func (m *mockErrorSmallGroupService) DeleteSmallGroup(id string) error {
	return errors.New("DeleteSmallGroup error")
}

func TestSmallGroupHandler_ErrorPaths(t *testing.T) {
	app := fiber.New()
	mockService := &mockErrorSmallGroupService{}
	handler := NewSmallGroupHandler(mockService)

	app.Get("/small-groups", handler.GetSmallGroups)
	app.Get("/small-groups/:id", handler.GetSmallGroup)
	app.Post("/small-groups", handler.CreateSmallGroup)
	app.Put("/small-groups/:id", handler.UpdateSmallGroup)
	app.Delete("/small-groups/:id", handler.DeleteSmallGroup)

	t.Run("GetSmallGroups error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/small-groups", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("GetSmallGroup error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/small-groups/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateSmallGroup error", func(t *testing.T) {
		sg := &models.SmallGroup{Name: "Test Small Group"}
		body, _ := json.Marshal(sg)
		req := httptest.NewRequest("POST", "/small-groups", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateSmallGroup bad request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/small-groups", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("UpdateSmallGroup error", func(t *testing.T) {
		sg := &models.SmallGroup{Name: "Test Small Group"}
		body, _ := json.Marshal(sg)
		req := httptest.NewRequest("PUT", "/small-groups/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("UpdateSmallGroup bad request", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/small-groups/1", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("DeleteSmallGroup error", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/small-groups/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})
}
