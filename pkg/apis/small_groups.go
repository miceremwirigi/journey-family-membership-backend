package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/services"
)

type SmallGroupHandler struct {
	service services.SmallGroupService
}

func NewSmallGroupHandler(service services.SmallGroupService) *SmallGroupHandler {
	return &SmallGroupHandler{service}
}

func (h *SmallGroupHandler) GetSmallGroups(c *fiber.Ctx) error {
	smallGroups, err := h.service.GetSmallGroups()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(smallGroups)
}

func (h *SmallGroupHandler) GetSmallGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	smallGroup, err := h.service.GetSmallGroup(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(smallGroup)
}

func (h *SmallGroupHandler) CreateSmallGroup(c *fiber.Ctx) error {
	smallGroup := new(models.SmallGroup)
	if err := c.BodyParser(smallGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.CreateSmallGroup(smallGroup)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(smallGroup)
}

func (h *SmallGroupHandler) UpdateSmallGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	smallGroup := new(models.SmallGroup)
	if err := c.BodyParser(smallGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.UpdateSmallGroup(id, smallGroup)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(smallGroup)
}

func (h *SmallGroupHandler) DeleteSmallGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeleteSmallGroup(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendString("SmallGroup successfully deleted")
}
