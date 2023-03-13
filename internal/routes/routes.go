package routes

import (
	"github.com/gofiber/fiber/v2"
	serverservice "github.com/tetrex/coffeeshop-crud-golang/internal/server/serverService"
)

func Initilize(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1") // /api/v1

	v1.Get("/", serverservice.Root)
}
