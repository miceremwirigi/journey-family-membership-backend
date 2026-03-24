package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/services"
)

type MemberHandler struct {
	service services.MemberService
}

func NewMemberHandler(service services.MemberService) *MemberHandler {
	return &MemberHandler{service}
}

func (h *MemberHandler) GetMembers(c *fiber.Ctx) error {
	members, err := h.service.GetMembers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(members)
}

func (h *MemberHandler) GetMember(c *fiber.Ctx) error {
	id := c.Params("id")
	member, err := h.service.GetMember(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}

func (h *MemberHandler) CreateMember(c *fiber.Ctx) error {
	member := new(models.Member)
	if err := c.BodyParser(member); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.CreateMember(member)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}

func (h *MemberHandler) UpdateMember(c *fiber.Ctx) error {
	id := c.Params("id")
	member := new(models.Member)
	if err := c.BodyParser(member); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.UpdateMember(id, member)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}

func (h *MemberHandler) DeleteMember(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeleteMember(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendString("Member successfully deleted")
}
