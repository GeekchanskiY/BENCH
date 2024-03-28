package schemas

import "Finance/models"

type EmployeeSchema struct {
	ID   uint   `json:"id"`
	Age  int    `json:"age" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (c *EmployeeSchema) ToModel(model *models.Employee) {
	model.Name = c.Name
	model.Age = c.Age
	model.ID = c.ID
}

func (c *EmployeeSchema) FromModel(model *models.Employee) {
	c.Name = model.Name
	c.Age = model.Age
	c.ID = model.ID
}
