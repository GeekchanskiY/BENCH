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

type CVProjectSchema struct {
	CVID        uint   `json:"cv_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Years       uint   `json:"years"`
}

func (c *CVProjectSchema) ToModel(model *models.CVProject) {
	model.CVID = c.CVID
	model.Name = c.Name
	model.Description = c.Description
	model.Years = c.Years

}

func (c *CVProjectSchema) FromModel(model *models.CVProject) {
	c.CVID = model.CVID
	c.Name = model.Name
	c.Description = model.Description
	c.Years = model.Years

}

type CVResponsibilitySchema struct {
	CVProjectID      uint `json:"cv_project_id" binding:"required"`
	ResponsibilityID uint `json:"responsibility_id" binding:"required"`
	Order            uint `json:"order"`
}

func (c *CVResponsibilitySchema) ToModel(model *models.CVResponsibility) {
	model.CVProjectID = c.CVProjectID
	model.ResponsibilityID = c.ResponsibilityID
	model.Order = c.Order

}

func (c *CVResponsibilitySchema) FromModel(model *models.CVResponsibility) {
	c.CVProjectID = model.CVProjectID
	c.ResponsibilityID = model.ResponsibilityID
	c.Order = model.Order

}
