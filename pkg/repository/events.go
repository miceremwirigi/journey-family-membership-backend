package repository

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"gorm.io/gorm"
)

type EventRepository interface {
	GetEvents() ([]models.Event, error)
	GetEvent(id string) (models.Event, error)
	CreateEvent(event *models.Event) error
	UpdateEvent(id string, event *models.Event) error
	DeleteEvent(id string) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) GetEvents() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) GetEvent(id string) (models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	return event, err
}

func (r *eventRepository) CreateEvent(event *models.Event) error {
	err := r.db.Create(&event).Error
	return err
}

func (r *eventRepository) UpdateEvent(id string, event *models.Event) error {
	err := r.db.Where("id = ?", id).Updates(&event).Error
	return err
}

func (r *eventRepository) DeleteEvent(id string) error {
	var event models.Event
	err := r.db.Delete(&event, id).Error
	return err
}
