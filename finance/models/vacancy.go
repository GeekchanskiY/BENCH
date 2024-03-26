package models

import (
	"time"

	"gorm.io/gorm"
)

type Vacancy struct {
	gorm.Model
	Name        string
	Company     Company `gorm:"references:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VacancyLink string
	Description string
	PubDate     time.Time
	Experience  int
}

func (v *Vacancy) TableName() string {
	return "vacancy"
}
