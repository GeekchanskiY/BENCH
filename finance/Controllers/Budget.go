package Controllers

import (
	"Finance/ApiHelpers"
	"Finance/Models"
	"log"

	"fmt"

	"github.com/gin-gonic/gin"
)

func GetOneBudget(c *gin.Context) {
	id := c.Params.ByName("id")
	var budget Models.Budget
	err := Models.GetOneBudget(&budget, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, budget)
	} else {
		ApiHelpers.RespondJSON(c, 200, budget)
	}
}

func ListBudget(c *gin.Context) {
	var budget []Models.Budget
	err := Models.GetAllBudgets(&budget)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, budget)
	} else {
		ApiHelpers.RespondJSON(c, 200, budget)
	}
}

func AddNewBudget(c *gin.Context) {
	var budget Models.Budget
	var err error
	c.BindJSON(&budget)
	fmt.Println(budget)

	err = budget.Validate()
	if err != nil {
		log.Printf("validation error")
		ApiHelpers.RespondJSON(c, 201, budget)
		return
	}

	err = Models.AddNewBudget(&budget)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, budget)
		return
	} else {
		ApiHelpers.RespondJSON(c, 200, budget)
		return
	}
}
