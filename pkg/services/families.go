package services

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/repository"
)

type FamilyService interface {
	GetFamilies() ([]models.Family, error)
	GetFamily(id string) (models.Family, error)
	CreateFamily(family *models.Family) error
	UpdateFamily(id string, family *models.Family) error
	DeleteFamily(id string) error
}

type familyService struct {
	repo repository.FamilyRepository
}

func NewFamilyService(repo repository.FamilyRepository) FamilyService {
	return &familyService{repo}
}

func (s *familyService) GetFamilies() ([]models.Family, error) {
	return s.repo.GetFamilies()
}

func (s *familyService) GetFamily(id string) (models.Family, error) {
	return s.repo.GetFamily(id)
}

func (s *familyService) CreateFamily(family *models.Family) error {
	return s.repo.CreateFamily(family)
}

func (s *familyService) UpdateFamily(id string, family *models.Family) error {
	return s.repo.UpdateFamily(id, family)
}

func (s *familyService) DeleteFamily(id string) error {
	return s.repo.DeleteFamily(id)
}
