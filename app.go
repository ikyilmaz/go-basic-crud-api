package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-rest-api/middlewares"
	"go-rest-api/models"
)

func main() {
	db := models.NewDB()
	app := gin.New()

	middlewares.InitMiddleWares(app, db)

	_ = app.Run(":8080")

}
