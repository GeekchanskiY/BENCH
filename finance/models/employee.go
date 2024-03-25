package models

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
