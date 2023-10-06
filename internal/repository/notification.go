package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/ewallet/domain"
)

type notificationRepository struct {
	db *goqu.Database
}

func NewNotification(con *sql.DB) domain.NotificationRepository {
	return &notificationRepository{
		db: goqu.New("default", con),
	}
}

func (n notificationRepository) FindByUser(ctx context.Context, user int64) (notification []domain.Notification, err error) {
	dataset := n.db.From("notifications").Where(goqu.Ex{
		"user_id": user,
	}).Order(goqu.I("created_at").Desc()).Limit(15)

	err = dataset.ScanStructsContext(ctx, &notification)
	return
}

func (n notificationRepository) Insert(ctx context.Context, notification *domain.Notification) error {
	executor := n.db.Insert("notifications").Rows(goqu.Record{
		"user_id":    notification.UserID,
		"title":      notification.Title,
		"body":       notification.Body,
		"status":     notification.Status,
		"is_read":    notification.IsRead,
		"created_at": notification.CreatedAt,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, notification)
	return err
}

func (n notificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	executor := n.db.Update("notifications").Where(goqu.Ex{
		"id": notification.ID,
	}).Set(goqu.Record{
		"title":   notification.Title,
		"body":    notification.Body,
		"status":  notification.Status,
		"is_read": notification.IsRead,
	}).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
