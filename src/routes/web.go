package routes

import (
	"go-template/src/controllers"
	"go-template/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupWebRoutes(app *fiber.App) {
	app.Post("/upload", middleware.AuthMiddleware(), controllers.HandleUpload)
}