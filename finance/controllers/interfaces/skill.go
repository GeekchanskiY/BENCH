package interfaces

import (
	"github.com/gin-gonic/gin"
)

type SkillController interface {
	// Skill
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)

	// Skill Dependency
	FindSkillDependencies(ctx *gin.Context)
	CreateSkillDependency(ctx *gin.Context)
	DeleteSkillDependency(ctx *gin.Context)
	FindAllDependency(ctx *gin.Context)

	// Skill Conflict
	CreateSkillConflict(ctx *gin.Context)
	DeleteSkillConflict(ctx *gin.Context)
	FindAllSkillConflicts(ctx *gin.Context)

	// Skill Domain
	CreateSkillDomain(ctx *gin.Context)
	DeleteSkillDomain(ctx *gin.Context)
	FindAllSkillDomains(ctx *gin.Context)
}
