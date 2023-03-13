package customerscontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Test(ctx *fiber.Ctx) error {

	// Return response in JSON format
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"status": http.StatusOK, "data": "customer-controller"})
}
