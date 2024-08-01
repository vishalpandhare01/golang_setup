package main

import (
	"log"
	initializers "note-taking-app/initializer"
	routes "note-taking-app/router"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	initializers.Migrate()
	app := fiber.New()
	r := gin.Default()
	routes.SetupRoutes(app) //all routes

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "server work well",
		})
	})

	log.Fatal(app.Listen(":8000"))
	r.Run()
}
