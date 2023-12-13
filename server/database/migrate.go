package database

import (
	"log"
	"server/model/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	var err error

	err = db.AutoMigrate(&domain.Announcement{})
	if err != nil {
		log.Fatal("ERROR: couldn't migrate `announcements`")
		return err
	}

	return err
}
