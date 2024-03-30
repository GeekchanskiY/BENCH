package schemas

import "Finance/models"

type ResponsibilitySchema struct {
	SkillID         uint   `json:"skill_id" binding:"required"`
	Priority        int    `json:"priority" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Comments        string `json:"comments" binding:"required"`
	ExperienceLevel int    `json:"experience_level" binding:"required"`
}

func (c *ResponsibilitySchema) ToModel(model *models.Responsibility) {
	model.SkillID = c.SkillID
	model.Name = c.Name
	model.Priority = c.Priority
	model.Comments = c.Comments
	model.ExperienceLevel = c.ExperienceLevel

}

func (c *ResponsibilitySchema) FromModel(model *models.Responsibility) {
	c.SkillID = model.SkillID
	c.Name = model.Name
	c.Priority = model.Priority
	c.Comments = model.Comments
	c.ExperienceLevel = model.ExperienceLevel

}

type ResponsibilitySynonimSchema struct {
	ResponsibilityID uint   `json:"responsibility_id" binding:"required"`
	Name             string `json:"name_id" binding:"required"`
}

func (c *ResponsibilitySynonimSchema) ToModel(model *models.ResponsibilitySynonim) {
	model.Name = c.Name
	model.ResponsibilityID = c.ResponsibilityID
}

func (c *ResponsibilitySynonimSchema) FromModel(model *models.ResponsibilitySynonim) {
	c.Name = model.Name
	c.ResponsibilityID = model.ResponsibilityID
}

type ResponsibilityConflictSchema struct {
	Responsibility1ID uint `json:"responsibility_1_id" binding:"required"`
	Responsibility2ID uint `json:"responsibility_2_id" binding:"required"`
	Priority          int  `json:"priority" binding:"required"`
}

func (c *ResponsibilityConflictSchema) ToModel(model *models.ResponsibilityConflict) {
	model.Responsibility1ID = c.Responsibility1ID
	model.Responsibility2ID = c.Responsibility2ID
	model.Priority = c.Priority
}

func (c *ResponsibilityConflictSchema) FromModel(model *models.ResponsibilityConflict) {
	c.Responsibility1ID = model.Responsibility1ID
	c.Responsibility2ID = model.Responsibility2ID
	c.Priority = model.Priority
}
