package controllers

import (
	controller "Finance/controllers/interfaces"
	"Finance/models"
	interfaces "Finance/repositories/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

type companyController struct {
	companyRepository interfaces.CompanyRepository
}

func NewCompanyController(repo interfaces.CompanyRepository) controller.CompanyController {
	return &companyController{
		companyRepository: repo,
	}
}

func (c *companyController) FindAll(ctx *gin.Context) {
	companies, err := c.companyRepository.FindAll()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, companies)
}

func (c *companyController) FindByID(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return

	}
	var uid uint = uint(i)
	company, err := c.companyRepository.FindByID(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, company)

}

func (c *companyController) Create(ctx *gin.Context) {
	var company models.Company
	ctx.BindJSON(&company)
	company, err := c.companyRepository.Create(company)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, company)
}

func (c *companyController) Delete(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	err = c.companyRepository.Delete(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}
