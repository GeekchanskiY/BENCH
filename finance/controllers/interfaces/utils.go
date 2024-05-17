package interfaces

import "github.com/gin-gonic/gin"

type UtilsController interface {
	Import(c *gin.Context)
	Export(c *gin.Context)
}
