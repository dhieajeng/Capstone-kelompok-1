package repository

import (
	"context"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tokenRepository struct {
	db *gorm.DB
}

func (r *tokenRepository) Create(c context.Context, token *entity.Token) (*entity.Token, error) {
	if err := r.db.WithContext(c).Create(&token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (r *tokenRepository) FindById(c context.Context, id uuid.UUID) (*entity.Token, error) {
	var tokenData entity.Token
	if err := r.db.WithContext(c).Where(
		"id = ? AND expires_at > NOW()",
		id).First(&tokenData).Error; err != nil {
		return nil, err
	}

	return &tokenData, nil

}

func (r *tokenRepository) FindByIdAndEmail(c context.Context, id uuid.UUID, email string) (*entity.Token, error) {
	var tokenData entity.Token
	if err := r.db.WithContext(c).Where(
		"id = ? AND email = ? AND expires_at > NOW()",
		id, email).First(&tokenData).Error; err != nil {
		return nil, err
	}

	return &tokenData, nil

}

func (r *tokenRepository) Delete(c context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(c).Where("id = ?", id).Delete(&entity.Token{}).Error; err != nil {
		return err
	}
	return nil
}

type TokenRepository interface {
	Create(c context.Context, token *entity.Token) (*entity.Token, error)
	FindById(c context.Context, id uuid.UUID) (*entity.Token, error)
	FindByIdAndEmail(c context.Context, id uuid.UUID, email string) (*entity.Token, error)
	Delete(c context.Context, id uuid.UUID) error
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{
		db: db,
	}
}
