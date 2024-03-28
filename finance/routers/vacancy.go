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
	var VacancyController crtl_interfaces.EmployeeController = controllers.NewVacancyController(VacancyRepository)

	employee := base_group.Group("/vacancy")
	{
		employee.GET("/", VacancyController.FindAll)
		employee.GET("/:id", VacancyController.FindByID)
		employee.DELETE("/:id", VacancyController.Delete)
		employee.POST("/", VacancyController.Create)
	}
}
