package utils

import (
	"server/model/domain"
	"server/model/web"
)

func AnnouncementDomainToWeb(annDom domain.Announcement) web.Announcement {
	return web.Announcement{
		ID:        annDom.ID,
		Title:     annDom.Title,
		Content:   annDom.Content,
		CreatedAt: annDom.CreatedAt,
	}
}

func AnnouncementWebToDomain(annWeb web.Announcement) domain.Announcement {
	return domain.Announcement{
		ID:      annWeb.ID,
		Title:   annWeb.Title,
		Content: annWeb.Content,
	}
}
