package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/services"
)

type VisitorHandler struct {
	service services.VisitorService
}

func NewVisitorHandler(service services.VisitorService) *VisitorHandler {
	return &VisitorHandler{service}
}

func (h *VisitorHandler) GetVisitors(c *fiber.Ctx) error {
	visitors, err := h.service.GetVisitors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(visitors)
}

func (h *VisitorHandler) GetVisitor(c *fiber.Ctx) error {
	id := c.Params("id")
	visitor, err := h.service.GetVisitor(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(visitor)
}

func (h *VisitorHandler) CreateVisitor(c *fiber.Ctx) error {
	visitor := new(models.Visitor)
	if err := c.BodyParser(visitor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.CreateVisitor(visitor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(visitor)
}

func (h *VisitorHandler) UpdateVisitor(c *fiber.Ctx) error {
	id := c.Params("id")
	visitor := new(models.Visitor)
	if err := c.BodyParser(visitor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.UpdateVisitor(id, visitor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(visitor)
}

func (h *VisitorHandler) DeleteVisitor(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeleteVisitor(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendString("Visitor successfully deleted")
}
