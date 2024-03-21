package main

import (
	"fmt"
	"log"

	"Finance/Models"
	"Finance/Routers"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
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
	var err error
	Models.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("status: ", err)
	}
	Models.Migrate()
	r := Routers.SetupRouter()

	r.Use(LoggerMiddleware())

	r.Run(":3001")
}
