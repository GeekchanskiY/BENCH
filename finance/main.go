package main

import (
	"bench/finance/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func getPersons(c *gin.Context) {
	persons, err := models.GetPersons(10)
	checkErr(err)

	if persons == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": persons})
	}
}
func main() {
	var r *gin.Engine = gin.Default()
	r.Use(LoggerMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", getPersons)
	v1 := r.Group("/v1")
	v2 := v1.Group("/v12")

	{
		v2.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	r.Run(":8085")
}
