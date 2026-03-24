package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/auth"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/common/dto"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/config"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/repository"
)

func Login(c *fiber.Ctx, memberRepo repository.MemberRepository, cfg *config.Config) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	member, err := memberRepo.FindByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	if err := member.CheckPassword(req.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	token, err := auth.GenerateToken(member, cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

// I need to add FindByEmail to the MemberRepository
type MemberRepository interface {
	FindByEmail(email string) (*models.Member, error)
}
