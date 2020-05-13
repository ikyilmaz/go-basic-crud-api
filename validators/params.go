package validators

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckIDParam(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	ctx.Set("id", id)

	if err != nil {
		err = errors.New("param 'id' must be int")
	}

	setValidationErrors(ctx, err)

	ctx.Next()
}
