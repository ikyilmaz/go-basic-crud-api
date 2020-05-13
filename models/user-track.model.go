package models

type UserTrack struct {
	ID      int `json:"id,omitempty"      gorm:"primary_key"`
	UserID  int `json:"userId,omitempty"  gorm:"not null;type:int;"`
	TrackID int `json:"trackId,omitempty" gorm:"not null;type:int;"`

	User  *User  `json:"user"`
	Track *Track `json:"track"`
}
