package services

import (
	"errors"
	"testing"

	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

// Mock FamilyRepository
type mockFamilyRepository struct {
	families map[string]models.Family
}

func newMockFamilyRepository() *mockFamilyRepository {
	return &mockFamilyRepository{
		families: make(map[string]models.Family),
	}
}

func (m *mockFamilyRepository) GetFamilies() ([]models.Family, error) {
	families := make([]models.Family, 0, len(m.families))
	for _, family := range m.families {
		families = append(families, family)
	}
	return families, nil
}

func (m *mockFamilyRepository) GetFamily(id string) (models.Family, error) {
	family, ok := m.families[id]
	if !ok {
		return models.Family{}, errors.New("family not found")
	}
	return family, nil
}

func (m *mockFamilyRepository) CreateFamily(family *models.Family) error {
	m.families["1"] = *family
	return nil
}

func (m *mockFamilyRepository) UpdateFamily(id string, family *models.Family) error {
	if _, ok := m.families[id]; !ok {
		return errors.New("family not found")
	}
	m.families[id] = *family
	return nil
}

func (m *mockFamilyRepository) DeleteFamily(id string) error {
	if _, ok := m.families[id]; !ok {
		return errors.New("family not found")
	}
	delete(m.families, id)
	return nil
}

func TestFamilyService(t *testing.T) {
	mockRepo := newMockFamilyRepository()
	service := NewFamilyService(mockRepo)

	t.Run("CreateFamily", func(t *testing.T) {
		family := &models.Family{Name: "The Test Family"}
		err := service.CreateFamily(family)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetFamilies", func(t *testing.T) {
		_, err := service.GetFamilies()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("GetFamily", func(t *testing.T) {
		_, err := service.GetFamily("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("UpdateFamily", func(t *testing.T) {
		family := &models.Family{Name: "The Updated Test Family"}
		err := service.UpdateFamily("1", family)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("DeleteFamily", func(t *testing.T) {
		err := service.DeleteFamily("1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
