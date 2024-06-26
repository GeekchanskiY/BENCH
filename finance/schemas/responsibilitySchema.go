package schemas

import "Finance/models"

type ResponsibilitySchema struct {
	ID              uint   `json:"id"`
	SkillID         uint   `json:"skill_id" binding:"required"`
	Priority        int    `json:"priority" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Comments        string `json:"comments" binding:"required"`
	ExperienceLevel int    `json:"experience_level" binding:"required"`
}

func (c *ResponsibilitySchema) ToModel(model *models.Responsibility) {
	model.ID = c.ID
	model.SkillID = c.SkillID
	model.Name = c.Name
	model.Priority = c.Priority
	model.Comments = c.Comments
	model.ExperienceLevel = c.ExperienceLevel

}

func (c *ResponsibilitySchema) FromModel(model *models.Responsibility) {
	c.ID = model.ID
	c.SkillID = model.SkillID
	c.Name = model.Name
	c.Priority = model.Priority
	c.Comments = model.Comments
	c.ExperienceLevel = model.ExperienceLevel

}

type ResponsibilitySynonimSchema struct {
	ID               uint   `json:"id"`
	ResponsibilityID uint   `json:"responsibility_id" binding:"required"`
	Name             string `json:"name" binding:"required"`
}

func (c *ResponsibilitySynonimSchema) ToModel(model *models.ResponsibilitySynonim) {
	model.ID = c.ID
	model.Name = c.Name
	model.ResponsibilityID = c.ResponsibilityID
}

func (c *ResponsibilitySynonimSchema) FromModel(model *models.ResponsibilitySynonim) {
	c.ID = model.ID
	c.Name = model.Name
	c.ResponsibilityID = model.ResponsibilityID
}

type ResponsibilityConflictSchema struct {
	ID                uint `json:"id"`
	Responsibility1ID uint `json:"responsibility_1_id" binding:"required"`
	Responsibility2ID uint `json:"responsibility_2_id" binding:"required"`
	Priority          int  `json:"priority" binding:"required"`
}

func (c *ResponsibilityConflictSchema) ToModel(model *models.ResponsibilityConflict) {
	model.ID = c.ID
	model.Responsibility1ID = c.Responsibility1ID
	model.Responsibility2ID = c.Responsibility2ID
	model.Priority = c.Priority
}

func (c *ResponsibilityConflictSchema) FromModel(model *models.ResponsibilityConflict) {
	c.ID = model.ID
	c.Responsibility1ID = model.Responsibility1ID
	c.Responsibility2ID = model.Responsibility2ID
	c.Priority = model.Priority
}
