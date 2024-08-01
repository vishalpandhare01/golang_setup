package handlers

import (
	models "note-taking-app/database"
	initializers "note-taking-app/initializer"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Body struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Published bool   `json:"published"`
}

// AddBook handles adding a new book with validation
func AddBook(c *fiber.Ctx) error {
	var body Body

	// Parse the request body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//adding in db
	book := models.Book{
		Title:     body.Title,
		Content:   body.Content,
		Category:  body.Category,
		Published: body.Published,
	}

	if err := initializers.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book added successfully",
		"book": fiber.Map{
			"book": book,
		},
	})
}

// GetAllBooks retrieves all books
func GetAllBooks(c *fiber.Ctx) error {
	var books []models.Book
	if err := initializers.DB.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

// update book
func UpdateBookByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var bookUpdateRequest models.Book

	// Parse the request body
	if err := c.BodyParser(&bookUpdateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Convert id to UUID if necessary
	var id uuid.UUID
	var err error
	if id, err = uuid.Parse(idStr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var book models.Book
	if err := initializers.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update the fields
	book.Title = bookUpdateRequest.Title
	book.Content = bookUpdateRequest.Content
	book.Category = bookUpdateRequest.Category
	book.Published = bookUpdateRequest.Published

	if err := initializers.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book updated successfully",
		"book":    book,
	})
}

// delete book
func DeleteBookByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id uuid.UUID
	var err error
	if id, err = uuid.Parse(idStr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Find the book to be deleted
	var book models.Book
	if err := initializers.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Delete the book
	if err := initializers.DB.Delete(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
