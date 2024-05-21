package handlers

import (
	"github.com/Admiral-Simo/shortly_backend/db"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userStore db.UserStorer
}

func NewAuthHandler(userStore db.UserStorer) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var loginRequest AuthRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return SendError(c, fiber.StatusBadRequest, ErrInvalidRequestBody)
	}
	username, password := loginRequest.Username, loginRequest.Password
	user, err := h.userStore.CheckUser(username, password)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, fiber.Map{"err": "incorrect username or password"})
	}
	return c.JSON(fiber.Map{"user": user})
}

func (h *AuthHandler) Signup(c *fiber.Ctx) error {
	var signupRequest AuthRequest
	if err := c.BodyParser(&signupRequest); err != nil {
		return SendError(c, fiber.StatusBadRequest, ErrInvalidRequestBody)
	}
	username, password := signupRequest.Username, signupRequest.Password
	user, err := h.userStore.CreateUser(username, password)
	if err != nil {
		return SendError(c, fiber.StatusConflict, fiber.Map{"err": "user already in signup you might want to login instead"})
	}
	return c.JSON(fiber.Map{"user": user})
}
