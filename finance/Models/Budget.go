package Models

import (
	"Finance/Config"
)

func GetAllBudgets(b *[]Budget) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBudget(b *Budget, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}
