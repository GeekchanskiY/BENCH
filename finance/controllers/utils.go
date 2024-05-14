package controllers

import (
	controllers "Finance/controllers/interfaces"
	interfaces "Finance/repositories/interfaces"

	//"Finance/schemas"
	"github.com/gin-gonic/gin"
)

type utilsController struct {
	companyRepository        interfaces.CompanyRepository
	domainRepository         interfaces.DomainRepository
	employeeRepository       interfaces.EmployeeRepository
	cvRepository             interfaces.CVRepository
	responsibilityRepository interfaces.ResponsibilityRepository
	skillRepository          interfaces.SkillRepository
	vacancyRepository        interfaces.VacancyRepository
}

func NewUtilsController(
	companyRepo interfaces.CompanyRepository,
	domainRepo interfaces.DomainRepository,
	employeeRepo interfaces.EmployeeRepository,
	cvRepo interfaces.CVRepository,
	respRepo interfaces.ResponsibilityRepository,
	skillRepo interfaces.SkillRepository,
	vacancyRepo interfaces.VacancyRepository) controllers.UtilsController {
	return &utilsController{
		companyRepository:        companyRepo,
		domainRepository:         domainRepo,
		employeeRepository:       employeeRepo,
		cvRepository:             cvRepo,
		responsibilityRepository: respRepo,
		skillRepository:          skillRepo,
		vacancyRepository:        vacancyRepo,
	}
}

func (c *utilsController) Import(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (c *utilsController) Export(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
