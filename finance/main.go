package main

import (
	"log"

	"Finance/config"
	"Finance/routers"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

func main() {
	var db *gorm.DB

	// config.SetupDBConnection(db)

	db = config.Setup(db)

	r := routers.SetupRouter(db)

	r.Use(gin.ErrorLogger())

	r.Use(LoggerMiddleware())

	r.Run(":3001")
}
