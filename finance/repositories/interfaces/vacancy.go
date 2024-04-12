package interfaces

import (
	"Finance/schemas"
)

type VacancyRepository interface {
	FindAll() ([]schemas.VacancySchema, error)
	FindByID(id uint) (schemas.VacancySchema, error)
	Create(vacancy schemas.VacancySchema) (schemas.VacancySchema, error)
	Delete(id uint) error

	FindVacancyDomain(id uint) ([]schemas.VacancyDomainSchema, error)
	FindVacancySkill(id uint) ([]schemas.VacancySkillSchema, error)

	FindAllVacancyDomain() ([]schemas.VacancyDomainSchema, error)
	FindAllVacancySkill() ([]schemas.VacancySkillSchema, error)
	CreateVacancyDomain(vacancyDomain schemas.VacancyDomainSchema) (schemas.VacancyDomainSchema, error)
	CreateVacancySkill(vacancySkill schemas.VacancySkillSchema) (schemas.VacancySkillSchema, error)
	DeleteVacancyDomain(vacancyDomain schemas.VacancyDomainSchema) error
	DeleteVacancySkill(vacancySkill schemas.VacancySkillSchema) error
}
