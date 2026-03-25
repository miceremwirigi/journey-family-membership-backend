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

// Mock EventService
type mockEventService struct {
    events map[string]models.Event
    nextID uint
}

func newMockEventService() *mockEventService {
    return &mockEventService{
        events: make(map[string]models.Event),
        nextID: 1,
    }
}

func (m *mockEventService) GetEvents() ([]models.Event, error) {
    events := make([]models.Event, 0, len(m.events))
    for _, e := range m.events {
        events = append(events, e)
    }
    return events, nil
}

func (m *mockEventService) GetEvent(id string) (models.Event, error) {
    e, ok := m.events[id]
    if !ok {
        return models.Event{}, errors.New("event not found")
    }
    return e, nil
}

func (m *mockEventService) CreateEvent(e *models.Event) error {
    e.ID = m.nextID
    m.events[strconv.Itoa(int(m.nextID))] = *e
    m.nextID++
    return nil
}

func (m *mockEventService) UpdateEvent(id string, e *models.Event) error {
    if _, ok := m.events[id]; !ok {
        return errors.New("event not found")
    }
    m.events[id] = *e
    return nil
}

func (m *mockEventService) DeleteEvent(id string) error {
    if _, ok := m.events[id]; !ok {
        return errors.New("event not found")
    }
    delete(m.events, id)
    return nil
}

// Mock EventService that returns errors
type mockErrorEventService struct{}

func (m *mockErrorEventService) GetEvents() ([]models.Event, error) {
	return nil, errors.New("GetEvents error")
}
func (m *mockErrorEventService) GetEvent(id string) (models.Event, error) {
	return models.Event{}, errors.New("GetEvent error")
}
func (m *mockErrorEventService) CreateEvent(e *models.Event) error {
	return errors.New("CreateEvent error")
}
func (m *mockErrorEventService) UpdateEvent(id string, e *models.Event) error {
	return errors.New("UpdateEvent error")
}
func (m *mockErrorEventService) DeleteEvent(id string) error {
	return errors.New("DeleteEvent error")
}

func TestEventHandler_ErrorPaths(t *testing.T) {
	app := fiber.New()
	mockService := &mockErrorEventService{}
	handler := NewEventHandler(mockService)

	app.Get("/events", handler.GetEvents)
	app.Get("/events/:id", handler.GetEvent)
	app.Post("/events", handler.CreateEvent)
	app.Put("/events/:id", handler.UpdateEvent)
	app.Delete("/events/:id", handler.DeleteEvent)

	t.Run("GetEvents error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/events", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("GetEvent error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/events/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateEvent error", func(t *testing.T) {
		e := &models.Event{Title: "Test Event"}
		body, _ := json.Marshal(e)
		req := httptest.NewRequest("POST", "/events", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateEvent bad request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/events", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("UpdateEvent error", func(t *testing.T) {
		e := &models.Event{Title: "Test Event"}
		body, _ := json.Marshal(e)
		req := httptest.NewRequest("PUT", "/events/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("UpdateEvent bad request", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/events/1", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("DeleteEvent error", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/events/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})
}
