package models

import (
	"time"

	"gorm.io/gorm"
)

type Vacancy struct {
	gorm.Model
	Name        string
	CompanyID   uint
	Company     Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VacancyLink string
	Description string
	PubDate     time.Time
	Experience  int
}

func (v *Vacancy) TableName() string {
	return "vacancy"
}

type VacancySkill struct {
	VacancyID uint
	Vacancy   Vacancy `gorm:"foreignKey:VacancyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SkillID   uint
	Skill     Skill `gorm:"foreignKey:SkillID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Priority  int
}

func (vs *VacancySkill) TableName() string {
	return "vacancy_skill"
}

type VacancyDomain struct {
	VacancyID uint
	Vacancy   Vacancy
	DomainID  uint
	Domain    Domain
	Priority  int
}

func (vd *VacancyDomain) TableName() string {
	return "vacancy_domain"
}
