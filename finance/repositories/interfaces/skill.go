package interfaces

import (
	"Finance/schemas"
)

type SkillRepository interface {
	// Skills
	FindAll() ([]schemas.SkillSchema, error)
	FindByID(id uint) (schemas.SkillSchema, error)
	Create(skill schemas.SkillSchema) (schemas.SkillSchema, error)
	Delete(id uint) error

	// Skill dependency
	CreateSkillDependency(skillDependency schemas.SkillDependencySchema) (schemas.SkillDependencySchema, error)
	DeleteSkillDependency(skillDependency schemas.SkillDependencySchema) error
	FindAllDependency() ([]schemas.SkillDependencySchema, error)

	// Skill conflict
	CreateSkillConflict(skillConflict schemas.SkillConflictSchema) (schemas.SkillConflictSchema, error)
	DeleteSkillConflict(skillConflict schemas.SkillConflictSchema) error
	FindAllSkillConflicts() ([]schemas.SkillConflictSchema, error)
}
