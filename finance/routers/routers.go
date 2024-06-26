package routers

import (
	docs "Finance/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
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
	SetupResponsibility(v1, db)
	SetupCV(v1, db)
	SetupUtils(v1, db)

	return r
}
