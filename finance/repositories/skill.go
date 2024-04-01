package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"

	"gorm.io/gorm"
)

type skillDatabase struct {
	DB *gorm.DB
}

func NewSkillRepository(DB *gorm.DB) interfaces.SkillRepository {
	return &skillDatabase{DB}
}

func (c *skillDatabase) FindAll() ([]schemas.SkillSchema, error) {
	var skills []models.Skill
	err := c.DB.Find(&skills).Error
	var skill_schemas []schemas.SkillSchema
	var schema schemas.SkillSchema
	for _, s := range skills {
		schema.FromModel(&s)
		skill_schemas = append(skill_schemas, schema)
	}
	return skill_schemas, err
}

func (c *skillDatabase) FindByID(id uint) (schemas.SkillSchema, error) {
	var skill models.Skill
	err := c.DB.First(&skill, id).Error
	var skill_schema schemas.SkillSchema = schemas.SkillSchema{}
	skill_schema.FromModel(&skill)
	return skill_schema, err
}

func (c *skillDatabase) Create(skill schemas.SkillSchema) (schemas.SkillSchema, error) {
	var skill_model models.Skill = models.Skill{}
	skill.ToModel(&skill_model)
	err := c.DB.Save(&skill_model).Error
	skill.FromModel(&skill_model)
	return skill, err
}

func (c *skillDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Skill{}, id).Error
	return err
}
