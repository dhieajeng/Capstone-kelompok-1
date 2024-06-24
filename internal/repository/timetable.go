package repository

import (
	"context"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"reflect"
)

type timetableRepository struct {
	db *gorm.DB
}

func (r *timetableRepository) Update(c context.Context, timetable *entity.Timetable) (*entity.Timetable, error) {
	var fields entity.Timetable

	val := reflect.ValueOf(timetable).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		if !field.IsZero() || !field.IsNil() {
			reflect.ValueOf(&fields).Elem().FieldByName(fieldName).Set(field)
		}
	}

	if err := r.db.WithContext(c).Model(&timetable).Where("id = ?", timetable.ID).Updates(fields).Error; err != nil {
		return nil, err
	}

	return timetable, nil
}

func (r *timetableRepository) FindById(c context.Context, id *uuid.UUID) (*entity.Timetable, error) {
	var timetable = new(entity.Timetable)
	if err := r.db.WithContext(c).Where("id = ?", id).Take(timetable).Error; err != nil {
		return nil, err
	}

	return timetable, nil
}

func (r *timetableRepository) FindByIds(c context.Context, ids []uuid.UUID) ([]entity.Timetable, error) {
	var timetables []entity.Timetable
	if err := r.db.WithContext(c).Where("id IN ?", ids).Find(&timetables).Error; err != nil {
		return nil, err
	}
	return timetables, nil
}

func (r *timetableRepository) UpdateStock(c context.Context, timetable *entity.Timetable) error {
	return r.db.WithContext(c).Save(timetable).Error
}

type TimetableRepository interface {
	FindById(c context.Context, id *uuid.UUID) (*entity.Timetable, error)
	FindByIds(c context.Context, ids []uuid.UUID) ([]entity.Timetable, error)
	Update(c context.Context, timetable *entity.Timetable) (*entity.Timetable, error)
	UpdateStock(c context.Context, timetable *entity.Timetable) error
}

func NewTimetableRepository(db *gorm.DB) TimetableRepository {
	return &timetableRepository{db: db}
}
