package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"

	"gorm.io/gorm"
)

type vacancyDatabase struct {
	DB *gorm.DB
}

func NewVacancyRepository(DB *gorm.DB) interfaces.VacancyRepository {
	return &vacancyDatabase{DB}
}

func (c *vacancyDatabase) FindAll() ([]schemas.VacancySchema, error) {
	var vacancies []models.Vacancy
	err := c.DB.Find(&vacancies).Error
	var vacancy_schemas []schemas.VacancySchema
	var schema schemas.VacancySchema
	for _, e := range vacancies {
		schema.FromModel(&e)
		vacancy_schemas = append(vacancy_schemas, schema)
	}
	return vacancy_schemas, err
}

func (c *vacancyDatabase) FindByID(id uint) (schemas.VacancySchema, error) {
	var vacancy models.Vacancy
	err := c.DB.First(&vacancy, id).Error
	var vacancy_schema schemas.VacancySchema = schemas.VacancySchema{}
	vacancy_schema.FromModel(&vacancy)
	return vacancy_schema, err
}

func (c *vacancyDatabase) Create(vacancy schemas.VacancySchema) (schemas.VacancySchema, error) {
	var model models.Vacancy = models.Vacancy{}
	var company models.Company
	err := c.DB.First(&company, vacancy.CompanyID).Error
	if err != nil {
		return vacancy, err
	}

	vacancy.ToModel(&model)
	model.Company = company
	err = c.DB.Save(&model).Error
	if err != nil {
		return vacancy, err
	}
	vacancy.FromModel(&model)
	return vacancy, err
}

func (c *vacancyDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Vacancy{}, id).Error
	return err
}
