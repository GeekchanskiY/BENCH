package models

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

func (d *Domain) TableName() string {
	return "domain"
}

type VacancyDomain struct {
	VacancyID uint
	Vacancy   Vacancy `json:"vacancy" binding:"required"`
	DomainID  uint
	Domain    Domain `json:"domain" binding:"required"`
	Priority  int    `json:"priority" binding:"required"`
}

func (vd *VacancyDomain) TableName() string {
	return "vacancy_domain"
}

type SkillDomain struct {
	SkillID  uint
	Skill    Skill `json:"skill" binding:"required"`
	DomainID uint
	Domain   Domain `json:"domain" binding:"required"`
	Priority int    `json:"priority" binding:"required"`
}

func (sd *SkillDomain) TableName() string {
	return "skill_domain"
}
