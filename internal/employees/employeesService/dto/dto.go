package dto

import employeesschema "github.com/tetrex/coffeeshop-crud-golang/internal/employees/employeesSchema"

type Employees struct {
	Role     employeesschema.EmployeeRoles `json:"role" validate:"required"`
	Name     string                        `json:"name" validate:"required"`
	Email    string                        `json:"email" validate:"email,required"`
	Password string                        `json:"password" validate:"required"`
}
type EmployeesCreateDto Employees
