package main

import (
	"jec-live-code/config"
	"jec-live-code/repository"
	"jec-live-code/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := config.InitDatabase()
	if err != nil {
		// panic(err)
		log.Println(err)
	}

	app := fiber.New()

	notificationRepository := repository.NewNotificationRepository(db)
	notificationService := service.NewNotificationService(notificationRepository)

	app.Post("/notifications", notificationService.CreateNotification)
	app.Get("/unsend-notifications", notificationService.GetUnsendNotification)

	app.Listen(":3000")
}
