package models

import (
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
