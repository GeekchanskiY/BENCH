package interfaces

import (
	"github.com/gin-gonic/gin"
)

type CVController interface {
	// CV
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)

	// CV Responsibility
	CreateCVResponsibility(ctx *gin.Context)
	DeleteCVResponsibility(ctx *gin.Context)
	FindAllCVResponsibility(ctx *gin.Context)

	// CV Project
	CreateCVProject(ctx *gin.Context)
	DeleteCVProject(ctx *gin.Context)
	FindAllCVProject(ctx *gin.Context)
}
