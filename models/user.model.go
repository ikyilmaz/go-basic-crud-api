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

	// ONE 2 MANY
	AlbumsOwned []*Album `json:"albumsOwned,omitempty" gorm:"foreignkey:OwnerID"`
	TracksOwned []*Track `json:"tracksOwned,omitempty" gorm:"foreignkey:OwnerID"`

	// MANY 2 MANY
	AlbumsParticipated []*Album `json:"albumsParticipated,omitempty"     gorm:"many2many:user_albums;"`
	TracksParticipated []*Track `json:"tracksParticipated,omitempty"     gorm:"many2many:user_tracks;"`
}

func (u *User) BeforeSave() {

	if &u.Password != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
		u.Password = string(hashedPassword)
	}

}
