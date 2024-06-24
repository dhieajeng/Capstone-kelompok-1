package repository

import (
	"context"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"reflect"
)

type eventRepository struct {
	db *gorm.DB
}

func (r *eventRepository) GetAll(c context.Context) ([]entity.Event, error) {
	events := make([]entity.Event, 0)
	err := r.db.WithContext(c).
		Where("is_public = ? AND is_approved = ?", true, true).Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *eventRepository) GetAllWithPaginateAndFilter(c context.Context,
	paginate binder.PaginateRequest,
	filter binder.FilterRequest,
	sort binder.SortRequest) ([]entity.Event, int64, error) {
	var totalItems int64
	events := make([]entity.Event, 0)

	err := r.db.WithContext(c).
		Scopes(util.Filter(&filter)).
		Model(&entity.Event{}).
		Where("is_public = ? AND is_approved = ?", true, true).
		Count(&totalItems).Error

	if err != nil {
		return nil, 0, err
	}
	if int(totalItems) <= 0 {
		return nil, 0, nil
	}

	err = r.db.WithContext(c).
		Scopes(util.Paginate(*paginate.Page, *paginate.Limit), util.Filter(&filter), util.Sort(&sort)).
		Where("is_public = ? AND is_approved = ?", true, true).Find(&events).Error
	if err != nil {
		return nil, 0, err
	}

	return events, totalItems, nil
}

func (r *eventRepository) Create(c context.Context, event *entity.Event) (*entity.Event, error) {
	if err := r.db.WithContext(c).Create(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r *eventRepository) FindById(c context.Context, id uuid.UUID) (*entity.Event, error) {
	event := new(entity.Event)
	if err := r.db.WithContext(c).
		Where("id = ?", id).
		Preload("User").
		Preload("Location").
		Preload("Category").
		Preload("Topic").
		Preload("Timetables").
		Take(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r *eventRepository) FindWithDetailById(c context.Context, id uuid.UUID) (*entity.Event, error) {
	event := new(entity.Event)
	if err := r.db.WithContext(c).
		Where("id = ?", id).
		Preload("User").
		Preload("Location").
		Preload("Category").
		Preload("Topic").
		Preload("Timetables").
		Preload("Timetables.Tickets", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN transactions ON transactions.id = tickets.transaction_id").Where("transactions.status = ?", "paid")
		}).
		Take(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r *eventRepository) Update(c context.Context, event *entity.Event) (*entity.Event, error) {
	var fields entity.Event
	val := reflect.ValueOf(event).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		if !field.IsZero() {
			reflect.ValueOf(&fields).Elem().FieldByName(fieldName).Set(field)
		}
	}

	if err := r.db.WithContext(c).Model(&event).Where("id = ?", event.ID).Updates(fields).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) Delete(c context.Context, event *entity.Event) error {
	if err := r.db.WithContext(c).Delete(&event).Error; err != nil {
		return err
	}

	return nil
}

type EventRepository interface {
	GetAll(c context.Context) ([]entity.Event, error)
	GetAllWithPaginateAndFilter(c context.Context, paginate binder.PaginateRequest, filter binder.FilterRequest, sort binder.SortRequest) ([]entity.Event, int64, error)
	Create(c context.Context, event *entity.Event) (*entity.Event, error)
	FindById(c context.Context, id uuid.UUID) (*entity.Event, error)
	FindWithDetailById(c context.Context, id uuid.UUID) (*entity.Event, error)
	Update(c context.Context, event *entity.Event) (*entity.Event, error)
	Delete(c context.Context, event *entity.Event) error
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}
