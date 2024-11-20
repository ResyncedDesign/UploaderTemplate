package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization") // Fetch the auth header from the request
		token := ""
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ") // Then we want to extract the token part
		}

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Here you should check if the token is valid or implement your own authentication logic
		// e.g. query the database to see who the token belongs to for example like this:

		// user := models.GetUser(token) // This is just an example, this should be implemented according to your needs
		// if user == nil {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"message": "Unauthorized",
		// 	})
		// }

		// c.Locals("user", user) // Then we could save the user in the locals so we can access it in the next handler / controller

		// Proceed to next middleware
		return c.Next()
	}
}
