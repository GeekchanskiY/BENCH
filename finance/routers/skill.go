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
		skill_group.GET("/", SkillController.FindAll)
		skill_group.GET("/:id", SkillController.FindByID)
		skill_group.DELETE("/:id", SkillController.Delete)
		skill_group.POST("/", SkillController.Create)

	}
}
