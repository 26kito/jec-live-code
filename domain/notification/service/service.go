package service

import (
	"fmt"
	"jec-live-code/domain/notification/entity"
	"jec-live-code/domain/notification/repository"
	"strings"

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

	if err := validateCreateNotificationRequest(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	s.notificationRepository.Create(payload)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Notification created",
	})
}

func (s *NotificationService) GetUnsendNotification(c *fiber.Ctx) error {
	notifications, err := s.notificationRepository.GetUnsendNotification()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    notifications,
	})
}

func (s *NotificationService) UpdateIsSendNotification(c *fiber.Ctx) error {
	// Get id from URL
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = s.notificationRepository.UpdateIsSendNotification(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Notification updated",
	})
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

	if payload.Type != "SMS" && payload.Type != "EMAIL" {
		return fmt.Errorf("invalid notification type")
	}

	return nil
}
