package repositories

import (
	"Finance/models"
	"Finance/repositories/interfaces"
	"Finance/schemas"
	"errors"

	"gorm.io/gorm"
)

type cvDatabase struct {
	DB *gorm.DB
}

func NewCVRepository(DB *gorm.DB) interfaces.CVRepository {
	return &cvDatabase{DB}
}

func (c *cvDatabase) FindAll() ([]schemas.CVSchema, error) {
	var cvs []models.CV
	var cv_schemas []schemas.CVSchema
	err := c.DB.Find(&cvs).Error
	if err != nil {
		return cv_schemas, err
	}

	var schema schemas.CVSchema
	for _, s := range cvs {
		schema.FromModel(&s)
		cv_schemas = append(cv_schemas, schema)
	}
	return cv_schemas, nil
}

func (c *cvDatabase) FindByID(id uint) (schemas.CVSchema, error) {
	var cv models.CV
	err := c.DB.First(&cv, id).Error
	var cv_schema schemas.CVSchema = schemas.CVSchema{}
	if err != nil {
		return cv_schema, err
	}
	cv_schema.FromModel(&cv)
	return cv_schema, err
}

func (c *cvDatabase) Create(cv schemas.CVSchema) (schemas.CVSchema, error) {
	var cv_model models.CV = models.CV{}
	cv.ToModel(&cv_model)
	err := c.DB.Save(&cv).Error
	if err != nil {
		return cv, err
	}
	cv.FromModel(&cv_model)
	return cv, err
}

func (c *cvDatabase) Delete(id uint) error {

	err := c.DB.Delete(&models.CV{}, id).Error
	return err
}

//
// CV Project
//

func (c *cvDatabase) FindAllCVProjects() ([]schemas.CVProjectSchema, error) {
	var projects []models.CVProject
	var project_schemas []schemas.CVProjectSchema
	err := c.DB.Find(&projects).Error
	if err != nil {
		return project_schemas, err
	}

	var schema schemas.CVProjectSchema
	for _, s := range projects {
		schema.FromModel(&s)
		project_schemas = append(project_schemas, schema)
	}
	return project_schemas, err
}

func (c *cvDatabase) CreateCVProject(project schemas.CVProjectSchema) (schemas.CVProjectSchema, error) {
	var project_model models.CVProject = models.CVProject{}
	var cv models.CV

	err := c.DB.First(&cv, project.CVID).Error
	if err != nil {
		return project, errors.New("cv not found")
	}

	project_model.CVID = cv.ID
	project_model.Name = project.Name
	project_model.Description = project.Description
	project_model.Years = project.Years
	err = c.DB.Create(project_model).Error
	if err != nil {
		return project, err
	}
	project.FromModel(&project_model)
	return project, err
}

func (c *cvDatabase) DeleteCVProject(project schemas.CVProjectSchema) error {
	var project_model models.CVProject = models.CVProject{}
	project.ToModel(&project_model)

	res := c.DB.Where("cv_id = ?", project.CVID).Delete(&project_model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return errors.New("cv project does not exists")
	}
	return nil
}

//
// CV Responsibility
//

func (c *cvDatabase) CreateCVResponsibility(cvResp schemas.CVResponsibilitySchema) (schemas.CVResponsibilitySchema, error) {
	var cv_resp_model models.CVResponsibility = models.CVResponsibility{}
	var project models.CVProject
	var resp models.Responsibility

	err := c.DB.First(&project, cvResp.CVProjectID).Error
	if err != nil {
		return cvResp, errors.New("cv_project not found")
	}
	err = c.DB.First(&resp, cvResp.ResponsibilityID).Error
	if err != nil {
		return cvResp, errors.New("responsibility not found")
	}
	cv_resp_model.CVProjectID = project.ID
	cv_resp_model.ResponsibilityID = resp.ID
	cv_resp_model.Order = cvResp.Order
	err = c.DB.Create(cv_resp_model).Error
	if err != nil {
		return cvResp, err
	}
	cvResp.FromModel(&cv_resp_model)
	return cvResp, err
}

func (c *cvDatabase) DeleteCVResponsibility(respConflict schemas.CVResponsibilitySchema) error {
	var resp_model models.CVResponsibility = models.CVResponsibility{}
	respConflict.ToModel(&resp_model)

	res := c.DB.Where("responsibility_id = ? AND cv_project_id = ?", resp_model.ResponsibilityID, resp_model.CVProjectID).Delete(&resp_model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return errors.New("responsibility for this project does not exists")
	}
	return nil
}

func (c *cvDatabase) FindAllCVResponsibilities() ([]schemas.CVResponsibilitySchema, error) {
	var resps []models.CVResponsibility
	var resp_schemas []schemas.CVResponsibilitySchema
	err := c.DB.Find(&resps).Error
	if err != nil {
		return resp_schemas, nil
	}

	var schema schemas.CVResponsibilitySchema
	for _, s := range resps {
		schema.FromModel(&s)
		resp_schemas = append(resp_schemas, schema)
	}
	return resp_schemas, err
}
