package interfaces

import (
	"Finance/schemas"
)

type ResponsibilityRepository interface {
	// Responsibilities
	FindAll() ([]schemas.ResponsibilitySchema, error)
	FindByID(id uint) (schemas.ResponsibilitySchema, error)
	Create(responsibility schemas.ResponsibilitySchema) (schemas.ResponsibilitySchema, error)
	Delete(id uint) error

	// Responsibility Synonims
	FindResponsibilitySynonim(id uint) ([]schemas.ResponsibilitySynonimSchema, error)
	CreateResponsibilitySynonim(respSynonim schemas.ResponsibilitySynonimSchema) (schemas.ResponsibilitySynonimSchema, error)
	DeleteResponsibilitySynonim(respSynonim schemas.ResponsibilitySynonimSchema) error
	FindAllResponsibilitySynonim() ([]schemas.ResponsibilitySynonimSchema, error)

	// Responsibility Conflicts
	FindResponsibilityConflict(id uint) ([]schemas.ResponsibilityConflictSchema, error)
	CreateResponsibilityConflict(respConflict schemas.ResponsibilityConflictSchema) (schemas.ResponsibilityConflictSchema, error)
	DeleteResponsibilityConflict(respConflict schemas.ResponsibilityConflictSchema) error
	FindAllResponsibilityConflicts() ([]schemas.ResponsibilityConflictSchema, error)
}
