package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Skill       []Skill `gorm:"foreignkey:ID;null" json:"parent_skill" binding:"required"`
	Description string  `json:"Description" binding:"required"`
}

func (s *Skill) TableName() string {
	return "skill"
}

type VacancySkill struct {
	VacancyID uint
	Vacancy   Vacancy `json:"vacancy" binding:"required"`
	SkillID   uint
	Skill     Skill `json:"skill" binding:"required"`
	Priority  int   `json:"priority" binding:"required"`
}

func (vs *VacancySkill) TableName() string {
	return "vacancy_skill"
}

type SkillConflict struct {
	gorm.Model
	Skill1ID uint
	Skill1   Skill `json:"skill_1" binding:"required"`
	Skill2ID uint
	Skill2   Skill  `json:"skill_2" binding:"required"`
	Comment  string `json:"Comment" binding:"required"`
	Priority int    `json:"Priority" binding:"required"`
}

func (s *SkillConflict) TableName() string {
	return "skill_conflict"
}
