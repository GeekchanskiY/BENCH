package Models

import (
	"errors"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
	// Level       int    `json:"level"`
	// Description string `json:"description"`
}

func (e *Employee) TableName() string {
	return "employee"
}
func (e *Employee) Validate() error {

	if e.Name == "" {
		return errors.New("name is required")
	}

	if e.Age < 18 || e.Age > 100 {
		return errors.New("invalid age")
	}

	return nil
}

func GetAllEmployee(e *[]Employee) (err error) {

	if err := DB.Find(e).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEmployee(e *Employee, id string) (err error) {
	if err := DB.Where("ID = ?", id).First(e).Error; err != nil {
		return err
	}
	return nil
}

func AddNewEmployee(e *Employee) (err error) {
	err = DB.Create(e).Error
	if err != nil {
		return err
	}
	return nil
}
