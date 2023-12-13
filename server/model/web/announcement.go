package web

import (
	"time"
)

type Announcement struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
