package repository

import (
	"context"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type oneTimePasswordRepository struct {
	db *gorm.DB
}

func (r *oneTimePasswordRepository) Create(c context.Context, otp *entity.OneTimePassword) (*entity.OneTimePassword, error) {
	if err := r.db.WithContext(c).Create(&otp).Error; err != nil {
		return otp, err
	}

	return otp, nil
}

func (r *oneTimePasswordRepository) FindOneByCodeAndEmail(c context.Context, email, code string) (*entity.OneTimePassword, error) {
	otp := new(entity.OneTimePassword)
	if err := r.db.WithContext(c).Where(
		"email = ? AND otp_code = ? AND expires_at > NOW()",
		email, code).Take(&otp).Error; err != nil {
		return otp, err
	}

	return otp, nil
}

func (r *oneTimePasswordRepository) Delete(c context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(c).Delete(&entity.OneTimePassword{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

type OneTimePasswordRepository interface {
	Create(c context.Context, otp *entity.OneTimePassword) (*entity.OneTimePassword, error)
	FindOneByCodeAndEmail(c context.Context, email, code string) (*entity.OneTimePassword, error)
	Delete(c context.Context, id uuid.UUID) error
}

func NewOneTimePasswordRepository(db *gorm.DB) OneTimePasswordRepository {
	return &oneTimePasswordRepository{db}
}
