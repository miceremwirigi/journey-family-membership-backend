package services

import (
	"errors"
	"testing"

	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

// Mock SmallGroupRepository
type mockSmallGroupRepository struct {
	smallGroups map[string]models.SmallGroup
}

func newMockSmallGroupRepository() *mockSmallGroupRepository {
	return &mockSmallGroupRepository{
		smallGroups: make(map[string]models.SmallGroup),
	}
}

func (m *mockSmallGroupRepository) GetSmallGroups() ([]models.SmallGroup, error) {
	sgs := make([]models.SmallGroup, 0, len(m.smallGroups))
	for _, sg := range m.smallGroups {
		sgs = append(sgs, sg)
	}
	return sgs, nil
}

func (m *mockSmallGroupRepository) GetSmallGroup(id string) (models.SmallGroup, error) {
	sg, ok := m.smallGroups[id]
	if !ok {
		return models.SmallGroup{}, errors.New("small group not found")
	}
	return sg, nil
}

func (m *mockSmallGroupRepository) CreateSmallGroup(sg *models.SmallGroup) error {
	m.smallGroups["1"] = *sg
	return nil
}

func (m *mockSmallGroupRepository) UpdateSmallGroup(id string, sg *models.SmallGroup) error {
	if _, ok := m.smallGroups[id]; !ok {
		return errors.New("small group not found")
	}
	m.smallGroups[id] = *sg
	return nil
}

func (m *mockSmallGroupRepository) DeleteSmallGroup(id string) error {
	if _, ok := m.smallGroups[id]; !ok {
		return errors.New("small group not found")
	}
	delete(m.smallGroups, id)
	return nil
}

func TestSmallGroupService(t *testing.T) {
	mockRepo := newMockSmallGroupRepository()
	service := NewSmallGroupService(mockRepo)

	t.Run("CreateSmallGroup", func(t *testing.T) {
		sg := &models.SmallGroup{Name: "Test Small Group"}
		err := service.CreateSmallGroup(sg)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetSmallGroups", func(t *testing.T) {
		_, err := service.GetSmallGroups()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetSmallGroup", func(t *testing.T) {
		_, err := service.GetSmallGroup("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("UpdateSmallGroup", func(t *testing.T) {
		sg := &models.SmallGroup{Name: "Updated Test Small Group"}
		err := service.UpdateSmallGroup("1", sg)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("DeleteSmallGroup", func(t *testing.T) {
		err := service.DeleteSmallGroup("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
