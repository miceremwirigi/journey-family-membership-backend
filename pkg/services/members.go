package services

import (
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/repository"
)

type MemberService interface {
	GetMembers() ([]models.Member, error)
	GetMember(id string) (models.Member, error)
	CreateMember(member *models.Member) error
	UpdateMember(id string, member *models.Member) error
	DeleteMember(id string) error
}

type memberService struct {
	repo repository.MemberRepository
}

func NewMemberService(repo repository.MemberRepository) MemberService {
	return &memberService{repo}
}

func (s *memberService) GetMembers() ([]models.Member, error) {
	return s.repo.GetMembers()
}

func (s *memberService) GetMember(id string) (models.Member, error) {
	return s.repo.GetMember(id)
}

func (s *memberService) CreateMember(member *models.Member) error {
	return s.repo.CreateMember(member)
}

func (s *memberService) UpdateMember(id string, member *models.Member) error {
	return s.repo.UpdateMember(id, member)
}

func (s *memberService) DeleteMember(id string) error {
	return s.repo.DeleteMember(id)
}
