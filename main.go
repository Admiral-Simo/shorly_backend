package main

import (
	"log"
	"time"

	"github.com/Admiral-Simo/shortly_backend/db"
	"github.com/Admiral-Simo/shortly_backend/handlers"
	"github.com/Admiral-Simo/shortly_backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var config = fiber.Config{
	ErrorHandler: handlers.ErrorHandler,
}

func main() {
	app := fiber.New(config)

	database, err := gorm.Open(sqlite.Open("tinyurl.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := models.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	// Middleware for CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	var (
		userStore   = db.NewUserStore(database)
		urlStore    = db.NewUrlStore(database)
		authHandler = handlers.NewAuthHandler(userStore)
		urlHandler  = handlers.NewUrlHandler(urlStore)
	)

	app.Post("/login", authHandler.Login)
	app.Post("/signup", authHandler.Signup)
	app.Get("/get/:id", urlHandler.GetUrl)

	// protected by middleware
	app.Use(handlers.JWTAuthentication(userStore))
	app.Get("/get", urlHandler.GetUrls)
	app.Post("/save", urlHandler.SaveUrl)
	app.Get("/check_authentication", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	app.Get("/logout", func(c *fiber.Ctx) error {
		// Remove the cookie
		c.Cookie(&fiber.Cookie{
			Name:     "accessToken",              // Replace with your cookie name
			Value:    "",                         // Set an empty value
			Expires:  time.Now().Add(-time.Hour), // Expire the cookie immediately
			HTTPOnly: true,                       // Ensure the cookie is HTTP only
		})

		// Redirect or return a response as needed
		return c.SendStatus(200)
	})

	app.Use(handlers.NotFoundHandler)

	app.Listen(":8080")
}
