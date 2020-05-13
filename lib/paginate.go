package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func Paginator(ctx *gin.Context, args ...int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, err := strconv.Atoi(ctx.Query("page"))

		if err != nil {
			page = 1
		}

		limit, err := strconv.Atoi(ctx.Query("limit"))

		if err != nil {
			limit = args[0]
		}

		var maxLimit = 50

		if len(args) == 2 {
			maxLimit = args[1]
		}

		if limit > maxLimit {
			limit = maxLimit
		}

		offset := (page * limit) - limit

		return db.Offset(offset).Limit(limit)
	}
}
