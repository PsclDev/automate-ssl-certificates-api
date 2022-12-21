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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load envs")
	}

	err = sentry.Init(sentry.ClientOptions{
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
	
	ssl := v1.Group("/ssl")
	ssl.Get("/", handlers.GetAll)

	app.Get("/health", handlers.HealthCheck)
	app.Use(handlers.NotFound)

	port := os.Getenv("PORT")
	listen := fmt.Sprintf(":%s", port)
    log.Fatal(app.Listen(listen))
}