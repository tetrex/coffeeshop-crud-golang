package productscontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	productsservice "github.com/tetrex/coffeeshop-crud-golang/internal/products/productsService"
	"github.com/tetrex/coffeeshop-crud-golang/internal/products/productsService/dto"
)

func CreateNewProduct(ctx *fiber.Ctx) error {

	// json parsing
	json := new(dto.CreateProduct)
	if err := ctx.BodyParser(json); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"error":  "Invalid JSON",
		})
	}

	// to create employee
	obj := dto.CreateProduct{
		Name:  json.Name,
		Price: json.Price,
		Qty:   json.Qty,
	}

	// actual service call
	_, err := productsservice.Create(&obj)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  "Invalid JSON",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   "product-created",
	})
}

func FindProduct(ctx *fiber.Ctx) error {
	// pram struct
	param := struct {
		ID string `params:"id"`
	}{}
	// parsing param
	err := ctx.ParamsParser(&param) // "{"id": 111}"
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  "Invalid param",
		})
	}
	// pass that pram to service
	// to create employee
	obj := dto.FindProduct{
		Id: param.ID,
	}
	result, err, isSucess := productsservice.Find(&obj)
	if !isSucess && err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  "Invalid param",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   result,
	})
}

func DeleteProduct(ctx *fiber.Ctx) error {
	// pram struct
	param := struct {
		ID string `params:"id"`
	}{}
	// parsing param
	err := ctx.ParamsParser(&param) // "{"id": 111}"
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  "Invalid param",
		})
	}
	// pass that pram to service
	// to create employee
	obj := dto.DeleteProduct{
		Id: param.ID,
	}
	isSucess, err := productsservice.Delete(&obj)
	if !isSucess && err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  "Invalid param",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   "deleted",
	})
}
