package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/controllers"
	"go-rest-api/filters"
	"go-rest-api/services"
	"go-rest-api/validators"
)

func userRoutes(router *gin.RouterGroup) {
	var userController = controllers.UserController{
		Service: &services.UserService{DB: db},
	}

	router.
		GET("/", userController.GetMany).
		POST(
			"/",
			validators.CreateOrUpdateUser(true),
			filters.CheckValidationResult,
			userController.Create,
		).
		GET(
			"/:id",
			validators.CheckIDParam,
			filters.CheckValidationResult,
			userController.Get,
		).
		PATCH(
			"/:id",
			validators.CreateOrUpdateUser(false),
			filters.CheckValidationResult,
			userController.Update,
		).
		DELETE("/:id", userController.Get)
}
