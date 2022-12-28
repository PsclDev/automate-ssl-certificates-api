package handlers

import "github.com/gofiber/fiber/v2"

func invalidResponseType(ctx *fiber.Ctx) error {
	return ctx.Status(400).SendString("Invalid response type requested - 'json' (default) or 'file' are allowed")
}

func errorResponse(ctx *fiber.Ctx, err error) error {
	return ctx.Status(500).SendString(err.Error())
}