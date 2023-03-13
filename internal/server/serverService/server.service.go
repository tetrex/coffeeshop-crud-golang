package serverservice

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Root(ctx *fiber.Ctx) error {

	// Return response in JSON format
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"status": http.StatusOK, "data": time.Now().UnixMilli()})
}
