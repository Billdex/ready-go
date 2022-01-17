package gormutil

import (
	"gorm.io/gorm"
	"ready-go/pkg/util/numutil"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page = numutil.IntLeast(page, 1)
		pageSize = numutil.IntBetween(pageSize, 1, 100)

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
