package repository

import (
	"jec-live-code/entity"

	"github.com/jmoiron/sqlx"
)

type NotificationRepository interface {
	Create(payload entity.InsertNotificationRequest)
}

type notificationRepository struct {
	*sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) Create(payload entity.InsertNotificationRequest) {
	tx := r.MustBegin()

	query := `INSERT INTO notifications (email, message, type, is_send) VALUES ($1, $2, $3, $4)`

	_, err := tx.Exec(query, payload.Email, payload.Message, payload.Type, false)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
}
