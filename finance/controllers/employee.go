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
	employees, err := c.employeeRepository.FindAll()
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, employees)
}

func (c *employeeController) FindByID(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return

	}
	var uid uint = uint(i)
	employee, err := c.employeeRepository.FindByID(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, employee)

}

func (c *employeeController) Create(ctx *gin.Context) {
	var employee models.Employee
	ctx.BindJSON(&employee)
	employee, err := c.employeeRepository.Create(employee)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, employee)
}

func (c *employeeController) Delete(ctx *gin.Context) {
	var params_id string = ctx.Params.ByName("id")
	i, err := strconv.ParseUint(params_id, 10, 32)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	var uid uint = uint(i)
	err = c.employeeRepository.Delete(uid)
	if err != nil {
		ctx.JSON(400, ctx.Error(err))
		return
	}
	ctx.JSON(200, gin.H{"deleted": true})
}
