package service

import (
	"context"
	"fmt"
	"server/exception"
	"server/internal/repository"
	"server/model/web"
	"server/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AnnouncementService interface {
	Create(ctx context.Context, annReq web.Announcement) (web.Announcement, error)
	FindAll(ctx context.Context, page int, pageSize int) ([]web.Announcement, error)
	FindById(ctx context.Context, id int) (web.Announcement, error)
	Update(ctx context.Context, annReq web.Announcement) (web.Announcement, error)
	Delete(ctx context.Context, id int) error
}

type announcementServiceImpl struct {
	DB                     *gorm.DB
	Validate               *validator.Validate
	AnnouncementRepository repository.AnnouncementRepository
}

func NewAnnouncementService(db *gorm.DB, validate *validator.Validate,
	announcementRepository repository.AnnouncementRepository) *announcementServiceImpl {
	return &announcementServiceImpl{
		DB:                     db,
		Validate:               validate,
		AnnouncementRepository: announcementRepository,
	}
}

func (s *announcementServiceImpl) Create(ctx context.Context, annReq web.Announcement) (web.Announcement, error) {
	var ann web.Announcement

	if err := s.Validate.Struct(annReq); err != nil {
		return ann, err
	}

	tx := s.DB.Begin()
	annDom, err := s.AnnouncementRepository.Save(ctx, tx, utils.AnnouncementWebToDomain(annReq))
	if err != nil {
		if err := tx.Rollback().Error; err != nil {
			return ann, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return ann, err
	}

	return utils.AnnouncementDomainToWeb(annDom), nil
}

func (s *announcementServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]web.Announcement, error) {
	var anns []web.Announcement

	tx := s.DB.Begin()
	annsDom, err := s.AnnouncementRepository.FindAll(ctx, tx, page, pageSize)
	if err != nil {
		if err := tx.Rollback().Error; err != nil {
			return anns, nil
		}
	}
	if err := tx.Commit().Error; err != nil {
		return anns, nil
	}

	for _, annDom := range annsDom {
		anns = append(anns, utils.AnnouncementDomainToWeb(annDom))
	}

	return anns, nil
}

func (s *announcementServiceImpl) FindById(ctx context.Context, id int) (web.Announcement, error) {
	var ann web.Announcement

	tx := s.DB.Begin()
	annDom, err := s.AnnouncementRepository.FindById(ctx, tx, id)
	if err != nil {
		return ann, err
	}

	ann = utils.AnnouncementDomainToWeb(annDom)
	return ann, nil
}

func (s *announcementServiceImpl) Update(ctx context.Context, annReq web.Announcement) (web.Announcement, error) {
	var ann web.Announcement

	if err := s.Validate.Struct(annReq); err != nil {
		return ann, err
	}

	tx := s.DB.Begin()
	if found, err := s.AnnouncementRepository.FindById(ctx, tx, int(annReq.ID)); err == nil && found.ID > 0 {
		fmt.Println(found)
		annDom, err := s.AnnouncementRepository.Save(ctx, tx, utils.AnnouncementWebToDomain(annReq))
		if err != nil {
			if err := tx.Rollback().Error; err != nil {
				return ann, err
			}
		}
		if err := tx.Commit().Error; err != nil {
			return ann, err
		}

		return utils.AnnouncementDomainToWeb(annDom), nil
	} else {
		return ann, &exception.NotFoundError{Entity: "announcement"}
	}
}

func (s *announcementServiceImpl) Delete(ctx context.Context, id int) error {
	tx := s.DB.Begin()
	if found, err := s.AnnouncementRepository.FindById(ctx, tx, id); err == nil && found.ID > 0 {
		err := s.AnnouncementRepository.Delete(ctx, tx, id)
		if err != nil {
			if err := tx.Rollback().Error; err != nil {
				return err
			}
		}
		if err := tx.Commit().Error; err != nil {
			return err
		}

		return nil
	} else {
		return &exception.NotFoundError{Entity: "announcement"}
	}
}
