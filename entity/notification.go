package entity

type Notification struct {
	ID      int    `db:"id"`
	Email   string `db:"email"`
	Message string `db:"message"`
	Type    string `db:"type"`
	IsSend  bool   `db:"is_send"`
}

type InsertNotificationRequest struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
