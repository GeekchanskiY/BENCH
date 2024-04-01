package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"
	"errors"

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

//
// Skill Dependency
//

func (c *skillDatabase) FindAllDependency() ([]schemas.SkillDependencySchema, error) {
	var skilldeps []models.SkillDependency
	err := c.DB.Find(&skilldeps).Error
	var skill_schemas []schemas.SkillDependencySchema
	var schema schemas.SkillDependencySchema
	for _, s := range skilldeps {
		schema.FromModel(&s)
		skill_schemas = append(skill_schemas, schema)
	}
	return skill_schemas, err
}

func (c *skillDatabase) CreateSkillDependency(skillDependency schemas.SkillDependencySchema) (schemas.SkillDependencySchema, error) {
	var skill_dependency_model models.SkillDependency = models.SkillDependency{}
	var parent_skill, child_skill models.Skill

	err := c.DB.First(&parent_skill, skillDependency.ParentSkillID).Error
	if err != nil {
		return skillDependency, errors.New("parent skill not found")
	}
	err = c.DB.First(&child_skill, skillDependency.ChildSkillID).Error
	if err != nil {
		return skillDependency, errors.New("child skill not found")
	}
	skill_dependency_model.ChildSkillID = child_skill.ID
	skill_dependency_model.ParentSkillID = parent_skill.ID
	err = c.DB.Create(skill_dependency_model).Error
	skillDependency.FromModel(&skill_dependency_model)
	return skillDependency, err
}

func (c *skillDatabase) DeleteSkillDependency(skillDependency schemas.SkillDependencySchema) error {
	var skill_dependency_model models.SkillDependency = models.SkillDependency{}
	skillDependency.ToModel(&skill_dependency_model)

	res := c.DB.Where("parent_skill_id = ? AND child_skill_id = ?", skill_dependency_model.ParentSkillID, skill_dependency_model.ChildSkillID).Delete(&skill_dependency_model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return errors.New("skill dependency does not exists")
	}
	return nil
}

//
// Skill Conflict
//

func (c *skillDatabase) CreateSkillConflict(skillConflict schemas.SkillConflictSchema) (schemas.SkillConflictSchema, error) {
	var skill_conflict_model models.SkillConflict = models.SkillConflict{}
	var skill_1, skill_2 models.Skill

	err := c.DB.First(&skill_1, skillConflict.Skill1ID).Error
	if err != nil {
		return skillConflict, errors.New("skill_1 not found")
	}
	err = c.DB.First(&skill_2, skillConflict.Skill2ID).Error
	if err != nil {
		return skillConflict, errors.New("child skill not found")
	}
	skill_conflict_model.Skill1ID = skill_1.ID
	skill_conflict_model.Skill2ID = skill_2.ID
	skill_conflict_model.Comment = skillConflict.Comment
	skill_conflict_model.Priority = skillConflict.Priority
	err = c.DB.Create(skill_conflict_model).Error
	skillConflict.FromModel(&skill_conflict_model)
	return skillConflict, err
}

func (c *skillDatabase) DeleteSkillConflict(skillConflict schemas.SkillConflictSchema) error {
	var skill_conflict_model models.SkillConflict = models.SkillConflict{}
	skillConflict.ToModel(&skill_conflict_model)

	res := c.DB.Where("skill_1_id = ? AND skill_2_id = ?", skill_conflict_model.Skill1ID, skill_conflict_model.Skill2ID).Delete(&skill_conflict_model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return errors.New("skill dependency does not exists")
	}
	return nil
}

func (c *skillDatabase) FindAllSkillConflicts() ([]schemas.SkillConflictSchema, error) {
	var skill_conflicts []models.SkillConflict
	err := c.DB.Find(&skill_conflicts).Error
	var skill_schemas []schemas.SkillConflictSchema
	var schema schemas.SkillConflictSchema
	for _, s := range skill_conflicts {
		schema.FromModel(&s)
		skill_schemas = append(skill_schemas, schema)
	}
	return skill_schemas, err
}
