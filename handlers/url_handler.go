package handlers

import "github.com/gofiber/fiber/v2"

type UrlHandler struct {
}

func NewUrlHandler() *UrlHandler {
	return &UrlHandler{}
}

func (h *UrlHandler) SaveUrl(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "to implement"})
}

func (h *UrlHandler) GetUrl(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "to implement"})
}
