package interfaces

import (
	"Finance/schemas"
)

type DomainRepository interface {
	FindAll() ([]schemas.DomainSchema, error)
	FindByID(id uint) (schemas.DomainSchema, error)
	Create(domain schemas.DomainSchema) (schemas.DomainSchema, error)
	Delete(id uint) error
}
