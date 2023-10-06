package service

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"time"

	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
	templateRepository     domain.TemplateRepository
	hub                    *dto.Hub
}

func NewNotification(notificationRepository domain.NotificationRepository, templateRepository domain.TemplateRepository, hub *dto.Hub) domain.NotificationService {
	return &notificationService{
		notificationRepository: notificationRepository,
		templateRepository:     templateRepository,
		hub:                    hub,
	}
}

func (n notificationService) FindByUser(ctx context.Context, user int64) ([]dto.NotificationData, error) {
	notifications, err := n.notificationRepository.FindByUser(ctx, user)
	if err != nil {
		return nil, err
	}

	var result []dto.NotificationData

	for _, v := range notifications {
		result = append(result, dto.NotificationData{
			ID:        v.ID,
			UserID:    v.UserID,
			Title:     v.Title,
			Body:      v.Body,
			IsRead:    v.IsRead,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
		})
	}

	if result == nil {
		result = make([]dto.NotificationData, 0)
	}

	return result, nil
}

func (n notificationService) Insert(ctx context.Context, code string, userId int64, data map[string]string) error {
	tmpl, err := n.templateRepository.FindByCode(ctx, code)
	if err != nil {
		return err
	}

	if tmpl == (domain.Template{}) {
		return errors.New("template not found")
	}

	body := new(bytes.Buffer)

	t := template.Must(template.New("notif").Parse(tmpl.Body))

	err = t.Execute(body, data)
	if err != nil {
		return err
	}

	notification := domain.Notification{
		UserID:    userId,
		Title:     tmpl.Title,
		Body:      body.String(),
		Status:    1,
		IsRead:    0,
		CreatedAt: time.Now(),
	}

	err = n.notificationRepository.Insert(ctx, &notification)
	if err != nil {
		return err
	}

	if channel, ok := n.hub.NotificationChannel[userId]; ok {
		channel <- dto.NotificationData{
			ID:        notification.ID,
			Title:     notification.Title,
			Body:      notification.Body,
			Status:    notification.Status,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		}
	}

	return nil
}
