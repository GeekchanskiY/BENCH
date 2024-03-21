package main

import (
	"log"

	"Finance/Routers"
	"Finance/config"
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
	config.SetupDBConnection(db)

	config.Setup()

	r := Routers.SetupRouter()

	r.Use(LoggerMiddleware())

	r.Run(":3001")
}
