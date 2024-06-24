package repository

import (
	"context"
	"github.com/google/uuid"
	"reflect"

	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(c context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(c).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindById(c context.Context, id uuid.UUID) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(c).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(c context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(c).Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Edit(c context.Context, user *entity.User) (*entity.User, error) {
	var fields entity.User

	val := reflect.ValueOf(user).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		if !field.IsZero() || !field.IsNil() {
			reflect.ValueOf(&fields).Elem().FieldByName(fieldName).Set(field)
		}
	}

	if err := r.db.WithContext(c).Model(&user).Where("id = ?", user.ID).Updates(fields).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(c context.Context, user *entity.User) error {
	if err := r.db.WithContext(c).Delete(&user, "id = ?", user.ID).Error; err != nil {
		return err
	}
	return nil
}

type UserRepository interface {
	Create(c context.Context, user *entity.User) (*entity.User, error)
	FindById(c context.Context, id uuid.UUID) (*entity.User, error)
	FindByEmail(c context.Context, email string) (*entity.User, error)
	Edit(c context.Context, user *entity.User) (*entity.User, error)
	Delete(c context.Context, user *entity.User) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
