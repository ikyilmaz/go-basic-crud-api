package models

type Album struct {
	BaseModel
	Name    string  `json:"name,omitempty"     gorm:"type:varchar(128);not null;"`
	Photo   string  `json:"photo,omitempty"    gorm:"default:'default.jpeg'"`
	OwnerID int     `json:"ownerId,omitempty"   gorm:"type:int;not null;"`
	Owner   *User   `json:"owner,omitempty"`
	Artists []*User `json:"artists,omitempty"  gorm:"many2many:user_albums;"`
}
