package models

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	Name string
}

func (d *Domain) TableName() string {
	return "domain"
}
