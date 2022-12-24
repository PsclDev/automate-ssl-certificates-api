package handlers

import (
	"api/models"
	"api/services"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetMakeRoot(ctx *fiber.Ctx) error {
	path, err := services.MakePath(services.MakeRootFileName)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.Status(200).SendFile(path)
}

func GetMakeCertificate(ctx *fiber.Ctx) error {
	path, err := services.MakePath(services.MakeCertFileName)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.Status(200).SendFile(path)
}

func PostConfig(ctx *fiber.Ctx) error {
	domainName := new(models.DomainName)

	if err := ctx.BodyParser(domainName); err != nil {
		return err
	}
	v := validator.New()
	err := v.Struct(domainName)
	if err != nil {
		errors := ""
		for _, e := range err.(validator.ValidationErrors) {
			errors += fmt.Sprintf("%s\n", e)
		}
		return ctx.Status(400).SendString(errors)
	}
	
	services.SetConfig(domainName)

	return ctx.Status(200).SendString("Config was set")
}