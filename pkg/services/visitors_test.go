package services

import (
	"errors"
	"testing"

	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

// Mock VisitorRepository
type mockVisitorRepository struct {
	visitors map[string]models.Visitor
}

func newMockVisitorRepository() *mockVisitorRepository {
	return &mockVisitorRepository{
		visitors: make(map[string]models.Visitor),
	}
}

func (m *mockVisitorRepository) GetVisitors() ([]models.Visitor, error) {
	visitors := make([]models.Visitor, 0, len(m.visitors))
	for _, v := range m.visitors {
		visitors = append(visitors, v)
	}
	return visitors, nil
}

func (m *mockVisitorRepository) GetVisitor(id string) (models.Visitor, error) {
	v, ok := m.visitors[id]
	if !ok {
		return models.Visitor{}, errors.New("visitor not found")
	}
	return v, nil
}

func (m *mockVisitorRepository) CreateVisitor(v *models.Visitor) error {
	m.visitors["1"] = *v
	return nil
}

func (m *mockVisitorRepository) UpdateVisitor(id string, v *models.Visitor) error {
	if _, ok := m.visitors[id]; !ok {
		return errors.New("visitor not found")
	}
	m.visitors[id] = *v
	return nil
}

func (m *mockVisitorRepository) DeleteVisitor(id string) error {
	if _, ok := m.visitors[id]; !ok {
		return errors.New("visitor not found")
	}
	delete(m.visitors, id)
	return nil
}

func TestVisitorService(t *testing.T) {
	mockRepo := newMockVisitorRepository()
	service := NewVisitorService(mockRepo)

	t.Run("CreateVisitor", func(t *testing.T) {
		v := &models.Visitor{Email: "visitor@example.com"}
		err := service.CreateVisitor(v)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetVisitors", func(t *testing.T) {
		_, err := service.GetVisitors()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetVisitor", func(t *testing.T) {
		_, err := service.GetVisitor("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("UpdateVisitor", func(t *testing.T) {
		v := &models.Visitor{Email: "updatedvisitor@example.com"}
		err := service.UpdateVisitor("1", v)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("DeleteVisitor", func(t *testing.T) {
		err := service.DeleteVisitor("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
