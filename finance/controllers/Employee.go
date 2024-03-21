package Controllers

import (
	"Finance/ApiHelpers"
	"Finance/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetOneEmployee(c *gin.Context) {
	id := c.Params.ByName("id")
	var employee models.Employee
	err := models.GetOneEmployee(&employee, id)
	if err == nil {
		// log.Printf(employee.Name)
		ApiHelpers.RespondJSON(c, 200, employee)
		return
	}
	ApiHelpers.RespondJSON(c, 404, employee)

}

func DeleteEmployee(c *gin.Context) {
	id := c.Params.ByName("id")
	err := models.DeleteEmployee(id)
	if err == nil {
		c.JSON(200, gin.H{"deleted": "True"})
		return
	}
	c.JSON(400, gin.H{"error": "Cant delete :("})

}

func AddNewEmployee(c *gin.Context) {
	var employee models.Employee
	var err error
	c.BindJSON(&employee)
	err = employee.Validate()
	if err != nil {
		log.Printf("validation error")
		ApiHelpers.RespondJSON(c, 400, employee)
		return
	}
	err = models.AddNewEmployee(&employee)
	if err != nil {
		ApiHelpers.RespondJSON(c, 400, employee)
		return
	}

	ApiHelpers.RespondJSON(c, 200, employee)

}
