package Routers

import (
	"Finance/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	employee := v1.Group("/employee")
	{
		// employee.GET("/",)
		employee.GET("/get/:id", Controllers.GetOneEmployee)
		employee.POST("/create", Controllers.AddNewEmployee)
		employee.DELETE("/delete/:id", Controllers.DeleteEmployee)
	}

	return r
}
