package domain

import (
	"time"
)

type Announcement struct {
	ID        uint64 `gorm:"primaryKey"`
	Title     string
	Content   string
	CreatedAt time.Time `gorm:"<-:create"`
}
