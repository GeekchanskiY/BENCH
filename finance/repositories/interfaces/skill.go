package interfaces

import (
	"Finance/schemas"
)

type SkillRepository interface {
	FindAll() ([]schemas.SkillSchema, error)
	FindByID(id uint) (schemas.SkillSchema, error)
	Create(domain schemas.SkillSchema) (schemas.SkillSchema, error)
	Delete(id uint) error
}
