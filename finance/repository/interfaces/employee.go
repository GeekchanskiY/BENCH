package interfaces

import (
	"Finance/models"
	"context"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context) ([]models.Employee, error)
	FindByID(ctx context.Context, id uint) (models.Employee, error)
	Create(ctx context.Context, employee models.Employee) (models.Employee, error)
	Delete(ctx context.Context, employee models.Employee) error
}
