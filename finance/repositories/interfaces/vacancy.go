package interfaces

import (
	"Finance/schemas"
)

type VacancyRepository interface {
	FindAll() ([]schemas.VacancySchema, error)
	FindByID(id uint) (schemas.VacancySchema, error)
	Create(vacancy schemas.VacancySchema) (schemas.VacancySchema, error)
	Delete(id uint) error

	FindAllVacancyDomain() []schemas.VacancyDomainSchema
	FindAllVacancySkill() []schemas.VacancySkillSchema
	CreateVacancyDomain(vacancyDomain schemas.VacancyDomainSchema) (schemas.VacancyDomainSchema, error)
	CreateVacancySkill(vacancySkill schemas.VacancySkillSchema) (schemas.VacancySkillSchema, error)
	DeleteVacancyDomain(vacancyDomain schemas.VacancyDomainSchema) error
	DeleteVacancySkill(vacancySkill schemas.VacancySkillSchema) error
}
