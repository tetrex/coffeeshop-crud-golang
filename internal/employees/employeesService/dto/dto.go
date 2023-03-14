package dto

import employeesschema "github.com/tetrex/coffeeshop-crud-golang/internal/employees/employeesSchema"

type Employees struct {
	Role     employeesschema.EmployeeRoles
	Name     string
	Email    string
	Password string
}
type EmployeesCreateDto Employees
