package util

import (
	"fmt"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"gorm.io/gorm"
	"time"
)

func Filter(filter *binder.FilterRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.Keyword != nil {
			db = db.Where("name ILIKE ?", fmt.Sprint("%", *filter.Keyword, "%"))
		}
		if filter.Topic != nil {
			db = db.Where("topic_id = ?", filter.Topic)
		}
		if filter.Category != nil {
			db = db.Where("category_id = ?", filter.Category)
		}
		if filter.Location != nil {
			db = db.Where("location_id = ?", filter.Location)
		}
		if filter.IsPaid != nil {
			db = db.Where("is_paid = ?", filter.IsPaid)
		}

		if filter.Time != nil {
			db = applyTimeFilter(db, *filter.Time)
		}

		return db
	}
}

func applyTimeFilter(db *gorm.DB, timeFilter binder.Time) *gorm.DB {
	now := time.Now()
	startOfDay := now.Truncate(24 * time.Hour)

	switch timeFilter {
	case binder.Today:
		db = db.Where("start >= ? AND start < ?", startOfDay, startOfDay.Add(24*time.Hour))
	case binder.Tomorrow:
		tomorrow := startOfDay.Add(24 * time.Hour)
		db = db.Where("start >= ? AND start < ?", tomorrow, tomorrow.Add(24*time.Hour))
	case binder.ThisWeek:
		startOfWeek := startOfDay.AddDate(0, 0, -int(startOfDay.Weekday()))
		endOfWeek := startOfWeek.AddDate(0, 0, 7)
		db = db.Where("start >= ? AND start < ?", startOfWeek, endOfWeek)
	case binder.NextWeek:
		startOfNextWeek := startOfDay.AddDate(0, 0, 7-int(startOfDay.Weekday()))
		endOfNextWeek := startOfNextWeek.AddDate(0, 0, 7)
		db = db.Where("start >= ? AND start < ?", startOfNextWeek, endOfNextWeek)
	case binder.ThisMonth:
		startOfMonth := startOfDay.AddDate(0, 0, -startOfDay.Day()+1)
		endOfMonth := startOfMonth.AddDate(0, 1, -1)
		db = db.Where("start >= ? AND start < ?", startOfMonth, endOfMonth)
	case binder.NextMonth:
		startOfNextMonth := startOfDay.AddDate(0, 1, -startOfDay.Day()+1)
		endOfNextMonth := startOfNextMonth.AddDate(0, 1, -1)
		db = db.Where("start >= ? AND start < ?", startOfNextMonth, endOfNextMonth)
	}

	return db
}
