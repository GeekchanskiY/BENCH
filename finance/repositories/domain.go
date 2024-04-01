package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"

	"gorm.io/gorm"
)

type domainDatabase struct {
	DB *gorm.DB
}

func NewDomainRepository(DB *gorm.DB) interfaces.DomainRepository {
	return &domainDatabase{DB}
}

func (c *domainDatabase) FindAll() ([]schemas.DomainSchema, error) {
	var domains []models.Domain
	err := c.DB.Find(&domains).Error
	var domain_schemas []schemas.DomainSchema
	var schema schemas.DomainSchema
	for _, d := range domains {
		schema.FromModel(&d)
		domain_schemas = append(domain_schemas, schema)
	}
	return domain_schemas, err
}

func (c *domainDatabase) FindByID(id uint) (schemas.DomainSchema, error) {
	var domain models.Domain
	err := c.DB.First(&domain, id).Error
	var domain_schema schemas.DomainSchema = schemas.DomainSchema{}
	domain_schema.FromModel(&domain)
	return domain_schema, err
}

func (c *domainDatabase) Create(domain schemas.DomainSchema) (schemas.DomainSchema, error) {
	var domain_model models.Domain = models.Domain{}
	domain.ToModel(&domain_model)
	err := c.DB.Save(&domain_model).Error
	domain.FromModel(&domain_model)
	return domain, err
}

func (c *domainDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Domain{}, id).Error
	return err
}
