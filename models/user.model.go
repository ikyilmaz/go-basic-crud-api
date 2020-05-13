package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	FirstName string `json:"firstName,omitempty" gorm:"type:varchar(64);"`
	LastName  string `json:"lastName,omitempty"  gorm:"type:varchar(64);"`
	Username  string `json:"username,omitempty"  gorm:"type:varchar(32);unique;not null;"`
	Email     string `json:"email,omitempty"     gorm:"type:varchar(64);unique;not null;"`
	Password  string `json:"password,omitempty"  gorm:"not null;"`
	Role      string `json:"role,omitempty"      gorm:"type:varchar(64);default:'user';type:enum('user','artist','admin')"`

	Albums []*Album `json:"albums,omitempty"     gorm:"many2many:users_albums;"`
}

func (u *User) BeforeSave() {

	if &u.Password != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
		u.Password = string(hashedPassword)
	}

}
