package schemas

import (
	"Finance/models"
	"time"
)

//
// 	VacancySchema
//

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
	c.CompanyID = model.CompanyID
	c.Description = model.Description
	c.Description = model.Description
	c.PubDate = model.PubDate
	c.Experience = model.Experience
}

//
// VacancySkillSchema
//

type VacancySkillSchema struct {
	VacancyID uint `json:"vacancy_id" binding:"required"`
	SkillID   uint `json:"skill_id" binding:"required"`
	Priority  int  `json:"priority" binding:"required"`
}

func (c *VacancySkillSchema) ToModel(model *models.VacancySkill) {
	model.VacancyID = c.VacancyID
	model.SkillID = c.SkillID
	model.Priority = c.Priority

}

func (c *VacancySkillSchema) FromModel(model *models.VacancySkill) {
	c.VacancyID = model.VacancyID
	c.SkillID = model.SkillID
	c.Priority = model.Priority

}

//
//	VacancyDomain
//

type VacancyDomainSchema struct {
	VacancyID uint `json:"vacancy_id" binding:"required"`
	DomainID  uint `json:"domain_id" binding:"required"`
	Priority  int  `json:"priority" binding:"required"`
}

func (c *VacancyDomainSchema) ToModel(model *models.VacancyDomain) {
	model.VacancyID = c.VacancyID
	model.DomainID = c.DomainID
	model.Priority = c.Priority

}

func (c *VacancyDomainSchema) FromModel(model *models.VacancyDomain) {
	c.VacancyID = model.VacancyID
	c.DomainID = model.DomainID
	c.Priority = model.Priority

}
