package handlers

import (
	"os"

	sentryfiber "github.com/aldy505/sentry-fiber"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kpango/glg"
)

func invalidResponseType(ctx *fiber.Ctx, resType string) error {
	glg.Warnf("Invalid response type request '%s'", resType)
	return ctx.Status(406).SendString("Invalid response type requested - 'json' (default) or 'file' are allowed")
}

func invalidBody(ctx *fiber.Ctx) error {
	return ctx.Status(422).SendString("Body couldn't be parsed, check the documentations for the correct request body")
}

func failedValidation(ctx *fiber.Ctx, err error) error {
	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = "is missing, but required"
	}

	glg.Warnf("validation failed, reason(s) '%s'", err.Error())
	return ctx.Status(400).JSON(errors)
}

func notFound(ctx *fiber.Ctx) error {
	return ctx.Status(404).SendString("Could not be found")
}

func alreadyExists(ctx *fiber.Ctx) error {
	return ctx.Status(409).SendString("Cert already exists, if you want to recreate use PATCH request")
}

func errorResponse(ctx *fiber.Ctx, err error) error {
	if os.IsNotExist(err) {
		return notFound(ctx)
	}
	glg.Warnf(err.Error())
	if err.Error() == "cert already exists and was not forced to recreate" {
		return alreadyExists(ctx)
	}

	glg.Errorf("An error has occurred: '%s'", err)
	if hub := sentryfiber.GetHubFromContext(ctx); hub != nil {
		hub.CaptureException(err)
	}
	return ctx.Status(500).SendString("An error has occurred, view the logs to get more informations")
}