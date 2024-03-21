package controllers

import (
	controller "Finance/controllers/interfaces"
	"Finance/models"
	interfaces "Finance/repository/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

type employeeController struct {
	employeeRepository interfaces.EmployeeRepository
}

func NewUserController(repo interfaces.EmployeeRepository) controller.EmployeeController {
	return &employeeController{
		employeeRepository: repo,
	}
}

func (c *employeeController) FindAll(ctx *gin.Context) {
	c.employeeRepository.FindAll(ctx)

}

func (c *employeeController) FindByID(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		panic(err)
	}
	var ui uint = uint(i)
	employee, err := c.employeeRepository.FindByID(ctx, ui)
	ctx.JSON(200, gin.H{"nice": employee.Name})

}

func (c *employeeController) Create(ctx *gin.Context) {
	var employee models.Employee
	employee, err := c.employeeRepository.Create(ctx, employee)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "asd"})
	}
	ctx.JSON(200, gin.H{"name": employee.Name})

}

func (c *employeeController) Delete(ctx *gin.Context) {
	var employee models.Employee
	err := c.employeeRepository.Delete(ctx, employee)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "error"})
	}
}
