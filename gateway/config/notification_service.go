package config

import (
	pb "notification-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitNotificationServiceClient() (pb.NotificationServiceClient, error) {
	notificationConn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	notificationServiceClient := pb.NewNotificationServiceClient(notificationConn)

	return notificationServiceClient, nil
}
