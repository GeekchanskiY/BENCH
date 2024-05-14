package controllers

import (
	controller "Finance/controllers/interfaces"
	interfaces "Finance/repositories/interfaces"
	"Finance/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

type respController struct {
	respRepository interfaces.ResponsibilityRepository
}

func NewResponsibilityController(repo interfaces.ResponsibilityRepository) controller.ResponsibilityController {
	return &respController{
		respRepository: repo,
	}
}

func (c *respController) FindAll(ctx *gin.Context) {
	resps, err := c.respRepository.FindAll()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, resps)
}

func (c *respController) FindByID(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return

	}
	var uid uint = uint(i)
	resp, err := c.respRepository.FindByID(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, resp)

}

func (c *respController) Create(ctx *gin.Context) {
	var resp schemas.ResponsibilitySchema
	err := ctx.BindJSON(&resp)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	resp, err = c.respRepository.Create(resp)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, resp)
}

func (c *respController) Delete(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	err = c.respRepository.Delete(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}

//
// Responsibility Synonims
//

func (c *respController) FindAllSynonims(ctx *gin.Context) {
	skills, err := c.respRepository.FindAllResponsibilitySynonim()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *respController) FindSynonims(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	skills, err := c.respRepository.FindResponsibilitySynonim(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *respController) DeleteSynonim(ctx *gin.Context) {
	var synonim schemas.ResponsibilitySynonimSchema
	err := ctx.BindJSON(&synonim)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	err = c.respRepository.DeleteResponsibilitySynonim(synonim)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}

func (c *respController) CreateSynonim(ctx *gin.Context) {
	var synonim schemas.ResponsibilitySynonimSchema

	err := ctx.BindJSON(&synonim)
	if err != nil {

		ctx.JSON(400, ctx.Error(err))
		return
	}

	synonim, err = c.respRepository.CreateResponsibilitySynonim(synonim)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, synonim)
}

//
// Responsibility Conflict
//

func (c *respController) CreateRespConflict(ctx *gin.Context) {
	var conflict schemas.ResponsibilityConflictSchema

	err := ctx.BindJSON(&conflict)
	if err != nil {

		ctx.JSON(400, ctx.Error(err))
		return
	}

	conflict, err = c.respRepository.CreateResponsibilityConflict(conflict)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, conflict)
}

func (c *respController) FindConflicts(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	conflicts, err := c.respRepository.FindResponsibilityConflict(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, conflicts)
}

func (c *respController) FindAllRespConflicts(ctx *gin.Context) {
	conflicts, err := c.respRepository.FindAllResponsibilityConflicts()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, conflicts)
}

func (c *respController) DeleteRespConflict(ctx *gin.Context) {
	var conflict schemas.ResponsibilityConflictSchema
	err := ctx.BindJSON(&conflict)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	err = c.respRepository.DeleteResponsibilityConflict(conflict)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}
