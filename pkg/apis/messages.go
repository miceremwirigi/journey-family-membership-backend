package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/models"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/services"
)

type MessageHandler struct {
	service services.MessageService
}

func NewMessageHandler(service services.MessageService) *MessageHandler {
	return &MessageHandler{service}
}

func (h *MessageHandler) GetMessages(c *fiber.Ctx) error {
	messages, err := h.service.GetMessages()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(messages)
}

func (h *MessageHandler) GetMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	message, err := h.service.GetMessage(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(message)
}

func (h *MessageHandler) CreateMessage(c *fiber.Ctx) error {
	message := new(models.Message)
	if err := c.BodyParser(message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.CreateMessage(message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(message)
}

func (h *MessageHandler) UpdateMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	message := new(models.Message)
	if err := c.BodyParser(message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.service.UpdateMessage(id, message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(message)
}

func (h *MessageHandler) DeleteMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.DeleteMessage(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendString("Message successfully deleted")
}
