package employeescontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	employeesservice "github.com/tetrex/coffeeshop-crud-golang/internal/employees/employeesService"
	"github.com/tetrex/coffeeshop-crud-golang/internal/employees/employeesService/dto"
)

func CreateNewEmployee(ctx *fiber.Ctx) error {

	// json parsing
	json := new(dto.Employees)
	if err := ctx.BodyParser(json); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"error":  "Invalid JSON",
		})
	}

	// to create employee
	obj := dto.Employees{
		Role:     json.Role,
		Name:     json.Name,
		Email:    json.Email,
		Password: json.Password,
	}

	// actual service call
	_, err := employeesservice.Create(&obj)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  "Invalid JSON",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   "employee-created",
	})
}

func FindEmployee(ctx *fiber.Ctx) error {
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
	obj := dto.FindEmployee{
		Id: param.ID,
	}
	result, err, isSucess := employeesservice.Find(&obj)
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

func DeleteEmployee(ctx *fiber.Ctx) error {
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
	obj := dto.DeleteEmployee{
		Id: param.ID,
	}
	isSucess, err := employeesservice.Delete(&obj)
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
