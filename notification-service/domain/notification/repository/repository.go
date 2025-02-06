package repository

import (
	"fmt"
	"log"
	"notification-service/domain/notification/entity"

	"github.com/jmoiron/sqlx"
)

type NotificationRepository interface {
	CreateNotification(payload entity.InsertNotificationRequest) (*entity.Notification, error)
	GetUnsendNotification() ([]entity.Notification, error)
	UpdateIsSendNotification(id int) error
}

type notificationRepository struct {
	*sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) CreateNotification(payload entity.InsertNotificationRequest) (*entity.Notification, error) {
	var notification entity.Notification

	tx := r.MustBegin()

	query := `INSERT INTO notifications (email, message, type, is_send, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := tx.QueryRow(query, payload.Email, payload.Message, payload.Type, false, "now()", "now()").Scan(&notification.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// getQuery := `SELECT id, email, message, type, is_send, created_at, updated_at FROM notifications WHERE id = $1`
	// err = tx.Get(&notification, getQuery, notification.ID)
	// if err != nil {
	// 	return nil, err
	// }

	tx.Commit()

	return &notification, nil
}

func (r *notificationRepository) GetUnsendNotification() ([]entity.Notification, error) {
	notification := []entity.Notification{}

	tx := r.MustBegin()

	query := `SELECT * FROM notifications WHERE is_send = $1`

	err := tx.Select(&notification, query, false)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return notification, nil
}

func (r *notificationRepository) UpdateIsSendNotification(id int) error {
	tx := r.MustBegin()

	_, err := r.GetNotificationById(id)
	if err != nil {
		return err
	}

	query := `UPDATE notifications SET is_send = $1, updated_at = $2 WHERE id = $3`

	_, err = tx.Exec(query, true, "now()", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *notificationRepository) GetNotificationById(id int) (*entity.Notification, error) {
	notification := entity.Notification{}

	tx := r.MustBegin()

	query := `SELECT * FROM notifications WHERE id = $1`

	err := tx.Get(&notification, query, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Println("Notification not found")
			return nil, fmt.Errorf("Notification with id %v not found", id)
		}
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &notification, nil
}
