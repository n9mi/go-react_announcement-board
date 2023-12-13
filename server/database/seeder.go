package database

import (
	"context"
	"log"
	"server/internal/repository"
	"server/model/domain"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	var err error

	err = SeedAnnouncement(db)
	if err != nil {
		log.Fatal("ERROR: couldn't seed `announcements`")
		return err
	}

	return nil
}

func SeedAnnouncement(db *gorm.DB) error {
	anns := []domain.Announcement{
		{Title: "Announcement 1", Content: "Content of announcement 1"},
		{Title: "Announcement 2", Content: "Content of announcement 2"},
	}

	var err error
	ctx := context.Background()
	annRepository := repository.NewAnnouncementRepository()

	for i, ann := range anns {
		tx := db.Begin()
		result, _ := annRepository.Save(ctx, tx, ann)

		err = tx.Commit().Error
		if err != nil {
			break
		}

		anns[i] = result
	}

	return err
}
