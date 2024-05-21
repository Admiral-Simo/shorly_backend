package handlers

import (
	"github.com/Admiral-Simo/shortly_backend/db"
	"github.com/gofiber/fiber/v2"
)

type UrlHandler struct {
	urlStore db.UrlStorer
}

func NewUrlHandler(urlStore db.UrlStorer) *UrlHandler {
	return &UrlHandler{
		urlStore: urlStore,
	}
}

func (h *UrlHandler) SaveUrl(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "to implement"})
}

func (h *UrlHandler) GetUrl(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "to implement"})
}
