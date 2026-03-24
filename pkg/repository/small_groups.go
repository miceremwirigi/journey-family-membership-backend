package repository

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"gorm.io/gorm"
)

type SmallGroupRepository interface {
	GetSmallGroups() ([]models.SmallGroup, error)
	GetSmallGroup(id string) (models.SmallGroup, error)
	CreateSmallGroup(smallGroup *models.SmallGroup) error
	UpdateSmallGroup(id string, smallGroup *models.SmallGroup) error
	DeleteSmallGroup(id string) error
}

type smallGroupRepository struct {
	db *gorm.DB
}

func NewSmallGroupRepository(db *gorm.DB) SmallGroupRepository {
	return &smallGroupRepository{db}
}

func (r *smallGroupRepository) GetSmallGroups() ([]models.SmallGroup, error) {
	var smallGroups []models.SmallGroup
	err := r.db.Find(&smallGroups).Error
	return smallGroups, err
}

func (r *smallGroupRepository) GetSmallGroup(id string) (models.SmallGroup, error) {
	var smallGroup models.SmallGroup
	err := r.db.First(&smallGroup, id).Error
	return smallGroup, err
}

func (r *smallGroupRepository) CreateSmallGroup(smallGroup *models.SmallGroup) error {
	err := r.db.Create(&smallGroup).Error
	return err
}

func (r *smallGroupRepository) UpdateSmallGroup(id string, smallGroup *models.SmallGroup) error {
	err := r.db.Where("id = ?", id).Updates(&smallGroup).Error
	return err
}

func (r *smallGroupRepository) DeleteSmallGroup(id string) error {
	var smallGroup models.SmallGroup
	err := r.db.Delete(&smallGroup, id).Error
	return err
}
