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
	if err != nil {
		ctx.AbortWithStatusJSON(lib.NotFound())
		return
	}
	lib.SendResponse(lib.SendResponseOptions{Context: ctx, Data: user})
}

func (u UserController) GetMany(ctx *gin.Context) {
	var users []models.User
	u.Service.GetManyUser(ctx, &users)
	lib.SendResponse(lib.SendResponseOptions{Context: ctx, Data: users})
}

func (u UserController) Create(ctx *gin.Context) {
	var user = ctx.MustGet("create-user").(models.User)
	_ = u.Service.CreateUser(ctx, &user)
	u.Service.ShouldReturn(ctx, &user)
	lib.SendResponse(lib.SendResponseOptions{Context: ctx, Data: user, StatusCode: 201})
}

func (u UserController) Update(ctx *gin.Context) {
	var user = ctx.MustGet("update-user").(models.User)
	user.ID = ctx.GetInt("id")
	_ = u.Service.UpdateUser(ctx, &user)
	u.Service.ShouldReturn(ctx, &user)
	lib.SendResponse(lib.SendResponseOptions{Context: ctx})
}

func (u UserController) Delete(ctx *gin.Context) {
	var user models.User
	user.ID = ctx.GetInt("id")
	_ = u.Service.DeleteUser(ctx, &user)
	lib.SendResponse(lib.SendResponseOptions{Context: ctx, StatusCode: 204})
}
