package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string `json:"name"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	City        string `json:"city"`
	Link        string `json:"link"`
}

func (c *Company) TableName() string {
	return "company"
}

type Vacancy struct {
	gorm.Model
	Name        string `json:"name"`
	CompanyID   uint
	Company     *Company  `json:"company"`
	CompanyLink string    `json:"company_link"`
	VacancyLink string    `json:"vacancy_link"`
	Description string    `json:"description"`
	PubDate     time.Time `json:"publishedAt"`
	Experience  int       `json:"experience"`
}

func (v *Vacancy) TableName() string {
	return "vacancy"
}

type VacancySkill struct {
	VacancyID uint
	Vacancy   Vacancy `json:"vacancy"`
	SkillID   uint
	Skill     Skill `json:"skill"`
	Priority  int   `json:"priority"`
}

func (vs *VacancySkill) TableName() string {
	return "vacancy_skill"
}

type Skill struct {
	gorm.Model
	Name        string  `json:"name"`
	Skill       []Skill `gorm:"foreignkey:ID;null" json:"parent_skill"`
	Description string  `json:"Description"`
}

func (s *Skill) TableName() string {
	return "skill"
}

type SkillConflict struct {
	gorm.Model
	Skill1ID uint
	Skill1   Skill `json:"skill_1"`
	Skill2ID uint
	Skill2   Skill  `json:"skill_2"`
	Comment  string `json:"Comment"`
	Priority int    `json:"Priority"`
}

func (s *SkillConflict) TableName() string {
	return "skill_conflict"
}

type Domain struct {
	gorm.Model
	Name string `json:"name"`
}

func (d *Domain) TableName() string {
	return "domain"
}

type VacancyDomain struct {
	VacancyID uint
	Vacancy   Vacancy `json:"vacancy"`
	DomainID  uint
	Domain    Domain `json:"domain"`
	Priority  int    `json:"priority"`
}

func (vd *VacancyDomain) TableName() string {
	return "vacancy_domain"
}

type SkillDomain struct {
	SkillID  uint
	Skill    Skill `json:"skill"`
	DomainID uint
	Domain   Domain `json:"domain"`
	Priority int    `json:"priority"`
}

func (sd *SkillDomain) TableName() string {
	return "skill_domain"
}

type CV struct {
	gorm.Model
	EmployeeID uint
	Employee   Employee `json:"employee"`
	VacancyID  uint
	Vacancy    Vacancy `json:"vacancy"`
}

func (c *CV) TableName() string {
	return "cv"
}

type CVResponsibility struct {
	CVID    uint
	CV      CV `json:"CV"`
	SkillID uint
	Skill   Skill `json:"skill"`
	Years   int   `json:"years"`
}

func (cr *CVResponsibility) TableName() string {
	return "cv_responsibility"
}

type Responsibility struct {
	gorm.Model
	SkillID         uint
	Skill           Skill  `json:"skill"`
	Priority        int    `json:"priority"`
	Name            string `json:"name"`
	Comments        string `json:"comments" gorm:"null"`
	ExperienceLevel int    `json:"experience_level"`
}

func (r *Responsibility) TableName() string {
	return "responsibility"
}

type ResponsibilitySynonim struct {
	gorm.Model
	ResponsibilityID uint
	Responsibility   Responsibility `json:"responsibility"`
	Name             string         `json:"Name"`
}

func (rs *ResponsibilitySynonim) TableName() string {
	return "responsibility_synonim"
}

type ResponsibilityConflict struct {
	gorm.Model
	Responsibility1ID uint
	Responsibility1   Responsibility `json:"responsibility_1"`
	Responsibility2ID uint
	Responsibility2   Responsibility `json:"responsibility_2"`
	Priority          int            `json:"Priority"`
}

func (rc *ResponsibilityConflict) TableName() string {
	return "responsibility_conflict"
}
