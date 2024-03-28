package schemas

import "Finance/models"

type DomainSchema struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}

func (c *DomainSchema) ToModel(model *models.Domain) {
	model.ID = c.ID
	model.Name = c.Name

}

func (c *DomainSchema) FromModel(model *models.Domain) {
	c.ID = model.ID
	c.Name = model.Name

}
