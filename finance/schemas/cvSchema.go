package schemas

import "Finance/models"

type CVSchema struct {
	ID         uint `json:"id"`
	EmployeeID uint `json:"employee_id" binding:"required"`
	VacancyID  uint `json:"vacancy_id" binding:"required"`
}

func (c *CVSchema) ToModel(model *models.CV) {
	model.ID = c.ID
	model.EmployeeID = c.EmployeeID
	model.VacancyID = c.VacancyID

}

func (c *CVSchema) FromModel(model *models.CV) {
	c.ID = model.ID
	c.EmployeeID = model.EmployeeID
	c.VacancyID = model.VacancyID

}
