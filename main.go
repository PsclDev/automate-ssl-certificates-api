package main

import (
	"api/handlers"
	"fmt"
	"log"
	"os"

	sentryfiber "github.com/aldy505/sentry-fiber"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatal("Failed to initialize sentry")
	}

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(sentryfiber.New(sentryfiber.Options{}))

	v1 := app.Group("/api/v1")
	
	config := v1.Group("/config")
	config.Get("/root", handlers.GetMakeRoot)
	config.Get("/certificate", handlers.GetMakeCertificate)
	config.Post("/", handlers.PostConfig)

	cert := v1.Group("/cert")
	cert.Get("/", handlers.GetAllCerts)
	cert.Get("/root", handlers.GetRootCert)
	cert.Get("/:name", handlers.GetCertByName)
	cert.Post("/", handlers.CreateCert)
	cert.Patch("/", handlers.RecreateCert)
	cert.Delete("/:name", handlers.DeleteCert)

	app.Get("/health", handlers.HealthCheck)
	app.Use(handlers.NotFound)

	port := os.Getenv("PORT")
	listen := fmt.Sprintf(":%s", port)
    log.Fatal(app.Listen(listen))
}