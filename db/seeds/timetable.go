package seeds

import (
	"fmt"
	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
	"time"
)

func CreateTimetableSeeds(db *gorm.DB) {
	event := new(entity.Event)
	if err := db.First(event).Error; err != nil {
		fmt.Printf("Error finding event: %s", err)
	}

	var price int64 = 250000
	data := entity.Timetable{
		Name:        "Festival",
		Start:       time.Date(2024, 06, 22, 10, 0, 0, 0, time.Local),
		End:         time.Date(2024, 06, 22, 12, 0, 0, 0, time.Local),
		Description: nil,
		Stock:       1000,
		Price:       &price,
		EventID:     event.ID,
	}

	timetable := entity.NewTimetable(data.EventID,
		data.Name,
		data.Start,
		data.End,
		data.Description,
		data.Stock,
		data.Price)

	if err := db.Create(&timetable).Error; err != nil {
		fmt.Printf("Error when create timatable %s: %s\n", timetable.Name, err)
	}
}
