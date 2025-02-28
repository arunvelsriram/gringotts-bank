package payment

import (
	"context"
	"gringotts-bank/pkg/postgres"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	serviceName string
	listenAddr  string
	db          *gorm.DB
}

func (s Server) Run() error {
	app := fiber.New(fiber.Config{AppName: s.serviceName})

	app.Use(otelfiber.Middleware())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	app.Get("/customers/:id/transactions", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	})

	return app.Listen(s.listenAddr)
}

func NewServer(ctx context.Context, serviceName, listenAddr, dbConnUrl string) (*Server, error) {
	db, err := postgres.NewConnection(dbConnUrl)
	if err != nil {
		return nil, err
	}

	return &Server{
		serviceName: serviceName,
		listenAddr:  listenAddr,
		db:          db,
	}, nil
}
