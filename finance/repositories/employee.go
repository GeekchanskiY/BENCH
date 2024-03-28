package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"

	"gorm.io/gorm"
)

type employeeDatabase struct {
	DB *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) interfaces.EmployeeRepository {
	return &employeeDatabase{DB}
}

func (c *employeeDatabase) FindAll() ([]schemas.EmployeeSchema, error) {
	var employees []models.Employee
	err := c.DB.Find(&employees).Error
	var employee_schemas []schemas.EmployeeSchema
	var schema schemas.EmployeeSchema
	for _, e := range employees {
		schema.FromModel(&e)
		employee_schemas = append(employee_schemas, schema)
	}
	return employee_schemas, err
}

func (c *employeeDatabase) FindByID(id uint) (schemas.EmployeeSchema, error) {
	var employee models.Employee
	err := c.DB.First(&employee, id).Error
	var employee_schema schemas.EmployeeSchema = schemas.EmployeeSchema{}
	employee_schema.FromModel(&employee)
	return employee_schema, err
}

func (c *employeeDatabase) Create(employee schemas.EmployeeSchema) (schemas.EmployeeSchema, error) {
	var model models.Employee = models.Employee{}
	employee.ToModel(&model)
	err := c.DB.Save(&model).Error
	employee.FromModel(&model)
	return employee, err
}

func (c *employeeDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Employee{}, id).Error
	return err
}
