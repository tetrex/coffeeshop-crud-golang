package routes

import "github.com/gofiber/fiber/v2"

func FiberRouter(app *fiber.App) fiber.Router {

	api := app.Group("/api")
	v1 := api.Group("/v1") // /api/v1

	return v1
}
