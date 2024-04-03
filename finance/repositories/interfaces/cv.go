package interfaces

import (
	"Finance/schemas"
)

type CVRepository interface {
	FindAll() ([]schemas.CVSchema, error)
	FindByID(id uint) (schemas.CVSchema, error)
	Create(cv schemas.CVSchema) (schemas.CVSchema, error)
	Delete(id uint) error

	// CV Responsibility
	CreateCVResponsibility(cvResp schemas.CVResponsibilitySchema) (schemas.CVResponsibilitySchema, error)
	DeleteCVResponsibility(cvResp schemas.CVResponsibilitySchema) error
	FindAllCVResponsibilities() ([]schemas.CVResponsibilitySchema, error)

	// CV Projects
	CreateCVProject(cvProj schemas.CVProjectSchema) (schemas.CVProjectSchema, error)
	DeleteCVProject(cvProj schemas.CVProjectSchema) error
	FindAllCVProjects() ([]schemas.CVProjectSchema, error)
}
