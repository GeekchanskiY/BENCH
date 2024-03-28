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
