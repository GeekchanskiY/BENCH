package Controllers

import (
	"Finance/ApiHelpers"
	"Finance/Models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetOneEmployee(c *gin.Context) {
	id := c.Params.ByName("id")
	var employee Models.Employee
	err := Models.GetOneEmployee(&employee, id)
	if err == nil {
		// log.Printf(employee.Name)
		ApiHelpers.RespondJSON(c, 200, employee)
		return
	}
	ApiHelpers.RespondJSON(c, 404, employee)

}

func AddNewEmployee(c *gin.Context) {
	var employee Models.Employee
	var err error
	c.BindJSON(&employee)
	err = employee.Validate()
	if err != nil {
		log.Printf("validation error")
		ApiHelpers.RespondJSON(c, 400, employee)
		return
	}
	err = Models.AddNewEmployee(&employee)
	if err != nil {
		ApiHelpers.RespondJSON(c, 400, employee)
		return
	}

	ApiHelpers.RespondJSON(c, 200, employee)

}
