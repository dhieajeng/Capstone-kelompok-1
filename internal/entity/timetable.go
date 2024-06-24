package entity

import (
	"github.com/google/uuid"
	"time"
)

type Timetable struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Description *string   `json:"description,omitempty"`
	Stock       int32     `json:"stock"`
	Price       *int64    `json:"price,omitempty"`
	EventID     uuid.UUID `json:"-"`
	Event       *Event    `json:"event,omitempty"`
	Tickets     *[]Ticket `gorm:"TimetableID;references:ID" json:"tickets,omitempty"`
	Auditable
}

func NewTimetable(eventId uuid.UUID, name string, start, end time.Time, description *string, stock int32, price *int64) *Timetable {
	return &Timetable{
		ID:          uuid.New(),
		EventID:     eventId,
		Name:        name,
		Start:       start,
		End:         end,
		Description: description,
		Stock:       stock,
		Price:       price,
	}
}

func EditTimetable(id, eventId uuid.UUID, name string, start, end time.Time, description *string) *Timetable {
	return &Timetable{
		ID:          id,
		EventID:     eventId,
		Name:        name,
		Start:       start,
		End:         end,
		Description: description,
	}
}

func UpdateStock(id uuid.UUID, stock int32) *Timetable {
	return &Timetable{
		ID:    id,
		Stock: stock,
	}
}
