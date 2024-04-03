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

type CVProject struct {
	gorm.Model
	Name        string
	Description string
	Years       uint

	CVID uint
	CV   CV `gorm:"foreignKey:CVID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CVResponsibility struct {
	CVProjectID      uint
	CVProject        CVProject
	ResponsibilityID uint
	Responsibility   Responsibility `gorm:"foreignKey:ResponsibilityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Order            uint
}

func (cr *CVResponsibility) TableName() string {
	return "cv_responsibility"
}
