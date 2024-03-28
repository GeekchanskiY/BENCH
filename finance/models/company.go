package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string
	Rating      int
	Description string
	City        string
	Link        string
}

func (c *Company) TableName() string {
	return "company"
}
