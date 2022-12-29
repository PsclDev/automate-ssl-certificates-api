package handlers

import (
	sentryfiber "github.com/aldy505/sentry-fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/kpango/glg"
)

func invalidResponseType(ctx *fiber.Ctx, resType string) error {
	glg.Warnf("Invalid response type request '%s'", resType)
	return ctx.Status(400).SendString("Invalid response type requested - 'json' (default) or 'file' are allowed")
}

func errorResponse(ctx *fiber.Ctx, err error) error {
	glg.Errorf("An error has occurred: '%s'", err)
	if hub := sentryfiber.GetHubFromContext(ctx); hub != nil {
		hub.CaptureException(err)
	}
	return ctx.Status(500).SendString("An error has occurred, view the logs to get more informations")
}