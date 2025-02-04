package main

import (
	"jec-live-code/config"
	"jec-live-code/domain/notification/repository"
	"jec-live-code/domain/notification/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := config.InitDatabase()
	if err != nil {
		log.Println(err)
	}

	app := fiber.New()

	notificationRepository := repository.NewNotificationRepository(db)
	notificationService := service.NewNotificationService(notificationRepository)

	app.Post("/notifications", notificationService.CreateNotification)
	app.Get("/unsend-notifications", notificationService.GetUnsendNotification)
	app.Put("/notifications/:id", notificationService.UpdateIsSendNotification)

	app.Listen(":3000")
}
