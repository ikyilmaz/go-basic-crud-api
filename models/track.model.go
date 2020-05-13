package models

type Track struct {
	BaseModel

	Title string `json:"title,omitempty"     gorm:"type:varchar(128);not null;"`

	AlbumID int    `json:"albumId,omitempty"  gorm:"type:int;not null;"`
	Album   *Album `json:"album,omitempty"`

	OwnerID int   `json:"ownerId,omitempty"    gorm:"type:int;not null;"`
	Owner   *User `json:"owner,omitempty"`

	Artists []*User `json:"artists,omitempty"          gorm:"many2many:user_tracks;"`
}
