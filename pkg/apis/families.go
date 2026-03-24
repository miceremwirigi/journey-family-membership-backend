package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/services"
)

type FamilyHandler struct {
	service services.FamilyService
}

func NewFamilyHandler(service services.FamilyService) *FamilyHandler {
	return &FamilyHandler{service}
}

func (h *FamilyHandler) GetFamilies(c *fiber.Ctx) error {
	families, err := h.service.GetFamilies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(families)
}

func (h *FamilyHandler) GetFamily(c *fiber.Ctx) error {
	id := c.Params("id")
	family, err := h.service.GetFamily(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(family)
}

func (h *FamilyHandler) CreateFamily(c *fiber.Ctx) error {
	family := new(models.Family)
	if err := c.BodyParser(family); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.CreateFamily(family)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(family)
}

func (h *FamilyHandler) UpdateFamily(c *fiber.Ctx) error {
	id := c.Params("id")
	family := new(models.Family)
	if err := c.BodyParser(family); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.UpdateFamily(id, family)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(family)
}

func (h *FamilyHandler) DeleteFamily(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeleteFamily(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendString("Family successfully deleted")
}
