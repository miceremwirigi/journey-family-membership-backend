package repository

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"gorm.io/gorm"
)

type FamilyRepository interface {
	GetFamilies() ([]models.Family, error)
	GetFamily(id string) (models.Family, error)
	CreateFamily(family *models.Family) error
	UpdateFamily(id string, family *models.Family) error
	DeleteFamily(id string) error
}

type familyRepository struct {
	db *gorm.DB
}

func NewFamilyRepository(db *gorm.DB) FamilyRepository {
	return &familyRepository{db}
}

func (r *familyRepository) GetFamilies() ([]models.Family, error) {
	var families []models.Family
	err := r.db.Find(&families).Error
	return families, err
}

func (r *familyRepository) GetFamily(id string) (models.Family, error) {
	var family models.Family
	err := r.db.First(&family, id).Error
	return family, err
}

func (r *familyRepository) CreateFamily(family *models.Family) error {
	err := r.db.Create(&family).Error
	return err
}

func (r *familyRepository) UpdateFamily(id string, family *models.Family) error {
	err := r.db.Where("id = ?", id).Updates(&family).Error
	return err
}

func (r *familyRepository) DeleteFamily(id string) error {
	var family models.Family
	err := r.db.Delete(&family, id).Error
	return err
}
