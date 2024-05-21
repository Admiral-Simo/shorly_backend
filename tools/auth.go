package tools

import (
	"fmt"
	"time"

	"github.com/Admiral-Simo/shortly_backend/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateTokenFromUser(user *models.User) string {
	expires := time.Now().Add(time.Hour * 100)

	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"expires":  expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := "supersecretpassword"
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sign token with secret", err)
	}
	return tokenString
}
