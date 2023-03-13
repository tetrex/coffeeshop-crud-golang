package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tetrex/coffeeshop-crud-golang/pkg/response"
)

func HomeApi(ctx *fiber.Ctx) error {
	responseObj := response.FormattedResponse(&response.FormattedResponseDto{Data: "data", Status: http.StatusOK, Error: ""})

	// Return response in JSON format
	return ctx.Status(http.StatusOK).JSON(responseObj)
}
