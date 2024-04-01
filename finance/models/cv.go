package models

import "gorm.io/gorm"

type CV struct {
	gorm.Model
	EmployeeID uint
	Employee   Employee `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VacancyID  uint
	Vacancy    Vacancy `gorm:"foreignKey:VacancyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (c *CV) TableName() string {
	return "cv"
}

type CVResponsibility struct {
	CVID    uint
	CV      CV
	SkillID uint
	Skill   Skill `gorm:"foreignKey:SkillID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Years   int
}

func (cr *CVResponsibility) TableName() string {
	return "cv_responsibility"
}
