package handlers

import (
	"api/models"
	"api/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kpango/glg"
)

func GetMakeRoot(ctx *fiber.Ctx) error {
	glg.Trace("GetMakeRoot")
	path, err := services.MakePath(services.MakeRootFileName)
	if err != nil {
		return errorResponse(ctx, err)
	}

	return ctx.Status(200).SendFile(path)
}

func GetMakeCertificate(ctx *fiber.Ctx) error {
	glg.Trace("GetMakeCertificate")
	path, err := services.MakePath(services.MakeCertFileName)
	if err != nil {
		return errorResponse(ctx, err)
	}

	return ctx.Status(200).SendFile(path)
}

func PostConfig(ctx *fiber.Ctx) error {
	domainName := new(models.DomainName)

	if err := ctx.BodyParser(domainName); err != nil {
		return invalidBody(ctx)
	}
	glg.Tracef("PostConfig | with domain name '%s'", domainName)

	v := validator.New()
	err := v.Struct(domainName)
	if err != nil {
		return failedValidation(ctx, err)
	}
	
	if err := services.SetConfig(domainName); err != nil {
		return errorResponse(ctx, err)
	}

	return ctx.Status(200).SendString("Config was set")
}