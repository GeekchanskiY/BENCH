package controllers

import (
	controller "Finance/controllers/interfaces"
	interfaces "Finance/repositories/interfaces"
	"Finance/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

type domainController struct {
	domainRepository interfaces.DomainRepository
}

func NewDomainController(repo interfaces.DomainRepository) controller.DomainController {
	return &domainController{
		domainRepository: repo,
	}
}

func (c *domainController) FindAll(ctx *gin.Context) {
	companies, err := c.domainRepository.FindAll()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, companies)
}

func (c *domainController) FindByID(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return

	}
	var uid uint = uint(i)
	company, err := c.domainRepository.FindByID(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, company)

}

func (c *domainController) Create(ctx *gin.Context) {
	var domain schemas.DomainSchema
	err := ctx.BindJSON(&domain)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	domain, err = c.domainRepository.Create(domain)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, domain)
}

func (c *domainController) Delete(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	err = c.domainRepository.Delete(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}
