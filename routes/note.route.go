package routes

import (
	"note-taking-app/controller"

	"github.com/gofiber/fiber/v2"
)

func NoteApiRoutes(app fiber.Router) {
	app.Post("/notes", controller.CreateNoteHandler)
}
