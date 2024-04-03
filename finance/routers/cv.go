package routers

import (
	controllers "Finance/controllers"
	crtl_interfaces "Finance/controllers/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repo_interfaces "Finance/repositories/interfaces"

	repos "Finance/repositories"
)

func SetupCV(base_group *gin.RouterGroup, db *gorm.DB) {
	var cvRepository repo_interfaces.CVRepository = repos.NewCVRepository(db)
	var CVController crtl_interfaces.CVController = controllers.NewCVController(cvRepository)

	cv_group := base_group.Group("/cv")
	{
		cv_group.GET("/responsibility", CVController.FindAllCVResponsibility)
		cv_group.POST("/responsibility", CVController.CreateCVResponsibility)
		cv_group.DELETE("/responsibility", CVController.DeleteCVResponsibility)

		cv_group.GET("/project", CVController.FindAllCVProject)
		cv_group.POST("/project", CVController.CreateCVProject)
		cv_group.DELETE("/project", CVController.DeleteCVProject)

		cv_group.GET("/", CVController.FindAll)
		cv_group.GET("/:id", CVController.FindByID)
		cv_group.DELETE("/:id", CVController.Delete)
		cv_group.POST("/", CVController.Create)

	}
}
