package handlers

import (
	"github.com/Admiral-Simo/shortly_backend/db"
	"github.com/Admiral-Simo/shortly_backend/tools"
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
		return ErrBadRequest()
	}
	username, password := loginRequest.Username, loginRequest.Password
	user, err := h.userStore.CheckUser(username, password)
	if err != nil {
		// Assuming CheckUser returns an error for invalid credentials
		if err.Error() == "invalid credentials" {
			return ErrBadRequest()
		}
		return ErrInvalidCredentials()
	}
	tokenString := tools.CreateTokenFromUser(user)
	// i need to set this to the cookies
	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    tokenString,
		HTTPOnly: true,
		SameSite: "Strict",
	})
	return c.JSON(fiber.Map{"user": user})
}

func (h *AuthHandler) Signup(c *fiber.Ctx) error {
	var signupRequest AuthRequest
	if err := c.BodyParser(&signupRequest); err != nil {
		return ErrBadRequest()
	}
	username, password := signupRequest.Username, signupRequest.Password
	user, err := h.userStore.CreateUser(username, password)
	if err != nil {
		if err.Error() == "username already taken" {
			return ErrAlreadyExists("User")
		}
		return ErrInternalServerError()
	}
	tokenString := tools.CreateTokenFromUser(user)
	// i need to set this to the cookies
	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    tokenString,
		HTTPOnly: true,
		SameSite: "Strict",
	})
	return c.JSON(fiber.Map{"user": user})
}
