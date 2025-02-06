package service

import (
	"context"
	"fmt"
	"notification-service/domain/notification/entity"
	"notification-service/domain/notification/repository"
	"strings"

	pb "notification-service/proto"
)

type NotificationService struct {
	notificationRepository repository.NotificationRepository
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(notificationRepository repository.NotificationRepository) *NotificationService {
	return &NotificationService{notificationRepository: notificationRepository}
}

func (s *NotificationService) CreateNotification(ctx context.Context, req *pb.InsertNotificationRequest) (*pb.InsertNotificationResponse, error) {
	payload := entity.InsertNotificationRequest{
		Email:   req.Email,
		Message: req.Message,
		Type:    req.Type,
	}

	if err := validateCreateNotificationRequest(payload); err != nil {
		return nil, err
	}

	res, err := s.notificationRepository.CreateNotification(payload)
	if err != nil {
		return nil, err
	}

	return &pb.InsertNotificationResponse{Id: uint32(res.ID)}, nil
}

func (s *NotificationService) GetUnsendNotification(ctx context.Context, req *pb.Empty) (*pb.GetUnsendNotificationResponse, error) {
	notifications, err := s.notificationRepository.GetUnsendNotification()
	if err != nil {
		return nil, err
	}

	var res []*pb.Notification
	for _, notification := range notifications {
		res = append(res, &pb.Notification{
			Id:      uint32(notification.ID),
			Email:   notification.Email,
			Message: notification.Message,
			Type:    notification.Type,
			IsSend:  notification.IsSend,
		})
	}
	fmt.Println(notifications)
	fmt.Println(res)

	return &pb.GetUnsendNotificationResponse{Notifications: res}, nil
}

func (s *NotificationService) UpdateIsSendNotification(ctx context.Context, req *pb.UpdateIsSendNotificationRequest) (*pb.Empty, error) {
	err := s.notificationRepository.UpdateIsSendNotification(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validateCreateNotificationRequest(payload entity.InsertNotificationRequest) error {
	if payload.Email == "" {
		return fmt.Errorf("email is required")
	}

	if len(payload.Email) > 30 || len(payload.Email) < 10 {
		return fmt.Errorf("invalid email")
	}

	if !strings.Contains(payload.Email, "@") == true {
		return fmt.Errorf("invalid email")
	}

	if payload.Message == "" {
		return fmt.Errorf("message is required")
	}

	if payload.Type == "" {
		return fmt.Errorf("notification type is required")
	}

	if payload.Type != "sms" && payload.Type != "email" {
		return fmt.Errorf("invalid notification type")
	}

	return nil
}
