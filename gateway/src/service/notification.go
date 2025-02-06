package service

import (
	"context"
	entity "gateway/dto"
	pb "gateway/proto"
)

type NotificationService interface {
	CreateNotification(payload entity.InsertNotificationRequest) error
	GetUnsendNotification() (*pb.GetUnsendNotificationResponse, error)
	UpdateIsSendNotification(id uint32) error
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

func (n *notificationService) GetUnsendNotification() (*pb.GetUnsendNotificationResponse, error) {
	res, err := n.client.GetUnsendNotification(context.Background(), &pb.Empty{})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (n *notificationService) UpdateIsSendNotification(id uint32) error {
	_, err := n.client.UpdateIsSendNotification(context.Background(), &pb.UpdateIsSendNotificationRequest{
		Id: id,
	})

	if err != nil {
		return err
	}

	return nil
}
