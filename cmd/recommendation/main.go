package main

import (
	"context"
	"gringotts-bank/pkg/tracing"
	"log"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

const service = "recommendation"
const version = "1.0.0"

func main() {
	ctx := context.Background()

	tp := tracing.InitTracer(ctx, service, version)
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	app := fiber.New()
	app.Use(otelfiber.Middleware())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	if err := app.Listen(":8082"); err != nil {
		panic(err)
	}
}
