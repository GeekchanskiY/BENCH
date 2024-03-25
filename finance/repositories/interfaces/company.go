package interfaces

import (
	"Finance/models"
)

type CompanyRepository interface {
	FindAll() ([]models.Company, error)
	FindByID(id uint) (models.Company, error)
	Create(employee models.Company) (models.Company, error)
	Delete(id uint) error
}
