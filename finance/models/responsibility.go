package models

import "gorm.io/gorm"

type Responsibility struct {
	gorm.Model
	SkillID         uint
	Skill           Skill  `json:"skill" binding:"required"`
	Priority        int    `json:"priority" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Comments        string `json:"comments" gorm:"null" binding:"required"`
	ExperienceLevel int    `json:"experience_level" binding:"required"`
}

func (r *Responsibility) TableName() string {
	return "responsibility"
}

type ResponsibilitySynonim struct {
	gorm.Model
	ResponsibilityID uint
	Responsibility   Responsibility `json:"responsibility" binding:"required"`
	Name             string         `json:"Name" binding:"required"`
}

func (rs *ResponsibilitySynonim) TableName() string {
	return "responsibility_synonim"
}

type ResponsibilityConflict struct {
	gorm.Model
	Responsibility1ID uint
	Responsibility1   Responsibility `json:"responsibility_1" binding:"required"`
	Responsibility2ID uint
	Responsibility2   Responsibility `json:"responsibility_2" binding:"required"`
	Priority          int            `json:"Priority" binding:"required"`
}

func (rc *ResponsibilityConflict) TableName() string {
	return "responsibility_conflict"
}
