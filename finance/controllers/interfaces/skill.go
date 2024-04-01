package interfaces

import (
	"github.com/gin-gonic/gin"
)

type SkillController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	CreateSkillDependency(ctx *gin.Context)
	DeleteSkillDependency(ctx *gin.Context)
	FindAllDependency(ctx *gin.Context)
}
