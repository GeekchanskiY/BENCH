package controllers

import (
	controller "Finance/controllers/interfaces"
	interfaces "Finance/repositories/interfaces"
	"Finance/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

type vacancyController struct {
	vacancyRepository interfaces.VacancyRepository
}

func NewVacancyController(repo interfaces.VacancyRepository) controller.VacancyController {
	return &vacancyController{
		vacancyRepository: repo,
	}
}

func (c *vacancyController) FindAll(ctx *gin.Context) {
	companies, err := c.vacancyRepository.FindAll()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, companies)
}

func (c *vacancyController) FindByID(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return

	}
	var uid uint = uint(i)
	company, err := c.vacancyRepository.FindByID(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, company)

}

func (c *vacancyController) Create(ctx *gin.Context) {
	var vacancy schemas.VacancySchema
	err := ctx.BindJSON(&vacancy)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}

	vacancy, err = c.vacancyRepository.Create(vacancy)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, vacancy)
}

func (c *vacancyController) Delete(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	err = c.vacancyRepository.Delete(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}
