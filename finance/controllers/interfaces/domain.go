package interfaces

import (
	"github.com/gin-gonic/gin"
)

type DomainController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
