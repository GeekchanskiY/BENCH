package controllers

import (
	controller "Finance/controllers/interfaces"
	interfaces "Finance/repositories/interfaces"
	"Finance/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

type skillController struct {
	skillRepository interfaces.SkillRepository
}

func NewSkillController(repo interfaces.SkillRepository) controller.SkillController {
	return &skillController{
		skillRepository: repo,
	}
}

func (c *skillController) FindAll(ctx *gin.Context) {
	skills, err := c.skillRepository.FindAll()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *skillController) FindByID(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return

	}
	var uid uint = uint(i)
	skill, err := c.skillRepository.FindByID(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skill)

}

func (c *skillController) Create(ctx *gin.Context) {
	var skill schemas.SkillSchema
	err := ctx.BindJSON(&skill)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	skill, err = c.skillRepository.Create(skill)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skill)
}

func (c *skillController) Delete(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	err = c.skillRepository.Delete(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}

//
// Skill Dependency
//

func (c *skillController) FindSkillDependencies(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	skilldeps, err := c.skillRepository.FindSkillDependencies(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skilldeps)
}

func (c *skillController) FindAllDependency(ctx *gin.Context) {
	skills, err := c.skillRepository.FindAllDependency()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *skillController) DeleteSkillDependency(ctx *gin.Context) {
	var skilldep schemas.SkillDependencySchema
	err := ctx.BindJSON(&skilldep)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	err = c.skillRepository.DeleteSkillDependency(skilldep)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}

func (c *skillController) CreateSkillDependency(ctx *gin.Context) {
	var skilldep schemas.SkillDependencySchema

	err := ctx.BindJSON(&skilldep)
	if err != nil {

		ctx.JSON(400, ctx.Error(err))
		return
	}

	skilldep, err = c.skillRepository.CreateSkillDependency(skilldep)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skilldep)
}

//
// Skill Conflict
//

func (c *skillController) FindSkillConflicts(ctx *gin.Context) {
	var skillconflicts []schemas.SkillConflictSchema
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	skillconflicts, err = c.skillRepository.FindSkillConflicts(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skillconflicts)
}

func (c *skillController) CreateSkillConflict(ctx *gin.Context) {
	var skilldep schemas.SkillConflictSchema

	err := ctx.BindJSON(&skilldep)
	if err != nil {

		ctx.JSON(400, ctx.Error(err))
		return
	}

	skilldep, err = c.skillRepository.CreateSkillConflict(skilldep)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skilldep)
}

func (c *skillController) FindAllSkillConflicts(ctx *gin.Context) {
	skills, err := c.skillRepository.FindAllSkillConflicts()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *skillController) DeleteSkillConflict(ctx *gin.Context) {
	var skilldep schemas.SkillConflictSchema
	err := ctx.BindJSON(&skilldep)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	err = c.skillRepository.DeleteSkillConflict(skilldep)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}

//
// Skill Domain
//

func (c *skillController) FindSkillDomains(ctx *gin.Context) {
	var skilldomains []schemas.SkillDomainSchema
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	skilldomains, err = c.skillRepository.FindSkillDomains(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skilldomains)
}

func (c *skillController) CreateSkillDomain(ctx *gin.Context) {
	var skilldom schemas.SkillDomainSchema

	err := ctx.BindJSON(&skilldom)
	if err != nil {

		ctx.JSON(400, ctx.Error(err))
		return
	}

	skilldom, err = c.skillRepository.CreateSkillDomain(skilldom)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skilldom)
}

func (c *skillController) FindAllSkillDomains(ctx *gin.Context) {
	skills, err := c.skillRepository.FindAllSkillDomains()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, skills)
}

func (c *skillController) DeleteSkillDomain(ctx *gin.Context) {
	var skilldom schemas.SkillDomainSchema
	err := ctx.BindJSON(&skilldom)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	err = c.skillRepository.DeleteSkillDomain(skilldom)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}
