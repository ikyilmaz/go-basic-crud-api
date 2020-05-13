package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-api/routes"
)

func InitMiddleWares(app *gin.Engine, db *gorm.DB) {
	app.Use(gin.Recovery())
	app.Use(gin.Logger())
	app.Use(setFields)

	routes.InitRoutes(app, db)
}
