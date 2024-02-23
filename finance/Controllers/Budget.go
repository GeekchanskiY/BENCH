package Controllers

import (
	"Finance/ApiHelpers"
	"Finance/Models"

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
