package lib

import (
	"github.com/gin-gonic/gin"
)

type SendResponseOptions struct {
	Context    *gin.Context
	Data       interface{}
	StatusCode int
}

func SendResponse(options SendResponseOptions) {
	var ctx = options.Context

	var statusCode int

	if options.StatusCode != 0 {
		statusCode = options.StatusCode
	} else {
		statusCode = 200
	}

	var response = map[string]interface{}{
		"status": "success",
	}

	if &options.Data != nil {
		response["data"] = options.Data
	}

	ctx.JSON(statusCode, response)

}
