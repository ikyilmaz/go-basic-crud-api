package models

type UserAlbum struct {
	ID      int `json:"id,omitempty"      gorm:"primary_key"`
	UserID  int `json:"userId,omitempty"  gorm:"not null;type:int;"`
	AlbumID int `json:"albumId,omitempty" gorm:"not null;type:int;"`

	User  *User  `json:"user"`
	Album *Album `json:"album"`
}

func (u UserAlbum) TableName() string {
	return "users_albums"
}
