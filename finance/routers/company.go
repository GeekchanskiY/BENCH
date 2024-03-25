package routers

import (
	controllers "Finance/controllers"
	crtl_interfaces "Finance/controllers/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repo_interfaces "Finance/repositories/interfaces"

	repos "Finance/repositories"
)

func SetupCompany(base_group *gin.RouterGroup, db *gorm.DB) {
	var companyRepository repo_interfaces.CompanyRepository = repos.NewCompanyRepository(db)
	var CompanyController crtl_interfaces.CompanyController = controllers.NewCompanyController(companyRepository)

	employee := base_group.Group("/company")
	{
		employee.GET("/", CompanyController.FindAll)
		employee.GET("/:id", CompanyController.FindByID)
		employee.DELETE("/:id", CompanyController.Delete)
		employee.POST("/", CompanyController.Create)
	}
}
