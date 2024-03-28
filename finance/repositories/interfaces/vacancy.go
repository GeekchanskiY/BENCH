package interfaces

import (
	"Finance/schemas"
)

type VacancyRepository interface {
	FindAll() ([]schemas.VacancySchema, error)
	FindByID(id uint) (schemas.VacancySchema, error)
	Create(vacancy schemas.VacancySchema) (schemas.VacancySchema, error)
	Delete(id uint) error
}
