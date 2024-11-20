package main

import (
	"fmt"
	"go-template/src/config"
	"go-template/src/routes"
	"go-template/src/services"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	// Call the load func from the config package (honestly we could just call godotenv.Load() here but this is more organized I guess?)
	config.Load()

	// https://docs.gofiber.io/api/fiber for more information on the config options
	app := fiber.New(fiber.Config{
		Prefork:               true, // Depending on what you add later on, you might need to disable this (e.g. database connections)
		CaseSensitive:         true, // When enabled, /Foo and /foo are different routes. When disabled, /Fooand /foo are treated the same.
		StrictRouting:         true, // When enabled, the router treats /foo and /foo/ as different. Otherwise, the router treats /foo and /foo/ as the same.
		DisableStartupMessage: true, // Shhh 🤫
		ServerHeader:          "Go Webserver",
		AppName:               "Go Webserver with R2 integration",
		BodyLimit:             100 * 1024 * 1024, // 100MB limit (increase if needed)
	})

	app.Use(limiter.New(limiter.Config{
		Next:       nil,
		Max:        25,              // 25 requests (If you only use this for uploads, depending on your case you might want to decrease this so no one can spam your server)
		Expiration: 1 * time.Minute, // Per 1 minute
		LimitReached: func(c *fiber.Ctx) error {
			// Here you could perform more actions then just returning a status like logging the IP or something
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	}))

	R2Service, err := services.NewR2Service()
	if err != nil {
		log.Fatal(err)
	}

	// Store the R2Service in the context so we can easily access it in the controllers later on
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("R2Service", R2Service)
		return c.Next()
	})

	// Setup the routes
	routes.SetupWebRoutes(app)
	routes.SetupApiRoutes(app)

	port := config.GetServerPort()
	err = app.Listen(port)
	if err != nil {
		fmt.Println("Error starting web server:", err)
	}
}
