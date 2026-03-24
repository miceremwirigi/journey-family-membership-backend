package services

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/repository"
)

type SmallGroupService interface {
	GetSmallGroups() ([]models.SmallGroup, error)
	GetSmallGroup(id string) (models.SmallGroup, error)
	CreateSmallGroup(smallGroup *models.SmallGroup) error
	UpdateSmallGroup(id string, smallGroup *models.SmallGroup) error
	DeleteSmallGroup(id string) error
}

type smallGroupService struct {
	repo repository.SmallGroupRepository
}

func NewSmallGroupService(repo repository.SmallGroupRepository) SmallGroupService {
	return &smallGroupService{repo}
}

func (s *smallGroupService) GetSmallGroups() ([]models.SmallGroup, error) {
	return s.repo.GetSmallGroups()
}

func (s *smallGroupService) GetSmallGroup(id string) (models.SmallGroup, error) {
	return s.repo.GetSmallGroup(id)
}

func (s *smallGroupService) CreateSmallGroup(smallGroup *models.SmallGroup) error {
	return s.repo.CreateSmallGroup(smallGroup)
}

func (s *smallGroupService) UpdateSmallGroup(id string, smallGroup *models.SmallGroup) error {
	return s.repo.UpdateSmallGroup(id, smallGroup)
}

func (s *smallGroupService) DeleteSmallGroup(id string) error {
	return s.repo.DeleteSmallGroup(id)
}
