package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// AdminRoleMiddleware checks if the user has an admin role
func AdminRoleMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied. Admin role required.",
		})
	}

	return c.Next()
}

// VendorRoleMiddleware checks if the user has a vendor role
func VendorRoleMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "vendor" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied. Vendor role required.",
		})
	}

	return c.Next()
}
