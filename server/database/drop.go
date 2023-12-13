package database

import (
	"log"
	"server/model/domain"

	"gorm.io/gorm"
)

func Drop(db *gorm.DB) error {
	var err error

	err = db.Migrator().DropTable(&domain.Announcement{})
	if err != nil {
		log.Fatal("ERROR: couldn't drop `announcements`")
	}

	return err
}
