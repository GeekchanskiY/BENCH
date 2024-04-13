package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"
	"errors"

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

//
//	VacancyDomain
//

func (c *vacancyDatabase) FindVacancyDomain(id uint) ([]schemas.VacancyDomainSchema, error) {
	var vacancy_domains []models.VacancyDomain
	err := c.DB.Where("vacancy_id = ?", id).Find(&vacancy_domains).Error
	var skill_schemas []schemas.VacancyDomainSchema
	var schema schemas.VacancyDomainSchema
	for _, s := range vacancy_domains {
		schema.FromModel(&s)
		skill_schemas = append(skill_schemas, schema)
	}
	return skill_schemas, err
}

func (c *vacancyDatabase) CreateVacancyDomain(vacancyDomain schemas.VacancyDomainSchema) (schemas.VacancyDomainSchema, error) {
	var vacancy_domain_model models.VacancyDomain = models.VacancyDomain{}
	var vacancy models.Vacancy
	var domain models.Domain

	err := c.DB.First(&vacancy, vacancyDomain.VacancyID).Error
	if err != nil {
		return vacancyDomain, errors.New("vacancy not found")
	}
	err = c.DB.First(&domain, vacancyDomain.DomainID).Error
	if err != nil {
		return vacancyDomain, errors.New("domain not found")
	}
	vacancy_domain_model.VacancyID = vacancy.ID
	vacancy_domain_model.DomainID = domain.ID
	vacancy_domain_model.Priority = vacancyDomain.Priority
	err = c.DB.Create(vacancy_domain_model).Error
	vacancyDomain.FromModel(&vacancy_domain_model)
	return vacancyDomain, err
}

func (c *vacancyDatabase) DeleteVacancyDomain(skillDomain schemas.VacancyDomainSchema) error {
	var vacancy_domain_model models.VacancyDomain = models.VacancyDomain{}
	skillDomain.ToModel(&vacancy_domain_model)

	res := c.DB.Where("vacancy_id = ? AND domain_id = ?", vacancy_domain_model.VacancyID, vacancy_domain_model.DomainID).Delete(&vacancy_domain_model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return errors.New("vacancy domain does not exists")
	}
	return nil
}

func (c *vacancyDatabase) FindAllVacancyDomain() ([]schemas.VacancyDomainSchema, error) {
	var vacancy_domains []models.VacancyDomain
	err := c.DB.Find(&vacancy_domains).Error
	var skill_schemas []schemas.VacancyDomainSchema
	var schema schemas.VacancyDomainSchema
	for _, s := range vacancy_domains {
		schema.FromModel(&s)
		skill_schemas = append(skill_schemas, schema)
	}
	return skill_schemas, err
}

//
//	VacancySkill
//

func (c *vacancyDatabase) FindVacancySkill(id uint) ([]schemas.VacancySkillSchema, error) {
	var vacancy_skills []models.VacancySkill
	err := c.DB.Where("vacancy_id = ?", id).Find(&vacancy_skills).Error
	var skill_schemas []schemas.VacancySkillSchema
	var schema schemas.VacancySkillSchema
	for _, s := range vacancy_skills {
		schema.FromModel(&s)
		skill_schemas = append(skill_schemas, schema)
	}
	return skill_schemas, err
}

func (c *vacancyDatabase) CreateVacancySkill(vacancySkill schemas.VacancySkillSchema) (schemas.VacancySkillSchema, error) {
	var vacancy_skill_model models.VacancySkill = models.VacancySkill{}
	var vacancy models.Vacancy
	var skill models.Skill

	err := c.DB.First(&vacancy, vacancySkill.VacancyID).Error
	if err != nil {
		return vacancySkill, errors.New("vacancy not found")
	}
	err = c.DB.First(&skill, vacancySkill.SkillID).Error
	if err != nil {
		return vacancySkill, errors.New("skill not found")
	}
	vacancy_skill_model.SkillID = skill.ID
	vacancy_skill_model.VacancyID = vacancy.ID
	vacancy_skill_model.Priority = vacancySkill.Priority
	err = c.DB.Create(vacancy_skill_model).Error
	vacancySkill.FromModel(&vacancy_skill_model)
	return vacancySkill, err
}

func (c *vacancyDatabase) DeleteVacancySkill(vacancySkill schemas.VacancySkillSchema) error {
	var vacancy_skill_model models.VacancySkill = models.VacancySkill{}
	vacancySkill.ToModel(&vacancy_skill_model)

	res := c.DB.Where("vacancy_id = ? AND skill_id = ?", vacancy_skill_model.VacancyID, vacancy_skill_model.SkillID).Delete(&vacancy_skill_model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return errors.New("vacancy skill does not exists")
	}
	return nil
}

func (c *vacancyDatabase) FindAllVacancySkill() ([]schemas.VacancySkillSchema, error) {
	var vacancy_skills []models.VacancySkill
	err := c.DB.Find(&vacancy_skills).Error
	var skill_schemas []schemas.VacancySkillSchema
	var schema schemas.VacancySkillSchema
	for _, s := range vacancy_skills {
		schema.FromModel(&s)
		skill_schemas = append(skill_schemas, schema)
	}
	return skill_schemas, err
}
