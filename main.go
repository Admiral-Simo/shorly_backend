package main

import (
	"log"

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

	// protected by middleware
	app.Use(handlers.JWTAuthentication(userStore))

	app.Get("/:id", urlHandler.GetUrl)
	app.Post("/save", urlHandler.SaveUrl)

	app.Listen(":8080")
}
