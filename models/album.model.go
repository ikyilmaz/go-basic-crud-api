package models

type Album struct {
	BaseModel
	Name  string `json:"name,omitempty"   gorm:"type:varchar(128);not null;"`
	Photo string `json:"photo,omitempty"  gorm:"default:'default.jpeg'"`
	//UserID uint   `json:"userId,omitempty" gorm:"type:int;not null;"`
	Users []*User `json:"user,omitempty"  gorm:"many2many:users_albums;"`
}
