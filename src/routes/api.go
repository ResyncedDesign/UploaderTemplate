package routes

import (
	"go-template/src/controllers"
	"go-template/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupApiRoutes(app *fiber.App) {
	app.Post("/api", middleware.AuthMiddleware(), controllers.HandleUpload)
}
