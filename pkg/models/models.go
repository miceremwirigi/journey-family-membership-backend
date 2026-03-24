package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	FirstName                 string
	MiddleName                string
	LastName                  string
	Email                     string `gorm:"unique"`
	Gender                    string
	Zone                      string
	MaritalStatus             string
	Mobile                    string
	Residence                 string
	Birthday                  string
	WeddingAnniversary        string
	SpecialCelebration        string
	SpecialCelebrationDescription string
	FamilyID                  uint
	FamilyRole                string
	SmallGroupID              uint
	JoinDate                  string
	Password                  string
	Role                      string
}

func (m *Member) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	m.Password = string(bytes)
	return nil
}

func (m *Member) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

type Family struct {
	gorm.Model
	Name    string
	HeadID  uint
	Members []Member
}

type SmallGroup struct {
	gorm.Model
	Name        string
	Zone        string
	Location    string
	LeaderID    uint
	Description string
	MeetingDay  string
	MeetingTime string
	Members     []Member
}

type Visitor struct {
	gorm.Model
	FirstName  string
	MiddleName string
	LastName   string
	Gender     string
	Contact    string
	Email      string `gorm:"unique"`
	Visits     int
	FirstVisit string
	LastVisit  string
	Interest   string
	Status     string
}

type Message struct {
	gorm.Model
	Date       string
	Time       string
	Type       string
	Recipients string
	Message    string
	Delivered  int
}

type Event struct {
	gorm.Model
	Title    string
	Category string
	DateTime string
	Location string
	Status   string
}
