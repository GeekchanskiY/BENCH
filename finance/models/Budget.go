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

func GetAllBudgets(b *[]Budget) (err error) {
	if err = config.db.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBudget(b *Budget, id string) (err error) {
	if err := config.db.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBudget(b *Budget) (err error) {
	err = config.db.Create(b).Error
	if err != nil {
		return err
	}
	return nil
}
