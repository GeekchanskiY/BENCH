package models

import "gorm.io/gorm"

type CV struct {
	gorm.Model
	EmployeeID uint
	Employee   Employee `json:"employee" binding:"required"`
	VacancyID  uint
	Vacancy    Vacancy `json:"vacancy" binding:"required"`
}

func (c *CV) TableName() string {
	return "cv"
}

type CVResponsibility struct {
	CVID    uint
	CV      CV `json:"CV" binding:"required"`
	SkillID uint
	Skill   Skill `json:"skill" binding:"required"`
	Years   int   `json:"years" binding:"required"`
}

func (cr *CVResponsibility) TableName() string {
	return "cv_responsibility"
}
