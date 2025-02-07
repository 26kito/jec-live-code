package main

import (
	"log"
	"net"
	"notification-service/config"
	"notification-service/domain/notification/repository"
	"notification-service/domain/notification/service"
	pb "notification-service/proto"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	server := grpc.NewServer()

	notificationRepository := repository.NewNotificationRepository(db)
	notificationService := service.NewNotificationService(notificationRepository)

	pb.RegisterNotificationServiceServer(server, notificationService)

	log.Println("Starting server on port :50051")

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
