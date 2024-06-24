package util

import (
	"github.com/bloomingbug/depublic/internal/http/binder"
	"gorm.io/gorm"
)

func Sort(sort *binder.SortRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sort.Sort != nil {
			var sort = getSortOrder(*sort.Sort)
			db = db.Order(sort)
		}

		return db
	}
}

func getSortOrder(sort string) string {
	switch sort {
	case "naz":
		return "name asc"
	case "nza":
		return "name desc"
	case "daz":
		return "start asc"
	case "dza":
		return "start desc"
	default:
		return "start asc" // Default sorting
	}
}
