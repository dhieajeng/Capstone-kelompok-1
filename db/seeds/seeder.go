package seeds

import (
	"gorm.io/gorm"
	"time"
)

func Run(db *gorm.DB) {
	CreateUserSeeds(db)
	CreateLocationSeeds(db)
	CreateCategorySeeds(db)
	CreateTopicSeeds(db)
	CreateEventSeeder(db)
	CreateTimetableSeeds(db)
}

func ParseDate(date string) *time.Time {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil
	}
	return &parsedDate
}

func HandleTimeDereference(data *time.Time) time.Time {
	t := data
	return *t
}

func HandleTimeReference(data time.Time) *time.Time {
	t := data
	return &t
}
