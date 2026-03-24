package services

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/repository"
)

type MessageService interface {
	GetMessages() ([]models.Message, error)
	GetMessage(id string) (models.Message, error)
	CreateMessage(message *models.Message) error
	UpdateMessage(id string, message *models.Message) error
	DeleteMessage(id string) error
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{repo}
}

func (s *messageService) GetMessages() ([]models.Message, error) {
	return s.repo.GetMessages()
}

func (s *messageService) GetMessage(id string) (models.Message, error) {
	return s.repo.GetMessage(id)
}

func (s *messageService) CreateMessage(message *models.Message) error {
	return s.repo.CreateMessage(message)
}

func (s *messageService) UpdateMessage(id string, message *models.Message) error {
	return s.repo.UpdateMessage(id, message)
}

func (s *messageService) DeleteMessage(id string) error {
	return s.repo.DeleteMessage(id)
}
