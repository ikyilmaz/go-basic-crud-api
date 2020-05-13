package models

type UserAlbum struct {
	ID      int `json:"id,omitempty"      gorm:"primary_key"`
	UserID  int `json:"userId,omitempty"  gorm:"not null;"`
	AlbumID int `json:"albumId,omitempty" gorm:"not null;"`

	User  *User  `json:"user"`
	Album *Album `json:"album"`
}
