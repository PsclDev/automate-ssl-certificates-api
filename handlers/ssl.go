package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	return c.Status(200).SendString("Ok")
}