package repository

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"gorm.io/gorm"
)

type VisitorRepository interface {
	GetVisitors() ([]models.Visitor, error)
	GetVisitor(id string) (models.Visitor, error)
	CreateVisitor(visitor *models.Visitor) error
	UpdateVisitor(id string, visitor *models.Visitor) error
	DeleteVisitor(id string) error
}

type visitorRepository struct {
	db *gorm.DB
}

func NewVisitorRepository(db *gorm.DB) VisitorRepository {
	return &visitorRepository{db}
}

func (r *visitorRepository) GetVisitors() ([]models.Visitor, error) {
	var visitors []models.Visitor
	err := r.db.Find(&visitors).Error
	return visitors, err
}

func (r *visitorRepository) GetVisitor(id string) (models.Visitor, error) {
	var visitor models.Visitor
	err := r.db.First(&visitor, id).Error
	return visitor, err
}

func (r *visitorRepository) CreateVisitor(visitor *models.Visitor) error {
	err := r.db.Create(&visitor).Error
	return err
}

func (r *visitorRepository) UpdateVisitor(id string, visitor *models.Visitor) error {
	err := r.db.Where("id = ?", id).Updates(&visitor).Error
	return err
}

func (r *visitorRepository) DeleteVisitor(id string) error {
	var visitor models.Visitor
	err := r.db.Delete(&visitor, id).Error
	return err
}
