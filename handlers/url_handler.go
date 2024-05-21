package handlers

import (
	"fmt"

	"github.com/Admiral-Simo/shortly_backend/db"
	"github.com/Admiral-Simo/shortly_backend/models"
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
	user := c.Locals("user").(*models.User)
	if user == nil {
		fmt.Println("the save url is not having any user")
		ErrInternalServerError()
	}
	id := user.ID // userID
	fmt.Printf("hello you're %d\n", id)
	// i need to convert this user to a models.User
	return c.JSON(fiber.Map{"message": "to implement"})
}

func (h *UrlHandler) GetUrl(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	if user == nil {
		fmt.Println("the save url is not having any user")
		ErrInternalServerError()
	}
	id := user.ID // userID
	fmt.Printf("hello you're %d\n", id)
	// i need to convert this user to a models.User
	return c.JSON(fiber.Map{"message": "to implement"})
}
