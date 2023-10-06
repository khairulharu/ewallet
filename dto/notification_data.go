package dto

import "time"

type NotificationData struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Status    int8      `json:"status"`
	IsRead    int8      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
