package ApiHelpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Data   interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {

	var res ResponseData

	res.Status = status

	res.Data = payload

	w.JSON(status, res)
}
