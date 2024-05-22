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

type SaveRequest struct {
	Url string `json:"url"`
}

func (h *UrlHandler) SaveUrl(c *fiber.Ctx) error {
	var saveRequest SaveRequest
	if err := c.BodyParser(&saveRequest); err != nil {
		return ErrBadRequest()
	}
	user := c.Locals("user").(*models.User)
	if user == nil {
		fmt.Println("the save url is not having any user")
		return ErrInternalServerError()
	}
	userID := user.ID // userID
	url, err := h.urlStore.CreateUrl(userID, saveRequest.Url)
	if err != nil {
		return ErrInternalServerError()
	}
	// i need to convert this user to a models.User
	return c.JSON(fiber.Map{"url": url})
}

func (h *UrlHandler) GetUrls(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	if user == nil {
		fmt.Println("the save url is not having any user")
		return ErrInternalServerError()
	}
	userID := user.ID // userID
	urls, err := h.urlStore.GetUrls(userID)
	if err != nil {
		return err
	}
	// i need to convert this user to a models.User
	return c.JSON(fiber.Map{"urls": urls})
}

func (h *UrlHandler) GetUrl(c *fiber.Ctx) error {
	hash := c.Params("id")
	user := c.Locals("user").(*models.User)
	if user == nil {
		fmt.Println("the save url is not having any user")
		return ErrInternalServerError()
	}
	userID := user.ID // userID
	url, err := h.urlStore.GetUrl(userID, hash)
	if err != nil {
		return ErrInternalServerError()
	}
	// i need to convert this user to a models.User
	return c.JSON(fiber.Map{"url": url})
}
