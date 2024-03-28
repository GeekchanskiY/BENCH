package schemas

import (
	"Finance/models"
	"time"
)

type VacancySchema struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name" binding:"required"`
	CompanyID   uint      `json:"company_id" binding:"required"`
	VacancyLink string    `json:"vacancy_link" binding:"required"`
	Description string    `json:"description" binding:"required"`
	PubDate     time.Time `json:"published_at" binding:"required"`
	Experience  int       `json:"experience" binding:"required"`
}

func (c *VacancySchema) ToModel(model *models.Vacancy) {
	model.ID = c.ID
	model.Name = c.Name
	model.Company.ID = c.CompanyID
	model.Description = c.Description
	model.Description = c.Description
	model.PubDate = c.PubDate
	model.Experience = c.Experience
}

func (c *VacancySchema) FromModel(model *models.Vacancy) {
	c.ID = model.ID
	c.Name = model.Name
	c.CompanyID = model.Company.ID
	c.Description = model.Description
	c.Description = model.Description
	c.PubDate = model.PubDate
	c.Experience = model.Experience
}
