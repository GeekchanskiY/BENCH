package routers

import (
	controllers "Finance/controllers"
	crtl_interfaces "Finance/controllers/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repo_interfaces "Finance/repositories/interfaces"

	repos "Finance/repositories"
)

func SetupVacancy(base_group *gin.RouterGroup, db *gorm.DB) {
	var VacancyRepository repo_interfaces.VacancyRepository = repos.NewVacancyRepository(db)
	var VacancyController crtl_interfaces.VacancyController = controllers.NewVacancyController(VacancyRepository)

	vacancy := base_group.Group("/vacancy")
	{
		vacancy.GET("/skill", VacancyController.FindAllVacancySkill)
		vacancy.POST("/skill", VacancyController.CreateVacancySkill)
		vacancy.DELETE("/skill", VacancyController.DeleteVacancySkill)

		vacancy.GET("/domain", VacancyController.FindAllVacancyDomain)
		vacancy.POST("/domain", VacancyController.CreateVacancyDomain)
		vacancy.DELETE("/domain", VacancyController.DeleteVacancyDomain)

		vacancy.GET("/", VacancyController.FindAll)
		vacancy.GET("/:id", VacancyController.FindByID)
		vacancy.GET("/:id/domain", VacancyController.FindVacancyDomain)
		vacancy.GET("/:id/skill", VacancyController.FindVacancySkill)
		vacancy.DELETE("/:id", VacancyController.Delete)
		vacancy.POST("/", VacancyController.Create)
	}
}
