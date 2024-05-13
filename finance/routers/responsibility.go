package routers

import (
	controllers "Finance/controllers"
	crtl_interfaces "Finance/controllers/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repo_interfaces "Finance/repositories/interfaces"

	repos "Finance/repositories"
)

func SetupResponsibility(base_group *gin.RouterGroup, db *gorm.DB) {
	var responsibilityRepository repo_interfaces.ResponsibilityRepository = repos.NewResponsibilityRepository(db)
	var ResponsibilityController crtl_interfaces.ResponsibilityController = controllers.NewResponsibilityController(responsibilityRepository)

	responsibility_group := base_group.Group("/responsibility")
	{
		responsibility_group.GET("/synonim", ResponsibilityController.FindAllSynonims)
		responsibility_group.POST("/synonim", ResponsibilityController.CreateSynonim)
		responsibility_group.DELETE("/synonim", ResponsibilityController.DeleteSynonim)

		responsibility_group.GET("/conflict", ResponsibilityController.FindAllRespConflicts)
		responsibility_group.POST("/conflict", ResponsibilityController.CreateRespConflict)
		responsibility_group.DELETE("/conflict", ResponsibilityController.DeleteRespConflict)

		responsibility_group.GET("/", ResponsibilityController.FindAll)
		responsibility_group.GET("/:id", ResponsibilityController.FindByID)
		responsibility_group.DELETE("/:id", ResponsibilityController.Delete)
		responsibility_group.GET("/:id/synonims", ResponsibilityController.FindSynonims)
		responsibility_group.GET("/:id/conflicts", ResponsibilityController.FindConflicts)
		responsibility_group.POST("/", ResponsibilityController.Create)

	}
}
