package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"

	"gorm.io/gorm"
)

type companyDatabase struct {
	DB *gorm.DB
}

func NewCompanyRepository(DB *gorm.DB) interfaces.CompanyRepository {
	return &companyDatabase{DB}
}

func (c *companyDatabase) FindAll() ([]models.Company, error) {
	var companies []models.Company
	err := c.DB.Find(&companies).Error
	return companies, err
}

func (c *companyDatabase) FindByID(id uint) (models.Company, error) {
	var company models.Company
	err := c.DB.First(&company, id).Error
	return company, err
}

func (c *companyDatabase) Create(company models.Company) (models.Company, error) {
	err := c.DB.Save(&company).Error
	return company, err
}

func (c *companyDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Company{}, id).Error
	return err
}
