package interfaces

import (
	"Finance/schemas"
)

type SkillRepository interface {
	FindAll() ([]schemas.SkillSchema, error)
	FindByID(id uint) (schemas.SkillSchema, error)
	Create(skill schemas.SkillSchema) (schemas.SkillSchema, error)
	Delete(id uint) error
	CreateSkillDependency(skillDependency schemas.SkillDependencySchema) (schemas.SkillDependencySchema, error)
	DeleteSkillDependency(skillDependency schemas.SkillDependencySchema) error
	FindAllDependency() ([]schemas.SkillDependencySchema, error)
}
