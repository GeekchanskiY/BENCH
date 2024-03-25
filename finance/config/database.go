package config

import (
	"Finance/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}

func Setup(DB *gorm.DB) *gorm.DB {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("status: ", err)
	}
	db.AutoMigrate(&models.Budget{})
	db.AutoMigrate(&models.Employee{})
	db.AutoMigrate(&models.Company{})
	db.AutoMigrate(&models.Vacancy{})
	db.AutoMigrate(&models.Skill{})
	db.AutoMigrate(&models.VacancySkill{})
	db.AutoMigrate(&models.Domain{})
	db.AutoMigrate(&models.SkillDomain{})
	db.AutoMigrate(&models.SkillConflict{})
	db.AutoMigrate(&models.VacancyDomain{})
	db.AutoMigrate(&models.Responsibility{})
	db.AutoMigrate(&models.ResponsibilityConflict{})
	db.AutoMigrate(&models.ResponsibilitySynonim{})
	db.AutoMigrate(&models.CV{})
	db.AutoMigrate(&models.CVResponsibility{})
	return db
}
