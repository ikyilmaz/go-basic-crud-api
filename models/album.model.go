package models

type Album struct {
	Name   string `json:"omitempty,name" gorm:"type:varchar(128);not null;"`
	Photo  string `json:"omitempty,photo" gorm:"default:'default.jpeg'"`
	UserID uint   `json:"omitempty,userId" gorm:"type:int;not null;"`
	User   *User  `json:"omitempty,user"`
}
