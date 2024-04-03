package controllers

import (
	controller "Finance/controllers/interfaces"
	interfaces "Finance/repositories/interfaces"
	"Finance/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cvController struct {
	cvRepository interfaces.CVRepository
}

func NewCVController(repo interfaces.CVRepository) controller.CVController {
	return &cvController{
		cvRepository: repo,
	}
}

func (c *cvController) FindAll(ctx *gin.Context) {
	skills, err := c.cvRepository.FindAll()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *cvController) FindByID(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return

	}
	var uid uint = uint(i)
	skill, err := c.cvRepository.FindByID(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skill)

}

func (c *cvController) Create(ctx *gin.Context) {
	var cv schemas.CVSchema
	err := ctx.BindJSON(&cv)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	cv, err = c.cvRepository.Create(cv)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, cv)
}

func (c *cvController) Delete(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	err = c.cvRepository.Delete(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}

//
// CV Responsibility
//

func (c *cvController) FindAllCVResponsibility(ctx *gin.Context) {
	skills, err := c.cvRepository.FindAllCVResponsibilities()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *cvController) DeleteCVResponsibility(ctx *gin.Context) {
	var cvresp schemas.CVResponsibilitySchema
	err := ctx.BindJSON(&cvresp)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	err = c.cvRepository.DeleteCVResponsibility(cvresp)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}

func (c *cvController) CreateCVResponsibility(ctx *gin.Context) {
	var cvresp schemas.CVResponsibilitySchema

	err := ctx.BindJSON(&cvresp)
	if err != nil {

		ctx.JSON(400, ctx.Error(err))
		return
	}

	cvresp, err = c.cvRepository.CreateCVResponsibility(cvresp)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, cvresp)
}

//
// CV Project
//

func (c *cvController) CreateCVProject(ctx *gin.Context) {
	var cv_proj schemas.CVProjectSchema

	err := ctx.BindJSON(&cv_proj)
	if err != nil {

		ctx.JSON(400, ctx.Error(err))
		return
	}

	cv_proj, err = c.cvRepository.CreateCVProject(cv_proj)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, cv_proj)
}

func (c *cvController) FindAllCVProject(ctx *gin.Context) {
	skills, err := c.cvRepository.FindAllCVProjects()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *cvController) DeleteCVProject(ctx *gin.Context) {
	var cv_proj schemas.CVProjectSchema
	err := ctx.BindJSON(&cv_proj)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	err = c.cvRepository.DeleteCVProject(cv_proj)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}
