package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	FirstName string `json:"firstName,omitempty" gorm:"type:varchar(64);"`
	LastName  string `json:"lastName,omitempty" gorm:"type:varchar(64);"`
	Username  string `json:"username,omitempty" gorm:"type:varchar(32);unique;not null;"`
	Email     string `json:"email,omitempty" gorm:"type:varchar(64);unique;not null;"`
	Password  string `json:"password,omitempty" gorm:"not null;"`

	Albums []*Album `json:"albums,omitempty"`
}

func (u *User) BeforeSave() {

	if &u.Password != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
		u.Password = string(hashedPassword)
	}

}

func (u *User) AfterFind() (err error) {
	fmt.Printf("Hop Didik --> %v\n", u)
	return nil
}
