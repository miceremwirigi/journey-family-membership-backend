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

// Mock MessageService
type mockMessageService struct {
    messages map[string]models.Message
    nextID   uint
}

func newMockMessageService() *mockMessageService {
    return &mockMessageService{
        messages: make(map[string]models.Message),
        nextID:   1,
    }
}

func (m *mockMessageService) GetMessages() ([]models.Message, error) {
    messages := make([]models.Message, 0, len(m.messages))
    for _, msg := range m.messages {
        messages = append(messages, msg)
    }
    return messages, nil
}

func (m *mockMessageService) GetMessage(id string) (models.Message, error) {
    msg, ok := m.messages[id]
    if !ok {
        return models.Message{}, errors.New("message not found")
    }
    return msg, nil
}

func (m *mockMessageService) CreateMessage(msg *models.Message) error {
    msg.ID = m.nextID
    m.messages[strconv.Itoa(int(m.nextID))] = *msg
    m.nextID++
    return nil
}

func (m *mockMessageService) UpdateMessage(id string, msg *models.Message) error {
    if _, ok := m.messages[id]; !ok {
        return errors.New("message not found")
    }
    m.messages[id] = *msg
    return nil
}

func (m *mockMessageService) DeleteMessage(id string) error {
    if _, ok := m.messages[id]; !ok {
        return errors.New("message not found")
    }
    delete(m.messages, id)
    return nil
}

// Mock MessageService that returns errors
type mockErrorMessageService struct{}

func (m *mockErrorMessageService) GetMessages() ([]models.Message, error) {
	return nil, errors.New("GetMessages error")
}
func (m *mockErrorMessageService) GetMessage(id string) (models.Message, error) {
	return models.Message{}, errors.New("GetMessage error")
}
func (m *mockErrorMessageService) CreateMessage(msg *models.Message) error {
	return errors.New("CreateMessage error")
}
func (m *mockErrorMessageService) UpdateMessage(id string, msg *models.Message) error {
	return errors.New("UpdateMessage error")
}
func (m *mockErrorMessageService) DeleteMessage(id string) error {
	return errors.New("DeleteMessage error")
}

func TestMessageHandler_ErrorPaths(t *testing.T) {
	app := fiber.New()
	mockService := &mockErrorMessageService{}
	handler := NewMessageHandler(mockService)

	app.Get("/messages", handler.GetMessages)
	app.Get("/messages/:id", handler.GetMessage)
	app.Post("/messages", handler.CreateMessage)
	app.Put("/messages/:id", handler.UpdateMessage)
	app.Delete("/messages/:id", handler.DeleteMessage)

	t.Run("GetMessages error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/messages", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("GetMessage error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/messages/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateMessage error", func(t *testing.T) {
		msg := &models.Message{Message: "Test Message"}
		body, _ := json.Marshal(msg)
		req := httptest.NewRequest("POST", "/messages", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("CreateMessage bad request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/messages", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("UpdateMessage error", func(t *testing.T) {
		msg := &models.Message{Message: "Test Message"}
		body, _ := json.Marshal(msg)
		req := httptest.NewRequest("PUT", "/messages/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("UpdateMessage bad request", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/messages/1", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("DeleteMessage error", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/messages/1", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", fiber.StatusInternalServerError, resp.StatusCode)
		}
	})
}
