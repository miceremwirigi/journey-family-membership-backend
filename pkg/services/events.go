package services

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/repository"
)

type EventService interface {
	GetEvents() ([]models.Event, error)
	GetEvent(id string) (models.Event, error)
	CreateEvent(event *models.Event) error
	UpdateEvent(id string, event *models.Event) error
	DeleteEvent(id string) error
}

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) GetEvents() ([]models.Event, error) {
	return s.repo.GetEvents()
}

func (s *eventService) GetEvent(id string) (models.Event, error) {
	return s.repo.GetEvent(id)
}

func (s *eventService) CreateEvent(event *models.Event) error {
	return s.repo.CreateEvent(event)
}

func (s *eventService) UpdateEvent(id string, event *models.Event) error {
	return s.repo.UpdateEvent(id, event)
}

func (s *eventService) DeleteEvent(id string) error {
	return s.repo.DeleteEvent(id)
}
