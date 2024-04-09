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
	FindSkillDependencies(skill_id uint) ([]schemas.SkillSchema, error)
	CreateSkillDependency(skillDependency schemas.SkillDependencySchema) (schemas.SkillDependencySchema, error)
	DeleteSkillDependency(skillDependency schemas.SkillDependencySchema) error
	FindAllDependency() ([]schemas.SkillDependencySchema, error)

	// Skill conflict
	FindSkillConflicts(skill_id uint) ([]schemas.SkillConflictSchema, error)
	CreateSkillConflict(skillConflict schemas.SkillConflictSchema) (schemas.SkillConflictSchema, error)
	DeleteSkillConflict(skillConflict schemas.SkillConflictSchema) error
	FindAllSkillConflicts() ([]schemas.SkillConflictSchema, error)

	// Skill domain
	FindSkillDomains(skill_id uint) ([]schemas.SkillDomainSchema, error)
	CreateSkillDomain(skillDomain schemas.SkillDomainSchema) (schemas.SkillDomainSchema, error)
	DeleteSkillDomain(skillDomain schemas.SkillDomainSchema) error
	FindAllSkillDomains() ([]schemas.SkillDomainSchema, error)
}
