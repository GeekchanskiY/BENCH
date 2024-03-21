package models

import (
	"errors"

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

func (b *Budget) Validate() error {

	if b.Name == "" {
		return errors.New("name is required")
	}

	if b.Author == "" {
		return errors.New("author is required")
	}

	if b.CurrentAmount <= 0 {
		return errors.New("price cant be less 0")
	}

	return nil
}
