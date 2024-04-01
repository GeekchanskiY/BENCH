package routers

import (
	controllers "Finance/controllers"
	crtl_interfaces "Finance/controllers/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repo_interfaces "Finance/repositories/interfaces"

	repos "Finance/repositories"
)

func SetupDomain(base_group *gin.RouterGroup, db *gorm.DB) {
	var domainRepository repo_interfaces.DomainRepository = repos.NewDomainRepository(db)
	var DomainController crtl_interfaces.DomainController = controllers.NewDomainController(domainRepository)

	employee := base_group.Group("/domain")
	{
		employee.GET("/", DomainController.FindAll)
		employee.GET("/:id", DomainController.FindByID)
		employee.DELETE("/:id", DomainController.Delete)
		employee.POST("/", DomainController.Create)
	}
}
