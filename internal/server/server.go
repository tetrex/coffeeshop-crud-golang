package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tetrex/coffeeshop-crud-golang/internal/routes"
)

func FiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost, nikschaefer.tech"
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	app.Use(logger.New())

	// ---------------
	// ROUTER
	routes.Initilize(app)
	// ---------------
	//
	// ---------------

	return app
}
