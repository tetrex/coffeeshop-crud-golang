package routes

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Initilize(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1") // /api/v1

	v1.Get("/", HomeApi)
}

func HomeApi(ctx *fiber.Ctx) error {

	// Return response in JSON format
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"status": http.StatusOK, "data": time.Now().UnixMilli()})
}
