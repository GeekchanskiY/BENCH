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

	employee := base_group.Group("/skill")
	{
		employee.GET("/", SkillController.FindAll)
		employee.GET("/:id", SkillController.FindByID)
		employee.DELETE("/:id", SkillController.Delete)
		employee.POST("/", SkillController.Create)
	}
}
