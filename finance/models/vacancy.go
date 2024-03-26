package models

import (
	"time"

	"gorm.io/gorm"
)

type Vacancy struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	CompanyID   uint
	Company     *Company  `json:"company" binding:"required"`
	CompanyLink string    `json:"company_link" binding:"required"`
	VacancyLink string    `json:"vacancy_link" binding:"required"`
	Description string    `json:"description" binding:"required"`
	PubDate     time.Time `json:"publishedAt" binding:"required"`
	Experience  int       `json:"experience" binding:"required"`
}

func (v *Vacancy) TableName() string {
	return "vacancy"
}
