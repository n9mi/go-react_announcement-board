package utils

import "gorm.io/gorm"

func Paginate(page int, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 0 {
			page = 0
		}

		if pageSize < 0 {
			pageSize = 0
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
