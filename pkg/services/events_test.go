package services

import (
	"errors"
	"testing"

	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

// Mock EventRepository
type mockEventRepository struct {
	events map[string]models.Event
}

func newMockEventRepository() *mockEventRepository {
	return &mockEventRepository{
		events: make(map[string]models.Event),
	}
}

func (m *mockEventRepository) GetEvents() ([]models.Event, error) {
	events := make([]models.Event, 0, len(m.events))
	for _, e := range m.events {
		events = append(events, e)
	}
	return events, nil
}

func (m *mockEventRepository) GetEvent(id string) (models.Event, error) {
	e, ok := m.events[id]
	if !ok {
		return models.Event{}, errors.New("event not found")
	}
	return e, nil
}

func (m *mockEventRepository) CreateEvent(e *models.Event) error {
	m.events["1"] = *e
	return nil
}

func (m *mockEventRepository) UpdateEvent(id string, e *models.Event) error {
	if _, ok := m.events[id]; !ok {
		return errors.New("event not found")
	}
	m.events[id] = *e
	return nil
}

func (m *mockEventRepository) DeleteEvent(id string) error {
	if _, ok := m.events[id]; !ok {
		return errors.New("event not found")
	}
	delete(m.events, id)
	return nil
}

func TestEventService(t *testing.T) {
	mockRepo := newMockEventRepository()
	service := NewEventService(mockRepo)

	t.Run("CreateEvent", func(t *testing.T) {
		e := &models.Event{Title: "Test Event"}
		err := service.CreateEvent(e)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetEvents", func(t *testing.T) {
		_, err := service.GetEvents()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetEvent", func(t *testing.T) {
		_, err := service.GetEvent("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("UpdateEvent", func(t *testing.T) {
		e := &models.Event{Title: "Updated Test Event"}
		err := service.UpdateEvent("1", e)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("DeleteEvent", func(t *testing.T) {
		err := service.DeleteEvent("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
