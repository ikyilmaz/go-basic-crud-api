package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func InitRoutes(app *gin.Engine, db2 *gorm.DB) {
	var router = app.Group("api/v1")
	db = db2
	userRoutes(router.Group("/users"))
}
