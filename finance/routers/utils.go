package routers

import (
	controllers "Finance/controllers"
	ctrl_interfaces "Finance/controllers/interfaces"

	repos "Finance/repositories"
	interfaces "Finance/repositories/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUtils(base_group *gin.RouterGroup, db *gorm.DB) {
	var skillRepo interfaces.SkillRepository = repos.NewSkillRepository(db)
	var cvRepo interfaces.CVRepository = repos.NewCVRepository(db)
	var domainRepo interfaces.DomainRepository = repos.NewDomainRepository(db)
	var companyRepo interfaces.CompanyRepository = repos.NewCompanyRepository(db)
	var employeeRepo interfaces.EmployeeRepository = repos.NewEmployeeRepository(db)
	var responsibilityRepo interfaces.ResponsibilityRepository = repos.NewResponsibilityRepository(db)
	var vacancyRepo interfaces.VacancyRepository = repos.NewVacancyRepository(db)

	var utilsController ctrl_interfaces.UtilsController = controllers.NewUtilsController(
		companyRepo,
		domainRepo,
		employeeRepo,
		cvRepo,
		responsibilityRepo,
		skillRepo,
		vacancyRepo,
	)
	utils := base_group.Group("/utils")
	{
		utils.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
		utils.POST("/export", utilsController.Export)
		utils.GET("/import", utilsController.Import)
	}
}
