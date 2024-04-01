package routers

import (
	docs "Finance/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1"
	v1 := r.Group("/v1")
	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong1",
			})
		})
	}

	SetupEmployee(v1, db)
	SetupCompany(v1, db)
	SetupVacancy(v1, db)
	SetupDomain(v1, db)
	SetupSkill(v1, db)

	return r
}
