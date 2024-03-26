package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"

	"gorm.io/gorm"
)

type companyDatabase struct {
	DB *gorm.DB
}

func NewCompanyRepository(DB *gorm.DB) interfaces.CompanyRepository {
	return &companyDatabase{DB}
}

func (c *companyDatabase) FindAll() ([]schemas.CompanySchema, error) {
	var companies []models.Company
	err := c.DB.Find(&companies).Error
	var company_schemas []schemas.CompanySchema
	var schema schemas.CompanySchema
	for _, c := range companies {
		schema.FromModel(&c)
		company_schemas = append(company_schemas, schema)
	}
	return company_schemas, err
}

func (c *companyDatabase) FindByID(id uint) (schemas.CompanySchema, error) {
	var company models.Company
	err := c.DB.First(&company, id).Error
	var company_schema schemas.CompanySchema = schemas.CompanySchema{}
	company_schema.FromModel(&company)
	return company_schema, err
}

func (c *companyDatabase) Create(company schemas.CompanySchema) (schemas.CompanySchema, error) {
	var company_model models.Company = models.Company{}
	company.ToModel(&company_model)
	err := c.DB.Save(&company_model).Error
	return company, err
}

func (c *companyDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Company{}, id).Error
	return err
}
