package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest-api/lib"
	"go-rest-api/models"
	"strconv"
)

func getUserDefaults(db *gorm.DB) *gorm.DB {
	return db.Select("id, first_name, last_name, username")
}

type UserService struct {
	*gorm.DB
}

func (u UserService) GetUser(ctx *gin.Context, user *models.User) error {
	user.ID = ctx.GetInt("id")
	return u.Scopes(getUserDefaults).Find(user).Error
}

func (u UserService) GetUserWith(ctx *gin.Context, user *models.User, preload string) error {
	user.ID = ctx.GetInt("id")
	return u.Preload(preload).Scopes(getUserDefaults).Find(user).Error
}

func (u UserService) GetManyUser(ctx *gin.Context, users *[]models.User) error {
	return u.Scopes(getUserDefaults, lib.Paginator(ctx, 20, 25)).Find(users).Error
}

func (u UserService) CreateUser(user *models.User) error {
	return u.Create(user).Error
}

func (u UserService) UpdateUser(ctx *gin.Context, user *models.User) error {
	user.ID = ctx.GetInt("id")
	return u.Table("users").Update(user).Error
}

func (u UserService) DeleteUser(ctx *gin.Context, user *models.User) error {
	user.ID = ctx.GetInt("id")
	return u.Delete(user).Error
}

func (u UserService) ShouldReturn(ctx *gin.Context, user *models.User) {
	if returning, err := strconv.ParseBool(ctx.DefaultQuery("returning", "false")); returning && err != nil {
		_ = u.GetUser(ctx, user)
	}
}
