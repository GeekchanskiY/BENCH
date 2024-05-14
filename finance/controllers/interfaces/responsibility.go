package interfaces

import (
	"github.com/gin-gonic/gin"
)

type ResponsibilityController interface {
	// Responsibility
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)

	// Responsibility Synonims
	CreateSynonim(ctx *gin.Context)
	DeleteSynonim(ctx *gin.Context)
	FindSynonims(ctx *gin.Context)
	FindAllSynonims(ctx *gin.Context)

	// Responsibility Conflict
	CreateRespConflict(ctx *gin.Context)
	DeleteRespConflict(ctx *gin.Context)
	FindConflicts(ctx *gin.Context)
	FindAllRespConflicts(ctx *gin.Context)
}
