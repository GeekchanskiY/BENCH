package schemas

import "Finance/models"

type CompanySchema struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Rating      int    `json:"rating" binding:"required"`
	Description string `json:"description" binding:"required"`
	City        string `json:"city" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

func (c *CompanySchema) ToModel(model *models.Company) {
	model.ID = c.ID
	model.Name = c.Name
	model.City = c.City
	model.Description = c.Description
	model.Rating = c.Rating
	model.Link = c.Link
}

func (c *CompanySchema) FromModel(model *models.Company) {
	c.ID = model.ID
	c.Name = model.Name
	c.City = model.City
	c.Rating = model.Rating
	c.Description = model.Description
	c.Link = model.Link
}
