package binder

import (
	"github.com/google/uuid"
)

type CreateEventRequest struct {
	Name             string  `form:"name" json:"name" validate:"required"`
	LocationID       int64   `form:"location_id" json:"location_id" validate:"required"`
	CategoryID       int64   `form:"category_id" json:"category_id" validate:"required"`
	TopicID          int64   `form:"topic_id" json:"topic_id" validate:"required"`
	Start            string  `form:"start" json:"start" validate:"required"`
	End              string  `form:"end" json:"end" validate:"required"`
	Address          string  `form:"address" json:"address" validate:"required"`
	AddressLink      string  `form:"address_link" json:"address_link" validate:"required"`
	Organizer        string  `form:"organizer" json:"organizer" validate:"required"`
	Description      string  `form:"description" json:"description" validate:"required"`
	TermAndCondition string  `form:"term_and_condition" json:"term_and_condition" validate:"required"`
	Cover            *string `form:"cover" json:"cover"`
	OrganizerLogo    *string `form:"organizer_logo" json:"organizer_logo"`
	IsPaid           bool    `form:"is_paid" json:"is_paid" validate:"required"`
	IsPublic         bool    `form:"is_public" json:"is_public" validate:"required"`
}

type EditEventRequest struct {
	ID               uuid.UUID `param:"id" json:"id" validate:"required"`
	Name             string    `form:"name" json:"name" validate:"required"`
	LocationID       int64     `form:"location_id" json:"location_id" validate:"required"`
	CategoryID       int64     `form:"category_id" json:"category_id" validate:"required"`
	TopicID          int64     `form:"topic_id" json:"topic_id" validate:"required"`
	Start            string    `form:"start" json:"start" validate:"required"`
	End              string    `form:"end" json:"end" validate:"required"`
	Address          string    `form:"address" json:"address" validate:"required"`
	AddressLink      string    `form:"address_link" json:"address_link" validate:"required"`
	Organizer        string    `form:"organizer" json:"organizer" validate:"required"`
	Description      string    `form:"description" json:"description" validate:"required"`
	TermAndCondition string    `form:"term_and_condition" json:"term_and_condition" validate:"required"`
	Cover            *string   `form:"cover" json:"cover"`
	OrganizerLogo    *string   `form:"organizer_logo" json:"organizer_logo"`
	IsPaid           bool      `form:"is_paid" json:"is_paid" validate:"required"`
	IsPublic         bool      `form:"is_public" json:"is_public" validate:"required"`
}

type ApproveEventRequest struct {
	ID        uuid.UUID `param:"id" json:"id" validate:"required"`
	IsApprove bool      `json:"is_Approve" validate:"required"`
}
