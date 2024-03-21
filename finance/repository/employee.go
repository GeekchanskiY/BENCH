package repository

import (
	"Finance/models"
	"Finance/repository/interfaces"
	"context"

	"gorm.io/gorm"
)

type employeeDatabase struct {
	DB *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) interfaces.EmployeeRepository {
	return &employeeDatabase{DB}
}

func (c *employeeDatabase) FindAll(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee
	err := c.DB.Find(&employees).Error
	return employees, err
}

func (c *employeeDatabase) FindByID(ctx context.Context, id uint) (models.Employee, error) {
	var employee models.Employee
	err := c.DB.First(&employee, id).Error
	return employee, err
}

func (c *employeeDatabase) Create(ctx context.Context, employee models.Employee) (models.Employee, error) {
	err := c.DB.Save(&employee).Error
	return employee, err
}

func (c *employeeDatabase) Delete(ctx context.Context, employee models.Employee) error {
	err := c.DB.Delete(&employee).Error
	return err
}
