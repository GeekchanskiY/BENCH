package Models

import (
	"Finance/Config"
	"time"

	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	Name          string  `json:"name"`
	Author        string  `json:"author"`
	CurrentAmount float64 `json:"currentAmount"`
}

func (b *Budget) TableName() string {
	return "budget"
}

func Migrate() {
	Config.DB.AutoMigrate(&Budget{})
}

type Employee struct {
	gorm.Model
	Name string `json:"name"`
}

type Company struct {
	gorm.Model
	Name        string `json:"name"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	City        string `json:"city"`
	Link        string `json:"link"`
}

type Vacancy struct {
	gorm.Model
	Name        string    `json:"name"`
	Company     string    `json:"company"`
	CompanyLink string    `json:"company_link"`
	VacancyLink string    `json:"vacancy_link"`
	Description string    `json:"description"`
	PubDate     time.Time `json:"publishedAt"`
	Experience  int       `json:"experience"`
}

type VacancySkills struct {
	Vacancy  Vacancy `json:"vacancy"`
	Skill    Skill   `json:"skill"`
	Priority int     `json:"priority"`
}

type Skill struct {
	gorm.Model
	Name        string `json:"name"`
	Skill       *Skill `gorm:"null" json:"parent_skill"`
	Description string `json:"Description"`
}

type SkillConflict struct {
	gorm.Model
	Skill1   *Skill `json:"skill_1"`
	Skill2   *Skill `json:"skill_2"`
	Comment  string `json:"Comment"`
	Priority int    `json:"Priority"`
}

type Domain struct {
	gorm.Model
	Name string `json:"name"`
}

type VacancyDomain struct {
	Vacancy  Vacancy `json:"vacancy"`
	Domain   Domain  `json:"domain"`
	Priority int     `json:"priority"`
}

type SkillDomain struct {
	Skill    *Skill  `json:"skill"`
	Domain   *Domain `json:"domain"`
	Priority int     `json:"priority"`
}

type CV struct {
	gorm.Model
	Employee *Employee `json:"employee"`
	Vacancy  *Vacancy  `json:"vacancy"`
}

type CVSkills struct {
	CV    *CV    `json:"CV"`
	Skill *Skill `json:"skill"`
	Years int    `json:"years"`
}

type Responsibility struct {
	gorm.Model
	Skill           *Skill `json:"skill"`
	Priority        int    `json:"priority"`
	Name            string `json:"name"`
	Comments        string `json:"comments" gorm:"null"`
	ExperienceLevel int    `json:"experience_level"`
}

type ResponsibilitySynonim struct {
	gorm.Model
	Responsibility *Responsibility `json:"responsibility"`
	Name           string          `json:"Name"`
}

type ResponsibilityConflict struct {
	gorm.Model
	Responsibility1 *Responsibility `json:"responsibility_1"`
	Responsibility2 *Responsibility `json:"responsibility_2"`
	Priority        int             `json:"Priority"`
}
