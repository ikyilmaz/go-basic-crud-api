package middlewares

import "github.com/gin-gonic/gin"

func setFields(ctx *gin.Context) {
	var validationErrors []error
	ctx.Set("validation-errors", validationErrors)
}
