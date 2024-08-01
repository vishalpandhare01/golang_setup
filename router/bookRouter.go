package routes

import (
	bookController "note-taking-app/controller"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(app *fiber.App) {
	app.Post("/addBook", bookController.AddBook)
	app.Get("/getBooks", bookController.GetAllBooks)
	app.Put("/updateBook/:id", bookController.UpdateBookByID)
	app.Delete("deleteBook/:id", bookController.DeleteBookByID)
}
