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
				"message": "pong2",
			})
		})
		v1.GET("/", Controllers.ListBudget)
		v1.POST("/add", Controllers.AddNewBudget)
		v1.GET("/get/:id", Controllers.GetOneBudget)
	}

	return r
}
