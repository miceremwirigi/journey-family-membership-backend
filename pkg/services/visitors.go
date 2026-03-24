package services

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/repository"
)

type VisitorService interface {
	GetVisitors() ([]models.Visitor, error)
	GetVisitor(id string) (models.Visitor, error)
	CreateVisitor(visitor *models.Visitor) error
	UpdateVisitor(id string, visitor *models.Visitor) error
	DeleteVisitor(id string) error
}

type visitorService struct {
	repo repository.VisitorRepository
}

func NewVisitorService(repo repository.VisitorRepository) VisitorService {
	return &visitorService{repo}
}

func (s *visitorService) GetVisitors() ([]models.Visitor, error) {
	return s.repo.GetVisitors()
}

func (s *visitorService) GetVisitor(id string) (models.Visitor, error) {
	return s.repo.GetVisitor(id)
}

func (s *visitorService) CreateVisitor(visitor *models.Visitor) error {
	return s.repo.CreateVisitor(visitor)
}

func (s *visitorService) UpdateVisitor(id string, visitor *models.Visitor) error {
	return s.repo.UpdateVisitor(id, visitor)
}

func (s *visitorService) DeleteVisitor(id string) error {
	return s.repo.DeleteVisitor(id)
}
