package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tetrex/coffeeshop-crud-golang/internal/routes"
	"golang.org/x/sync/errgroup"
)

func FiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	// ---------------
	// ROUTER
	// ---------------
	//
	// ---------------

	var eg errgroup.Group
	eg.Go(func() error {
		return app.Listen(":8080")
	})

	return app
}

func RootApi(app *fiber.App) {
	rootRouter := routes.FiberRouter(app)

	rootRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("respond with a resource")
	})
}
