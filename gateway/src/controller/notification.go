package controller

import (
	entity "gateway/dto"
	"gateway/src/service"

	"github.com/gofiber/fiber/v2"
)

type NotificationController struct {
	notificationService service.NotificationService
}

func NewNotificationController(notificationService service.NotificationService) *NotificationController {
	return &NotificationController{notificationService}
}

func (n *NotificationController) CreateNotification(c *fiber.Ctx) error {
	var payload entity.InsertNotificationRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := n.notificationService.CreateNotification(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Notification created",
	})
}

func (n *NotificationController) GetUnsendNotification(c *fiber.Ctx) error {
	res, err := n.notificationService.GetUnsendNotification()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}

func (n *NotificationController) UpdateIsSendNotification(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := n.notificationService.UpdateIsSendNotification(uint32(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Notification updated",
	})
}
