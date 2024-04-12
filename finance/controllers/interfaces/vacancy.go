package interfaces

import (
	"github.com/gin-gonic/gin"
)

type VacancyController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)

	FindVacancyDomain(ctx *gin.Context)
	FindAllVacancyDomain(ctx *gin.Context)
	CreateVacancyDomain(ctx *gin.Context)
	DeleteVacancyDomain(ctx *gin.Context)

	FindVacancySkill(ctx *gin.Context)
	FindAllVacancySkill(ctx *gin.Context)
	CreateVacancySkill(ctx *gin.Context)
	DeleteVacancySkill(ctx *gin.Context)
}
