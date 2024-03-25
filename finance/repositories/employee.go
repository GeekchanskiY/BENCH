package repository

import (
	"Finance/models"
	"Finance/repositories/interfaces"

	"gorm.io/gorm"
)

type employeeDatabase struct {
	DB *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) interfaces.EmployeeRepository {
	return &employeeDatabase{DB}
}

func (c *employeeDatabase) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := c.DB.Find(&employees).Error
	return employees, err
}

func (c *employeeDatabase) FindByID(id uint) (models.Employee, error) {
	var employee models.Employee
	err := c.DB.First(&employee, id).Error
	return employee, err
}

func (c *employeeDatabase) Create(employee models.Employee) (models.Employee, error) {
	err := c.DB.Save(&employee).Error
	return employee, err
}

func (c *employeeDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Employee{}, id).Error
	return err
}
