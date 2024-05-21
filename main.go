package main

import (
	"log"

	"github.com/Admiral-Simo/shortly_backend/db"
	"github.com/Admiral-Simo/shortly_backend/handlers"
	"github.com/Admiral-Simo/shortly_backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	database, err := gorm.Open(sqlite.Open("tinyurl.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := models.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	var (
		userStore   = db.NewUserStore(database)
		authHandler = handlers.NewAuthHandler(userStore)
		urlHandler  = handlers.NewUrlHandler()
	)

	app.Post("/login", authHandler.Login)
	app.Post("/signup", authHandler.Signup)

	// protected by middleware
	app.Get("/:id", urlHandler.GetUrl)
	app.Post("/save", urlHandler.SaveUrl)

	app.Listen(":8080")
}
