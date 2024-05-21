package controllers

import (
	controllers "Finance/controllers/interfaces"
	interfaces "Finance/repositories/interfaces"
	"Finance/schemas"
	"fmt"
	"io"

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

type Backup struct {
	Companies []schemas.CompanySchema  `json:"companies"`
	Domains   []schemas.DomainSchema   `json:"domains"`
	Employees []schemas.EmployeeSchema `json:"employees"`
	Cvs       []schemas.CVSchema       `json:"cvs"`
	Skills    []schemas.SkillSchema    `json:"skills"`
	Vacancies []schemas.VacancySchema  `json:"vacancies"`

	SkillDependencies []schemas.SkillDependencySchema `json:"skillDependencies"`
	SkillConflicts    []schemas.SkillConflictSchema   `json:"skillConflicts"`
	SkillDomains      []schemas.SkillDomainSchema     `json:"skillDomains"`

	ResponsibilitySynonims  []schemas.ResponsibilitySynonimSchema  `json:"responsibilitySynonims"`
	ResponsibilityConflicts []schemas.ResponsibilityConflictSchema `json:"responsibilityConflicts"`

	VacancySkills  []schemas.VacancySkillSchema  `json:"vacancySkills"`
	VacancyDomains []schemas.VacancyDomainSchema `json:"vacancyDomains"`
}

func (c *utilsController) Import(ctx *gin.Context) {
	var backup Backup = Backup{}
	var errors []string = []string{}
	body, _ := io.ReadAll(ctx.Request.Body)
	println(string(body))
	err := ctx.BindJSON(&backup)
	if err != nil {
		fmt.Println(err)
		errors = append(errors, err.Error())
	}

	for _, company := range backup.Companies {
		_, err = c.companyRepository.FindByName(company.Name)
		if err == nil {
			continue
		} else {
			fmt.Println(err)
		}
		_, err := c.companyRepository.Create(company)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		ctx.JSON(400, gin.H{
			"message": errors,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Imported",
	})
}

func (c *utilsController) Export(ctx *gin.Context) {

	companies, err := c.companyRepository.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	domains, err := c.domainRepository.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	employees, err := c.employeeRepository.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	cvs, err := c.cvRepository.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	skills, err := c.skillRepository.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	vacancies, err := c.vacancyRepository.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	skillDependencies, err := c.skillRepository.FindAllDependency()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	skillConflicts, err := c.skillRepository.FindAllSkillConflicts()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	skillDomains, err := c.skillRepository.FindAllSkillDomains()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	responsibilitySynonims, err := c.responsibilityRepository.FindAllResponsibilitySynonim()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	responsibilityConflicts, err := c.responsibilityRepository.FindAllResponsibilityConflicts()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	vacancySkills, err := c.vacancyRepository.FindAllVacancySkill()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	vacancyDomains, err := c.vacancyRepository.FindAllVacancyDomain()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	backup := Backup{
		Companies: companies,
		Domains:   domains,
		Employees: employees,
		Cvs:       cvs,
		Skills:    skills,
		Vacancies: vacancies,

		SkillDependencies: skillDependencies,
		SkillConflicts:    skillConflicts,
		SkillDomains:      skillDomains,

		ResponsibilitySynonims:  responsibilitySynonims,
		ResponsibilityConflicts: responsibilityConflicts,

		VacancySkills:  vacancySkills,
		VacancyDomains: vacancyDomains,
	}
	// fmt.Println(backup)

	ctx.JSON(200, gin.H{
		"content": backup,
	})
}
