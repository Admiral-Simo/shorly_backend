package handlers

import (
	"net/http"
	"time"

	"github.com/Admiral-Simo/shortly_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(userStore db.UserStorer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the cookie named "accessToken"
		tokenString := c.Cookies("accessToken")

		claims, err := validateToken(tokenString)

		if err != nil {
			return ErrUnAuthorized()
		}

		expires, err := time.Parse(time.RFC3339, claims["expires"].(string))

		if err != nil {
			return ErrUnAuthorized()
		}

		if time.Now().After(expires) {
			return NewError(http.StatusUnauthorized, "token expired")
		}

		// userID is globally accesible in every handler
		userID := int(claims["id"].(float64))
		user, err := userStore.GetUserById(userID)
		if err != nil {
			return ErrUnAuthorized()
		}

		// Set the current authenticated user to the context value
		c.Locals("user", user)

		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnAuthorized()
		}

		secret := "supersecretpassword"

		return []byte(secret), nil
	})

	if err != nil {
		return nil, ErrUnAuthorized()
	}

	if !token.Valid {
		return nil, ErrUnAuthorized()
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, ErrUnAuthorized()
	}

	return claims, nil
}
