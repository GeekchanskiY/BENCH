package routers

import (
	controllers "Finance/controllers"
	crtl_interfaces "Finance/controllers/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repo_interfaces "Finance/repositories/interfaces"

	repos "Finance/repositories"
)

func SetupSkill(base_group *gin.RouterGroup, db *gorm.DB) {
	var skillRepository repo_interfaces.SkillRepository = repos.NewSkillRepository(db)
	var SkillController crtl_interfaces.SkillController = controllers.NewSkillController(skillRepository)

	skill_group := base_group.Group("/skill")
	{
		skill_group.GET("/dependency", SkillController.FindAllDependency)
		skill_group.POST("/dependency", SkillController.CreateSkillDependency)
		skill_group.DELETE("/dependency", SkillController.DeleteSkillDependency)

		skill_group.GET("/conflict", SkillController.FindAllSkillConflicts)
		skill_group.POST("/conflict", SkillController.CreateSkillConflict)
		skill_group.DELETE("/conflict", SkillController.DeleteSkillConflict)

		skill_group.GET("/domain", SkillController.FindAllSkillDomains)
		skill_group.POST("/domain", SkillController.CreateSkillDomain)
		skill_group.DELETE("/domain", SkillController.DeleteSkillDomain)

		skill_group.GET("/", SkillController.FindAll)
		skill_group.GET("/:id", SkillController.FindByID)
		skill_group.GET("/:id/dependencies", SkillController.FindSkillDependencies)
		skill_group.GET("/:id/conflicts", SkillController.FindSkillConflicts)
		skill_group.GET("/:id/domains", SkillController.FindSkillDomains)
		skill_group.DELETE("/:id", SkillController.Delete)
		skill_group.POST("/", SkillController.Create)

	}
}
