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
		vacancy.GET("/skills", VacancyController.FindAllVacancySkill)
		vacancy.POST("/skills", VacancyController.CreateVacancySkill)
		vacancy.DELETE("/skills", VacancyController.DeleteVacancySkill)

		vacancy.GET("/domains", VacancyController.FindAllVacancyDomain)
		vacancy.POST("/domains", VacancyController.CreateVacancyDomain)
		vacancy.DELETE("/domains", VacancyController.DeleteVacancyDomain)

		vacancy.GET("/", VacancyController.FindAll)
		vacancy.GET("/:id", VacancyController.FindByID)
		vacancy.GET("/:id/domains", VacancyController.FindVacancyDomain)
		vacancy.GET("/:id/skills", VacancyController.FindVacancySkill)
		vacancy.DELETE("/:id", VacancyController.Delete)
		vacancy.POST("/", VacancyController.Create)
	}
}
