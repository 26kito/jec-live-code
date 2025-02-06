package routes

import (
	"gateway/config"
	"gateway/src/controller"
	"gateway/src/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func NewRouter() *fiber.App {
	app := fiber.New()

	notificationClient, err := config.InitNotificationServiceClient()
	if err != nil {
		log.Println(err)
	}

	notificationService := service.NewNotificationService(notificationClient)
	notificationController := controller.NewNotificationController(notificationService)

	app.Post("/notifications", notificationController.CreateNotification)
	app.Get("/unsend-notifications", notificationController.GetUnsendNotification)
	app.Put("/notifications/:id", notificationController.UpdateIsSendNotification)

	return app
}
