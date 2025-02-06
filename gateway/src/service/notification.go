package service

import (
	"context"
	entity "notification-service/domain/notification/entity"
	pb "notification-service/proto"
)

type NotificationService interface {
	CreateNotification(payload entity.InsertNotificationRequest) error
}

type notificationService struct {
	client pb.NotificationServiceClient
}

func NewNotificationService(client pb.NotificationServiceClient) NotificationService {
	return &notificationService{client}
}

func (n *notificationService) CreateNotification(payload entity.InsertNotificationRequest) error {
	_, err := n.client.CreateNotification(context.Background(), &pb.InsertNotificationRequest{
		Email:   payload.Email,
		Message: payload.Message,
		Type:    payload.Type,
	})

	if err != nil {
		return err
	}

	return nil
}
