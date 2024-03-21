package routers

import (
	controllers "Finance/controllers"
	crtl_interfaces "Finance/controllers/interfaces"

	repos "Finance/repository"
	repo_interfaces "Finance/repository/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	var EmployeeRepository repo_interfaces.EmployeeRepository = repos.NewEmployeeRepository(db)
	var EmployeeController crtl_interfaces.EmployeeController = controllers.NewUserController(EmployeeRepository)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong1",
			})
		})
	}

	employee := v1.Group("/employee")
	{
		// employee.GET("/",)
		employee.GET("/get/:id", EmployeeController.FindByID)
		// employee.POST("/create", controllers.AddNewEmployee)
		// employee.DELETE("/delete/:id", controllers.DeleteEmployee)
	}

	return r
}
