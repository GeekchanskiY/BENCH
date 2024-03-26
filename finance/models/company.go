package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Rating      int    `json:"rating" binding:"required"`
	Description string `json:"description" binding:"required"`
	City        string `json:"city" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

func (c *Company) TableName() string {
	return "company"
}
