package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"
	"errors"

	"gorm.io/gorm"
)

type respDatabase struct {
	DB *gorm.DB
}

func NewResponsibilityRepository(DB *gorm.DB) interfaces.ResponsibilityRepository {
	return &respDatabase{DB}
}

func (c *respDatabase) FindAll() ([]schemas.ResponsibilitySchema, error) {
	var resps []models.Responsibility
	var resp_schemas []schemas.ResponsibilitySchema
	err := c.DB.Find(&resps).Error
	if err != nil {
		return resp_schemas, err
	}

	var schema schemas.ResponsibilitySchema
	for _, s := range resps {
		schema.FromModel(&s)
		resp_schemas = append(resp_schemas, schema)
	}
	return resp_schemas, nil
}

func (c *respDatabase) FindByID(id uint) (schemas.ResponsibilitySchema, error) {
	var resp models.Responsibility
	err := c.DB.First(&resp, id).Error
	var resp_schema schemas.ResponsibilitySchema = schemas.ResponsibilitySchema{}
	if err != nil {
		return resp_schema, err
	}
	resp_schema.FromModel(&resp)
	return resp_schema, err
}

func (c *respDatabase) Create(resp schemas.ResponsibilitySchema) (schemas.ResponsibilitySchema, error) {
	var resp_model models.Responsibility = models.Responsibility{}
	resp.ToModel(&resp_model)
	err := c.DB.Save(&resp_model).Error
	if err != nil {
		return resp, err
	}
	resp.FromModel(&resp_model)
	return resp, err
}

func (c *respDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.Responsibility{}, id).Error
	return err
}

//
// Responsibility synonim
//

func (c *respDatabase) FindAllResponsibilitySynonim() ([]schemas.ResponsibilitySynonimSchema, error) {
	var resp_synonims []models.ResponsibilitySynonim
	var synonim_schemas []schemas.ResponsibilitySynonimSchema
	err := c.DB.Find(&resp_synonims).Error
	if err != nil {
		return synonim_schemas, err
	}

	var schema schemas.ResponsibilitySynonimSchema
	for _, s := range resp_synonims {
		schema.FromModel(&s)
		synonim_schemas = append(synonim_schemas, schema)
	}
	return synonim_schemas, err
}

func (c *respDatabase) CreateResponsibilitySynonim(synonim schemas.ResponsibilitySynonimSchema) (schemas.ResponsibilitySynonimSchema, error) {
	var synonim_model models.ResponsibilitySynonim = models.ResponsibilitySynonim{}
	var resp models.Responsibility

	err := c.DB.First(&resp, synonim.ResponsibilityID).Error
	if err != nil {
		return synonim, errors.New("responsibility not found")
	}

	synonim_model.ResponsibilityID = resp.ID
	synonim_model.Name = synonim.Name
	err = c.DB.Create(synonim_model).Error
	if err != nil {
		return synonim, err
	}
	synonim.FromModel(&synonim_model)
	return synonim, err
}

func (c *respDatabase) DeleteResponsibilitySynonim(synonim schemas.ResponsibilitySynonimSchema) error {
	var synonim_model models.ResponsibilitySynonim = models.ResponsibilitySynonim{}
	synonim.ToModel(&synonim_model)

	res := c.DB.Where("responsibility_id = ?", synonim.ResponsibilityID).Delete(&synonim_model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return errors.New("synonim does not exists")
	}
	return nil
}

//
// Responsibility Conflict
//

func (c *respDatabase) CreateResponsibilityConflict(respConflict schemas.ResponsibilityConflictSchema) (schemas.ResponsibilityConflictSchema, error) {
	var resp_conflict_model models.ResponsibilityConflict = models.ResponsibilityConflict{}
	var resp_1, resp_2 models.Responsibility

	err := c.DB.First(&resp_1, respConflict.Responsibility1ID).Error
	if err != nil {
		return respConflict, errors.New("responsibility_1 not found")
	}
	err = c.DB.First(&resp_2, respConflict.Responsibility2ID).Error
	if err != nil {
		return respConflict, errors.New("responsibility_2 not found")
	}
	resp_conflict_model.Responsibility1ID = resp_1.ID
	resp_conflict_model.Responsibility2ID = resp_2.ID
	resp_conflict_model.Priority = respConflict.Priority
	err = c.DB.Create(resp_conflict_model).Error
	if err != nil {
		return respConflict, err
	}
	respConflict.FromModel(&resp_conflict_model)
	return respConflict, err
}

func (c *respDatabase) DeleteResponsibilityConflict(respConflict schemas.ResponsibilityConflictSchema) error {
	var resp_conflict_model models.ResponsibilityConflict = models.ResponsibilityConflict{}
	respConflict.ToModel(&resp_conflict_model)

	res := c.DB.Where("responsibility_1_id = ? AND responsibility_2_id = ?", resp_conflict_model.Responsibility1ID, resp_conflict_model.Responsibility2ID).Delete(&resp_conflict_model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return errors.New("responsibility conflict does not exists")
	}
	return nil
}

func (c *respDatabase) FindAllResponsibilityConflicts() ([]schemas.ResponsibilityConflictSchema, error) {
	var resp_conflicts []models.ResponsibilityConflict
	var resp_schemas []schemas.ResponsibilityConflictSchema
	err := c.DB.Find(&resp_conflicts).Error
	if err != nil {
		return resp_schemas, nil
	}

	var schema schemas.ResponsibilityConflictSchema
	for _, s := range resp_conflicts {
		schema.FromModel(&s)
		resp_schemas = append(resp_schemas, schema)
	}
	return resp_schemas, err
}
