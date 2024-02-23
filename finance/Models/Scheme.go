package Models

import (
	"Finance/Config"

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

func Migrate() {
	Config.DB.AutoMigrate(&Budget{})
}
