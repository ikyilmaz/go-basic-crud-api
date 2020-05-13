package filters

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/lib"
)

func CheckValidationResult(ctx *gin.Context) {
	value := ctx.MustGet("validation-errors").([]error)

	if value != nil && len(value) > 0 && value[0] != nil {
		ctx.AbortWithStatusJSON(lib.AppError(value[0].Error(), 400))
	}

	ctx.Next()
}
