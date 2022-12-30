package main

import (
	"api/handlers"
	"api/util"
	"fmt"
	"os"

	sentryfiber "github.com/aldy505/sentry-fiber"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/kpango/glg"
)

func main() {
	godotenv.Load()

	sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})

	if os.Getenv("NETLOG_URL") != "" {
		glg.Get().
		SetMode(glg.BOTH).
		AddLevelWriter(glg.TRACE, util.NetworkLogger{}).
		AddLevelWriter(glg.PRINT, util.NetworkLogger{}).
		AddLevelWriter(glg.LOG, util.NetworkLogger{}).
		AddLevelWriter(glg.INFO, util.NetworkLogger{}).
		AddLevelWriter(glg.OK, util.NetworkLogger{}).
		AddLevelWriter(glg.WARN, util.NetworkLogger{}).
		AddLevelWriter(glg.ERR, util.NetworkLogger{}).
		AddLevelWriter(glg.FAIL, util.NetworkLogger{}).
		AddLevelWriter(glg.FATAL, util.NetworkLogger{})
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
	glg.Infof("App listening on port %s", port)
	glg.Fatal(app.Listen(listen))
}