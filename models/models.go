package models

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-rest-api/lib"
	"time"
)

func NewDB() *gorm.DB {
	var db, err = gorm.Open("mysql", connectionString)

	lib.CheckErr(err)

	sync(db, SyncOptions{
		Force:   true,
		LogMode: true,
	})

	return db
}

var (
	config           = mysql.Config{User: "root", ParseTime: true, DBName: "music"}
	connectionString = config.FormatDSN()
)

type SyncOptions struct {
	Force   bool
	LogMode bool
}

func sync(db *gorm.DB, options SyncOptions) {
	var album Album
	var user User
	var userAlbum UserAlbum

	if options.Force {
		db.DropTableIfExists(&userAlbum, &album, &user)
	}

	db.AutoMigrate(&userAlbum, &user, &album)

	//db.Model(&album).
	//	AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Table("users_albums").
		AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT").
		AddForeignKey("album_id", "albums(id)", "CASCADE", "RESTRICT")

}

type BaseModel struct {
	ID        int        `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
}
