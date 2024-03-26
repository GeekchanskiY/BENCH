package interfaces

import (
	"Finance/models"
)

type VacancyRepository interface {
	FindAll() ([]models.Vacancy, error)
	FindByID(id uint) (models.Vacancy, error)
	Create(vacancy models.Vacancy) (models.Vacancy, error)
	Delete(id uint) error
}
