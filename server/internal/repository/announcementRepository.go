package repository

import (
	"context"
	"server/exception"
	"server/model/domain"
	"server/utils"

	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	Save(ctx context.Context, tx *gorm.DB, ann domain.Announcement) (domain.Announcement, error)
	FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]domain.Announcement, error)
	FindById(ctx context.Context, tx *gorm.DB, id int) (domain.Announcement, error)
	Delete(ctx context.Context, tx *gorm.DB, id int) error
}

func NewAnnouncementRepository() *announcementRepositoryImpl {
	return &announcementRepositoryImpl{}
}

type announcementRepositoryImpl struct {
}

func (r *announcementRepositoryImpl) Save(ctx context.Context, tx *gorm.DB,
	ann domain.Announcement) (domain.Announcement, error) {
	tx = tx.WithContext(ctx)

	if err := tx.Save(&ann).Error; err != nil {
		return ann, err
	}

	return ann, nil
}

func (r *announcementRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB,
	page int, pageSize int) ([]domain.Announcement, error) {
	tx = tx.WithContext(ctx)
	var anns []domain.Announcement
	if page > 0 && pageSize > 0 {
		tx = tx.Scopes(utils.Paginate(page, pageSize))
	}

	if err := tx.Find(&anns).Error; err != nil {
		return anns, err
	}

	return anns, nil
}

func (r *announcementRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB,
	id int) (domain.Announcement, error) {
	tx = tx.WithContext(ctx)
	var ann domain.Announcement

	if err := tx.First(&ann, id).Error; err != nil {
		return ann, err
	}

	if ann.ID == 0 {
		return ann, &exception.NotFoundError{Entity: "announcement"}
	}

	return ann, nil
}

func (r *announcementRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB,
	id int) error {
	tx = tx.WithContext(ctx)

	if err := tx.Delete(&domain.Announcement{}, id).Error; err != nil {
		return err
	}

	return nil
}
