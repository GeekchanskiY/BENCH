package interfaces

import (
	"Finance/schemas"
)

type CompanyRepository interface {
	FindAll() ([]schemas.CompanySchema, error)
	FindByID(id uint) (schemas.CompanySchema, error)
	Create(employee schemas.CompanySchema) (schemas.CompanySchema, error)
	Delete(id uint) error
}
