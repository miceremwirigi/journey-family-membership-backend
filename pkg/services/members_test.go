package services

import (
    "errors"
    "testing"

    "github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
)

// Mock MemberRepository
type mockMemberRepository struct {
    members map[string]models.Member
}

func newMockMemberRepository() *mockMemberRepository {
    return &mockMemberRepository{
        members: make(map[string]models.Member),
    }
}

func (m *mockMemberRepository) GetMembers() ([]models.Member, error) {
    members := make([]models.Member, 0, len(m.members))
    for _, member := range m.members {
        members = append(members, member)
    }
    return members, nil
}

func (m *mockMemberRepository) GetMember(id string) (models.Member, error) {
    member, ok := m.members[id]
    if !ok {
        return models.Member{}, errors.New("member not found")
    }
    return member, nil
}

func (m *mockMemberRepository) CreateMember(member *models.Member) error {
    m.members["1"] = *member
    return nil
}

func (m *mockMemberRepository) UpdateMember(id string, member *models.Member) error {
    if _, ok := m.members[id]; !ok {
        return errors.New("member not found")
    }
    m.members[id] = *member
    return nil
}

func (m *mockMemberRepository) DeleteMember(id string) error {
    if _, ok := m.members[id]; !ok {
        return errors.New("member not found")
    }
    delete(m.members, id)
    return nil
}

func (m *mockMemberRepository) FindByEmail(email string) (*models.Member, error) {
	return nil, nil
}

func TestMemberService(t *testing.T) {
    mockRepo := newMockMemberRepository()
    service := NewMemberService(mockRepo)

    t.Run("CreateMember", func(t *testing.T) {
        member := &models.Member{Email: "test@example.com"}
        err := service.CreateMember(member)
        if err != nil {
            t.Errorf("unexpected error: %v", err)
        }
    })

    t.Run("GetMembers", func(t *testing.T) {
        _, err := service.GetMembers()
        if err != nil {
            t.Errorf("unexpected error: %v", err)
        }
    })

    t.Run("GetMember", func(t *testing.T) {
        _, err := service.GetMember("1")
        if err != nil {
            t.Errorf("unexpected error: %v", err)
        }
    })

    t.Run("UpdateMember", func(t *testing.T) {
        member := &models.Member{Email: "updated@example.com"}
        err := service.UpdateMember("1", member)
        if err != nil {
            t.Errorf("unexpected error: %v", err)
        }
    })

    t.Run("DeleteMember", func(t *testing.T) {
        err := service.DeleteMember("1")
        if err != nil {
            t.Errorf("unexpected error: %v", err)
        }
    })
}
