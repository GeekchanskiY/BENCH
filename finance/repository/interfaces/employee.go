package interfaces

import (
	"Finance/models"
)

type EmployeeRepository interface {
	FindAll() ([]models.Employee, error)
	FindByID(id uint) (models.Employee, error)
	Create(employee models.Employee) (models.Employee, error)
	Delete(id uint) error
}
