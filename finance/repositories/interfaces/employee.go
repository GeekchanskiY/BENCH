package interfaces

import (
	"Finance/schemas"
)

type EmployeeRepository interface {
	FindAll() ([]schemas.EmployeeSchema, error)
	FindByID(id uint) (schemas.EmployeeSchema, error)
	Create(employee schemas.EmployeeSchema) (schemas.EmployeeSchema, error)
	Delete(id uint) error
}
