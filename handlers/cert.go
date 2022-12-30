package handlers

import (
	"api/models"
	"api/services"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kpango/glg"
)

const defaultResponse string = "json"

func GetAllCerts(ctx *fiber.Ctx) error {
	res := ctx.Query("res", defaultResponse)
	glg.Tracef("GetAllCerts | with response type '%s'", res)

	switch res {
	case "file":
		path, err := services.CreateCompleteArchive()
		if err != nil {
			return errorResponse(ctx, err)
		}
		
		return fileResponse(ctx, "archive.zip", path)
	case "json":
		certs, err := services.GetAllCertsAsJson()
		if err != nil {
			return errorResponse(ctx, err)
		}
		return ctx.Status(200).JSON(certs)
	}

	return invalidResponseType(ctx, res)
}

func GetRootCert(ctx *fiber.Ctx) error {
	glg.Trace("GetRootCert")
	err := services.CheckRootCertificate()
	if err != nil {
		return errorResponse(ctx, err)
	}

	path, err := services.CreateCertArchive("root")
	if err != nil {
		return errorResponse(ctx, err)
	}

	return fileResponse(ctx, "root.zip", path)
}

func GetCertByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	res := ctx.Query("res", defaultResponse)
	glg.Tracef("GetCertByName | for name: '%s' with response type '%s'", name, res)

	switch res {
	case "file":
		return certArchive(ctx, name)
	case "json":
		cert, err := services.GetCertAsJson(name)
		if err != nil {
			return errorResponse(ctx, err)
		}
		return ctx.Status(200).JSON(cert)
	}

	return invalidResponseType(ctx, res)
}

func CreateCert(ctx *fiber.Ctx) error {
	return createCert(ctx, false)
}

func RecreateCert(ctx *fiber.Ctx) error {
	return createCert(ctx, true)
}

func createCert(ctx *fiber.Ctx, forceCreate bool) error {
	res := ctx.Query("res", defaultResponse)
	cert := new(models.Certificate)

	if err := ctx.BodyParser(cert); err != nil {
		return invalidBody(ctx)
	}
	glg.Tracef("createCert | with response type '%s' and cert '%s'", res, cert)

	v := validator.New()
	err := v.Struct(cert)
	if err != nil {
		return failedValidation(ctx, err)
	}

	if err := services.CreateCert(cert, forceCreate); err != nil {
		return errorResponse(ctx, err)
	}


	switch res {
	case "file":
		return certArchive(ctx, cert.Name)
	case "json":
		return ctx.Status(200).JSON(cert)
	}

	return invalidResponseType(ctx, res)
}


func DeleteCert(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if err := services.DeleteCert(name); err != nil {
		return errorResponse(ctx, err)
	}

	return ctx.Status(200).SendString("Deleted")
}

func certArchive(ctx *fiber.Ctx, certName string) error {
	zipPath, err := services.CreateCertArchive(certName)
	if err != nil {
		return errorResponse(ctx, err)
	}

	return fileResponse(ctx, fmt.Sprintf("%s.zip", certName), zipPath)
}

func fileResponse(ctx *fiber.Ctx, filename, path string) error {
	disposition := fmt.Sprintf("attachment; filename=\"%s\"", filename)
	glg.Tracef("File response | with content-disposition: '%s'", disposition)
	ctx.Set("Content-Disposition", disposition)
	return ctx.Status(200).SendFile(path)
}