package Models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// _ "github.com/mattn/go-sqlite3"
)

type Person struct {
	gorm.Model
	Name    string
	Surname string
	Email   string
}

func init_model() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Person{})
	db.Create(&Person{Name: "Dmitry", Surname: "Astrovskiy", Email: "test@gmail.com"})
}
