package dto

type Hub struct {
	NotificationChannel map[int64]chan NotificationData
}
