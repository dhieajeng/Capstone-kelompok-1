package entity

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID               uuid.UUID    `json:"id"`
	Name             string       `json:"name"`
	Start            time.Time    `json:"start"`
	End              time.Time    `json:"end"`
	Address          string       `json:"address"`
	AddressLink      string       `json:"address_link"`
	Organizer        string       `json:"organizer"`
	OrganizerLogo    *string      `json:"organizer_logo,omitempty"`
	Cover            *string      `json:"cover,omitempty"`
	Description      string       `json:"description"`
	TermAndCondition string       `json:"term_and_condition"`
	IsPaid           bool         `json:"is_paid"`
	IsPublic         bool         `json:"is_public"`
	IsApproved       bool         `json:"is_approved"`
	ApprovedAt       *time.Time   `json:"approved_at,omitempty"`
	UserID           uuid.UUID    `json:"-"`
	User             *User        `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	LocationID       int64        `json:"-"`
	Location         *Location    `gorm:"foreignKey:LocationID;references:ID" json:"location,omitempty"`
	CategoryID       int64        `json:"-"`
	Category         *Category    `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`
	TopicID          int64        `json:"-"`
	Topic            *Topic       `gorm:"foreignKey:TopicID;references:ID" json:"topic,omitempty"`
	Timetables       *[]Timetable `gorm:"foreignKey:EventID;references:ID" json:"timetables,omitempty"`
	Auditable
}

// NewEventParams Struct untuk menyimpan parameter NewEvent
type NewEventParams struct {
	Name             string
	UserID           uuid.UUID
	LocationID       int64
	CategoryID       int64
	TopicID          int64
	Start            time.Time
	End              time.Time
	Address          string
	AddressLink      string
	Organizer        string
	Description      string
	Cover            *string
	OrganizerLogo    *string
	TermAndCondition string
	IsPaid           bool
	IsPublic         bool
	IsApproved       bool
}

func NewEvent(params NewEventParams) *Event {
	var approvedTime *time.Time = nil
	if params.IsApproved {
		now := time.Now()
		approvedTime = &now
	}
	return &Event{
		ID:               uuid.New(),
		Name:             params.Name,
		UserID:           params.UserID,
		LocationID:       params.LocationID,
		CategoryID:       params.CategoryID,
		TopicID:          params.TopicID,
		Start:            params.Start,
		End:              params.End,
		Address:          params.Address,
		AddressLink:      params.AddressLink,
		Organizer:        params.Organizer,
		Description:      params.Description,
		TermAndCondition: params.TermAndCondition,
		Cover:            params.Cover,
		OrganizerLogo:    params.OrganizerLogo,
		IsPaid:           params.IsPaid,
		IsPublic:         params.IsPublic,
		IsApproved:       params.IsApproved,
		ApprovedAt:       approvedTime,
	}
}

func ApproveEvent(id uuid.UUID) *Event {
	now := time.Now()
	return &Event{
		ID:         id,
		ApprovedAt: &now,
		IsApproved: true,
	}
}

type EditEventParams struct {
	ID               uuid.UUID
	Name             string
	LocationID       int64
	CategoryID       int64
	TopicID          int64
	Start            time.Time
	End              time.Time
	Address          string
	AddressLink      string
	Organizer        string
	Description      string
	TermAndCondition string
	Cover            *string
	OrganizerLogo    *string
	IsPaid           bool
	IsPublic         bool
}

func EditEvent(params EditEventParams) *Event {
	return &Event{
		ID:               params.ID,
		Name:             params.Name,
		LocationID:       params.LocationID,
		CategoryID:       params.CategoryID,
		TopicID:          params.TopicID,
		Start:            params.Start,
		End:              params.End,
		Address:          params.Address,
		AddressLink:      params.AddressLink,
		Organizer:        params.Organizer,
		Description:      params.Description,
		TermAndCondition: params.TermAndCondition,
		Cover:            params.Cover,
		OrganizerLogo:    params.OrganizerLogo,
		IsPaid:           params.IsPaid,
		IsPublic:         params.IsPublic,
	}
}
