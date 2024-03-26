package models

import (
	"time"

	"gorm.io/gorm"
)

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
