package Models

import (
	"errors"
)

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
	if err = DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBudget(b *Budget, id string) (err error) {
	if err := DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBudget(b *Budget) (err error) {
	err = DB.Create(b).Error
	if err != nil {
		return err
	}
	return nil
}
