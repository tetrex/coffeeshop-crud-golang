package routes

import (
	"github.com/gofiber/fiber/v2"
	customerscontroller "github.com/tetrex/coffeeshop-crud-golang/internal/customers/customersController"
	employeescontroller "github.com/tetrex/coffeeshop-crud-golang/internal/employees/employeesController"
	orderscontroller "github.com/tetrex/coffeeshop-crud-golang/internal/orders/ordersController"
	productscontroller "github.com/tetrex/coffeeshop-crud-golang/internal/products/productsController"
	serverservice "github.com/tetrex/coffeeshop-crud-golang/internal/server/serverService"
)

func Initilize(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1") // /api/v1

	v1.Get("/", serverservice.Root)

	// ------------------------------
	// customers
	customerGrp := v1.Group("/customer")

	customerGrp.Get("/", customerscontroller.Test)

	// ------------------------------
	// employees
	employeesGrp := v1.Group("/employees")
	employeesGrp.Put("/", employeescontroller.CreateNewEmployee)
	employeesGrp.Get("/:id", employeescontroller.FindEmployee)
	// ------------------------------
	// orders
	orders := v1.Group("/orders")
	orders.Get("/", orderscontroller.Test)
	// ------------------------------
	// products
	products := v1.Group("/products")
	products.Get("/", productscontroller.Test)

}
