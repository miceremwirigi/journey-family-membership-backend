package repository

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"gorm.io/gorm"
)

type MemberRepository interface {
	GetMembers() ([]models.Member, error)
	GetMember(id string) (models.Member, error)
	CreateMember(member *models.Member) error
	UpdateMember(id string, member *models.Member) error
	DeleteMember(id string) error
	FindByEmail(email string) (*models.Member, error)
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) GetMembers() ([]models.Member, error) {
	var members []models.Member
	err := r.db.Find(&members).Error
	return members, err
}

func (r *memberRepository) GetMember(id string) (models.Member, error) {
	var member models.Member
	err := r.db.First(&member, id).Error
	return member, err
}

func (r *memberRepository) CreateMember(member *models.Member) error {
	err := member.HashPassword(member.Password)
	if err != nil {
		return err
	}
	err = r.db.Create(&member).Error
	return err
}

func (r *memberRepository) UpdateMember(id string, member *models.Member) error {
	err := r.db.Where("id = ?", id).Updates(&member).Error
	return err
}

func (r *memberRepository) DeleteMember(id string) error {
	var member models.Member
	err := r.db.Delete(&member, id).Error
	return err
}

func (r *memberRepository) FindByEmail(email string) (*models.Member, error) {
	var member models.Member
	err := r.db.Where("email = ?", email).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}
