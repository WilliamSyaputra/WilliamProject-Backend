package main

// @title William Project API
// @version 1.0
// @description William Project API Documentation
// @host williamproject-backend-production.up.railway.app
// @BasePath /api/v1
// @schemes https

import (
	"log"
	"os"

	"william/backend/controllers"
	"william/backend/docs" // Import untuk Swagger

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	docs.SwaggerInfo.Schemes = []string{"https"}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))

	// Redirect root ("/") to Swagger UI
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", 301)
	})

	// Endpoint Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/swagger.html", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", 301)
	})

	api := app.Group("/api/v1")

	// Image routes
	api.Get("/images", controllers.GetImages)
	api.Get("/images/:name", controllers.GetImage)
	api.Post("/upload", controllers.UploadImage)
	api.Post("/uploads", controllers.UploadImages)
	api.Delete("/images/:name", controllers.DeleteImage)

	// Buat folder jika belum ada
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		os.Mkdir("images", os.ModePerm)
	}

	log.Fatal(app.Listen("0.0.0.0:3001"))
}
