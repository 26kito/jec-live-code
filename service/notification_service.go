package service

import (
	"jec-live-code/entity"
	"jec-live-code/repository"

	"github.com/gofiber/fiber/v2"
)

type NotificationService struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationService(notificationRepository repository.NotificationRepository) *NotificationService {
	return &NotificationService{notificationRepository}
}

func (s *NotificationService) CreateNotification(c *fiber.Ctx) error {
	var payload entity.InsertNotificationRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	s.notificationRepository.Create(payload)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Notification created",
	})
}
