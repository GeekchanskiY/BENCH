package routers

import (
	controllers "Finance/controllers"
	crtl_interfaces "Finance/controllers/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repo_interfaces "Finance/repository/interfaces"

	repos "Finance/repository"
)

func SetupEmployee(base_group *gin.RouterGroup, db *gorm.DB) {
	var EmployeeRepository repo_interfaces.EmployeeRepository = repos.NewEmployeeRepository(db)
	var EmployeeController crtl_interfaces.EmployeeController = controllers.NewUserController(EmployeeRepository)

	employee := base_group.Group("/employee")
	{
		employee.GET("/", EmployeeController.FindAll)
		employee.GET("/:id", EmployeeController.FindByID)
		employee.DELETE("/:id", EmployeeController.Delete)
		employee.POST("/", EmployeeController.Create)
	}
}
