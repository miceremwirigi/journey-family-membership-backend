package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/services"
)

type EventHandler struct {
	service services.EventService
}

func NewEventHandler(service services.EventService) *EventHandler {
	return &EventHandler{service}
}

func (h *EventHandler) GetEvents(c *fiber.Ctx) error {
	events, err := h.service.GetEvents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(events)
}

func (h *EventHandler) GetEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	event, err := h.service.GetEvent(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(event)
}

func (h *EventHandler) CreateEvent(c *fiber.Ctx) error {
	event := new(models.Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.CreateEvent(event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(event)
}

func (h *EventHandler) UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	event := new(models.Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.UpdateEvent(id, event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(event)
}

func (h *EventHandler) DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeleteEvent(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendString("Event successfully deleted")
}
