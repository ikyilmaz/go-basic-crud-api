package models

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-rest-api/lib"
	"log"
	"time"
)

func NewDB() *gorm.DB {
	var db, err = gorm.Open("mysql", connectionString)

	lib.CheckErr(err)

	sync(db, SyncOptions{
		Force:   true,
		LogMode: false,
	})

	return db
}

var (
	config           = mysql.Config{User: "root", ParseTime: true, DBName: "music"}
	connectionString = config.FormatDSN()
)

type SyncOptions struct {
	Force   bool
	LogMode bool
}

func sync(db *gorm.DB, options SyncOptions) {
	var album Album
	var user User
	var userAlbum UserAlbum
	var track Track
	var userTrack UserTrack

	err := db.Transaction(func(tx *gorm.DB) error {
		tx.LogMode(options.LogMode)

		if options.Force {
			tx.
				DropTableIfExists(&userTrack, &track, &userAlbum, &album, &user)
		}

		tx.
			AutoMigrate(&userAlbum, &userTrack, &user, &album, &track)

		tx.
			Model(&album).
			AddForeignKey("owner_id", "users(id)", "CASCADE", "RESTRICT")

		tx.
			Model(&userAlbum).
			AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT").
			AddForeignKey("album_id", "albums(id)", "CASCADE", "RESTRICT")

		tx.
			Model(&track).
			AddForeignKey("owner_id", "users(id)", "CASCADE", "RESTRICT").
			AddForeignKey("album_id", "albums(id)", "CASCADE", "RESTRICT")

		tx.
			Model(&userTrack).
			AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT").
			AddForeignKey("track_id", "tracks(id)", "CASCADE", "RESTRICT")

		var john = User{
			FirstName: "john",
			LastName:  "doe",
			Username:  "john_doe",
			Email:     "john@example.com",
			Password:  "87654321",
			Role:      "admin",
		}

		var charles = User{
			FirstName: "charles",
			LastName:  "dickens",
			Username:  "charles_dickens",
			Email:     "charles@example.com",
			Password:  "87654321",
		}

		var rick = User{
			FirstName: "rick",
			LastName:  "grimes",
			Username:  "rick_grimes",
			Email:     "rick@example.com",
			Password:  "87654321",
		}

		var err error

		err = tx.
			Create(&john).
			Create(&charles).
			Create(&rick).Error

		if err != nil {
			return errors.New("createUser --> " + err.Error())
		}

		var johnsAlbum = Album{
			Name:    "John is the best",
			OwnerID: john.ID,
		}

		var charlesAlbum = Album{
			Name:    "Charles is the best",
			OwnerID: charles.ID,
		}

		var ricksAlbum = Album{
			Name:    "Rick is the best",
			OwnerID: rick.ID,
		}

		err = tx.
			Create(&johnsAlbum).
			Create(&charlesAlbum).
			Create(&ricksAlbum).Error

		if err != nil {
			return errors.New("createAlbum --> " + err.Error())
		}

		var firstTrackForJohnsAlbum = Track{
			Title:   "First Track For John's Album",
			AlbumID: johnsAlbum.ID,
			OwnerID: john.ID,
		}

		var secondTrackForJohnsAlbum = Track{
			Title:   "Second Track For John's Album",
			AlbumID: johnsAlbum.ID,
			OwnerID: john.ID,
		}

		var firstTrackForCharlesAlbum = Track{
			Title:   "First Track For Charles's Album",
			AlbumID: charlesAlbum.ID,
			OwnerID: charles.ID,
		}

		var secondTrackForCharlesAlbum = Track{
			Title:   "Second Track For Charles's Album",
			AlbumID: charlesAlbum.ID,
			OwnerID: charles.ID,
		}

		var firstTrackForRicksAlbum = Track{
			Title:   "First Track For Rick's Album",
			AlbumID: ricksAlbum.ID,
			OwnerID: rick.ID,
		}

		var secondTrackForRicksAlbum = Track{
			Title:   "Second Track For Rick's Album",
			AlbumID: ricksAlbum.ID,
			OwnerID: rick.ID,
		}

		err = tx.
			Create(&firstTrackForJohnsAlbum).
			Create(&firstTrackForCharlesAlbum).
			Create(&firstTrackForRicksAlbum).
			Create(&secondTrackForJohnsAlbum).
			Create(&secondTrackForCharlesAlbum).
			Create(&secondTrackForRicksAlbum).
			Error

		if err != nil {
			return errors.New("createTrack --> " + err.Error())
		}

		err = tx.
			Model(johnsAlbum).
			Association("Artists").
			Append(rick).
			Error

		if err != nil {
			return errors.New("Association For John's Album --> " + err.Error())
		}

		err = tx.
			Model(charlesAlbum).
			Association("Artists").
			Append(rick).
			Error

		if err != nil {
			return errors.New("Association For Charles's Album --> " + err.Error())
		}

		err = tx.
			Model(firstTrackForJohnsAlbum).
			Association("Artists").
			Append(rick).
			Error

		if err != nil {
			return errors.New("Association For John's Track --> " + err.Error())
		}

		err = tx.
			Model(secondTrackForRicksAlbum).
			Association("Artists").
			Append(john).
			Error

		if err != nil {
			return errors.New("Association For Rick's Track --> " + err.Error())
		}

		err = tx.
			Model(firstTrackForCharlesAlbum).
			Association("Artists").
			Append(rick).
			Error

		if err != nil {
			return errors.New("Association For Charles's Track --> " + err.Error())
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}

type BaseModel struct {
	ID        int        `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
}
