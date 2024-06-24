package util

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	TotalItems int         `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Items      interface{} `json:"items"`
}

// NewPagination initializes a new Pagination struct
func NewPagination(limit int, page int, totalItems int, totalPages int, items interface{}) *Pagination {
	return &Pagination{
		Limit:      limit,
		Page:       page,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Items:      items,
	}
}

// Response returns the pagination response in echo.Map format
func (p *Pagination) Response() map[string]interface{} {
	return echo.Map{
		"limit":       p.Limit,
		"page":        p.Page,
		"total_items": p.TotalItems,
		"total_pages": p.TotalPages,
		"items":       p.Items,
	}
}

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
