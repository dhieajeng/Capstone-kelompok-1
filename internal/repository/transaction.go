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

type transactionRepository struct {
	db *gorm.DB
}

func (r *transactionRepository) Create(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	if err := r.db.WithContext(c).Create(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) FindById(c context.Context, id uuid.UUID) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	if err := r.db.WithContext(c).Where("id = ?", id).Take(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *transactionRepository) FindByIdWithDetails(c context.Context,
	id uuid.UUID,
	paginate binder.PaginateRequest) ([]entity.Transaction, int64, error) {
	var totalItems int64
	transactions := make([]entity.Transaction, 0)

	err := r.db.WithContext(c).
		Model(&entity.Transaction{}).
		Where("user_id = ?", id).
		Count(&totalItems).Error
	if err != nil || int(totalItems) <= 0 {
		return nil, 0, err
	}

	err = r.db.WithContext(c).
		Scopes(util.Paginate(*paginate.Page, *paginate.Limit)).
		Where("user_id = ?", id).
		Preload("Tickets.Timetable.Event").
		Order("created_at desc").Find(&transactions).Error

	if err != nil {
		return nil, 0, err
	}
	return transactions, totalItems, nil
}

func (r *transactionRepository) FindByInvoice(c context.Context, invoice string) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	if err := r.db.WithContext(c).Where("invoice = ?", invoice).Take(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *transactionRepository) Edit(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	var fields entity.Transaction

	val := reflect.ValueOf(transaction).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		if !field.IsZero() {
			reflect.ValueOf(&fields).Elem().FieldByName(fieldName).Set(field)
		}
	}

	if err := r.db.WithContext(c).Model(&transaction).Where("id = ?", transaction.ID).Updates(fields).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

type TransactionRepository interface {
	Create(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	FindById(c context.Context, id uuid.UUID) (*entity.Transaction, error)
	FindByInvoice(c context.Context, invoice string) (*entity.Transaction, error)
	FindByIdWithDetails(c context.Context, id uuid.UUID, paginate binder.PaginateRequest) ([]entity.Transaction, int64, error)
	Edit(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}
