package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/lib"
	"go-rest-api/models"
	"go-rest-api/services"
)

type UserController struct {
	Service *services.UserService
}

func (u UserController) Get(ctx *gin.Context) {
	var user models.User
	err := u.Service.GetUser(ctx, &user)
	lib.CheckErr(err)
	if err != nil {
		ctx.AbortWithStatusJSON(lib.NotFound())
		return
	}
	lib.SendResponse(lib.SendResponseOptions{Context: ctx, Data: user})
}

func (u UserController) GetMany(ctx *gin.Context) {
	var users []models.User
	err := u.Service.GetManyUser(ctx, &users)
	lib.CheckErr(err)
	lib.SendResponse(lib.SendResponseOptions{Context: ctx, Data: users})
}

func (u UserController) Create(ctx *gin.Context) {
	var user = ctx.MustGet("create-user").(models.User)
	err := u.Service.CreateUser(&user)
	lib.CheckErr(err)
	u.Service.ShouldReturn(ctx, &user)
	lib.SendResponse(lib.SendResponseOptions{Context: ctx, Data: user, StatusCode: 201})
}

func (u UserController) Update(ctx *gin.Context) {
	var user = ctx.MustGet("update-user").(models.User)
	err := u.Service.UpdateUser(ctx, &user)
	lib.CheckErr(err)
	u.Service.ShouldReturn(ctx, &user)
	lib.SendResponse(lib.SendResponseOptions{Context: ctx})
}

func (u UserController) Delete(ctx *gin.Context) {
	var user models.User
	err := u.Service.DeleteUser(ctx, &user)
	lib.CheckErr(err)
	lib.SendResponse(lib.SendResponseOptions{Context: ctx, StatusCode: 204})
}

func (u UserController) GetUserWith(preload string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var user models.User
		err := u.Service.GetUserWith(ctx, &user, preload)
		lib.CheckErr(err)
		lib.SendResponse(lib.SendResponseOptions{Context: ctx, Data: user, StatusCode: 200})
	}
}
