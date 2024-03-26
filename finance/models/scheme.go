package models

import (
	"gorm.io/gorm"
)

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

type Skill struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Skill       []Skill `gorm:"foreignkey:ID;null" json:"parent_skill" binding:"required"`
	Description string  `json:"Description" binding:"required"`
}

func (s *Skill) TableName() string {
	return "skill"
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
