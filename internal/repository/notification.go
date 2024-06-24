package repository

import (
	"context"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func (r *notificationRepository) GetByUserIdWithPagination(c context.Context,
	id uuid.UUID,
	paginate binder.PaginateRequest,
	isRead *bool) ([]entity.Notification, int64, error) {
	var totalItems int64
	notifications := make([]entity.Notification, 0)

	queryCount := r.db.WithContext(c).
		Model(&entity.Notification{}).
		Where("user_id = ?", id)
	if isRead != nil {
		queryCount = queryCount.Where("is_read = ?", isRead)
	}
	err := queryCount.Count(&totalItems).Error

	if err != nil || int(totalItems) <= 0 {
		return nil, 0, err
	}

	queryNotif := r.db.WithContext(c).
		Scopes(util.Paginate(*paginate.Page, *paginate.Limit)).
		Select("id", "title", "is_read", "user_id").
		Where("user_id = ?", id)
	if isRead != nil {
		queryNotif = queryNotif.Where("is_read = ?", isRead)
	}

	err = queryNotif.Order("created_at desc").Find(&notifications).Error
	if err != nil {
		return nil, 0, err
	}

	return notifications, totalItems, nil
}

func (r *notificationRepository) Create(c context.Context, notif *entity.Notification) (*entity.Notification, error) {
	if err := r.db.WithContext(c).Create(notif).Error; err != nil {
		return nil, err
	}

	return notif, nil
}

func (r *notificationRepository) FindById(c context.Context, id uuid.UUID) (*entity.Notification, error) {
	notification := new(entity.Notification)
	if err := r.db.WithContext(c).Where("id = ?", id).Take(&notification).Error; err != nil {
		return nil, err
	}

	return notification, nil
}

func (r *notificationRepository) Edit(c context.Context, notif *entity.Notification) error {
	fields := make(map[string]interface{})
	fields["is_read"] = notif.IsRead

	if err := r.db.WithContext(c).Model(notif).Where("id = ?", notif.ID).Updates(fields).Error; err != nil {
		return err
	}

	return nil
}

type NotificationRepository interface {
	GetByUserIdWithPagination(c context.Context, id uuid.UUID, paginate binder.PaginateRequest, isRead *bool) ([]entity.Notification, int64, error)
	Create(c context.Context, notif *entity.Notification) (*entity.Notification, error)
	FindById(c context.Context, id uuid.UUID) (*entity.Notification, error)
	Edit(c context.Context, notif *entity.Notification) error
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}
