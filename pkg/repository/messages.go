package repository

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"gorm.io/gorm"
)

type MessageRepository interface {
	GetMessages() ([]models.Message, error)
	GetMessage(id string) (models.Message, error)
	CreateMessage(message *models.Message) error
	UpdateMessage(id string, message *models.Message) error
	DeleteMessage(id string) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) GetMessages() ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) GetMessage(id string) (models.Message, error) {
	var message models.Message
	err := r.db.First(&message, id).Error
	return message, err
}

func (r *messageRepository) CreateMessage(message *models.Message) error {
	err := r.db.Create(&message).Error
	return err
}

func (r *messageRepository) UpdateMessage(id string, message *models.Message) error {
	err := r.db.Where("id = ?", id).Updates(&message).Error
	return err
}

func (r *messageRepository) DeleteMessage(id string) error {
	var message models.Message
	err := r.db.Delete(&message, id).Error
	return err
}
