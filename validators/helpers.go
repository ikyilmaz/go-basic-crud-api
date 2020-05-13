package validators

import (
	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
)

func setValidationErrors(ctx *gin.Context, err ...error) {
	ctx.Set("validation-errors", append(ctx.MustGet("validation-errors").([]error), err...))
}

func isRequired(schema *jio.StringSchema, required bool) {
	if required {
		schema = schema.Required()
	} else {
		schema = schema.Optional()
	}
}
