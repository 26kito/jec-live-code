package repository

import (
	"fmt"
	"jec-live-code/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

type NotificationRepository interface {
	Create(payload entity.InsertNotificationRequest)
	GetUnsendNotification() ([]entity.Notification, error)
	UpdateIsSendNotification(id int) error
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

	query := `UPDATE notifications SET is_send = $1 WHERE id = $2`

	_, err = tx.Exec(query, true, id)
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
