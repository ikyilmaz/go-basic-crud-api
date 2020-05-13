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

	sync(db, SyncOptions{Force: false})

	return db
}

var (
	config           = mysql.Config{User: "root", ParseTime: true, DBName: "music"}
	connectionString = config.FormatDSN()
)

type SyncOptions struct {
	Force bool
}

func sync(db *gorm.DB, options SyncOptions) {
	var album Album
	var user User

	if options.Force {
		db.DropTableIfExists(&album, &user)
	}

	db.AutoMigrate(&user, &album)

	db.Model(&album).
		AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}

type BaseModel struct {
	ID        int        `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
}
