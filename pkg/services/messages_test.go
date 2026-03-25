package services

import (
	"errors"
	"testing"

	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

// Mock MessageRepository
type mockMessageRepository struct {
	messages map[string]models.Message
}

func newMockMessageRepository() *mockMessageRepository {
	return &mockMessageRepository{
		messages: make(map[string]models.Message),
	}
}

func (m *mockMessageRepository) GetMessages() ([]models.Message, error) {
	messages := make([]models.Message, 0, len(m.messages))
	for _, msg := range m.messages {
		messages = append(messages, msg)
	}
	return messages, nil
}

func (m *mockMessageRepository) GetMessage(id string) (models.Message, error) {
	msg, ok := m.messages[id]
	if !ok {
		return models.Message{}, errors.New("message not found")
	}
	return msg, nil
}

func (m *mockMessageRepository) CreateMessage(msg *models.Message) error {
	m.messages["1"] = *msg
	return nil
}

func (m *mockMessageRepository) UpdateMessage(id string, msg *models.Message) error {
	if _, ok := m.messages[id]; !ok {
		return errors.New("message not found")
	}
	m.messages[id] = *msg
	return nil
}

func (m *mockMessageRepository) DeleteMessage(id string) error {
	if _, ok := m.messages[id]; !ok {
		return errors.New("message not found")
	}
	delete(m.messages, id)
	return nil
}

func TestMessageService(t *testing.T) {
	mockRepo := newMockMessageRepository()
	service := NewMessageService(mockRepo)

	t.Run("CreateMessage", func(t *testing.T) {
		msg := &models.Message{Message: "Test Message"}
		err := service.CreateMessage(msg)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetMessages", func(t *testing.T) {
		_, err := service.GetMessages()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetMessage", func(t *testing.T) {
		_, err := service.GetMessage("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("UpdateMessage", func(t *testing.T) {
		msg := &models.Message{Message: "Updated Test Message"}
		err := service.UpdateMessage("1", msg)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("DeleteMessage", func(t *testing.T) {
		err := service.DeleteMessage("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
