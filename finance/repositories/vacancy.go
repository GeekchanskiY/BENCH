package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"

	"gorm.io/gorm"
)

type vacancyDatabase struct {
	DB *gorm.DB
}

func NewVacancyRepository(DB *gorm.DB) interfaces.VacancyRepository {
	return &vacancyDatabase{DB}
}

func (c *vacancyDatabase) FindAll() ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
	err := c.DB.Find(&vacancies).Error
	return vacancies, err
}

func (c *vacancyDatabase) FindByID(id uint) (models.Vacancy, error) {
	var vacancy models.Vacancy
	err := c.DB.First(&vacancy, id).Error
	return vacancy, err
}

func (c *vacancyDatabase) Create(vacancy models.Vacancy) (models.Vacancy, error) {
	err := c.DB.Save(&vacancy).Error
	return vacancy, err
}

func (c *vacancyDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Vacancy{}, id).Error
	return err
}
