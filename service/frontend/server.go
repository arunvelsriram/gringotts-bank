package frontend

import (
	"context"
	"gringotts-bank/pkg/http"
	"gringotts-bank/service/customer"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	serviceName    string
	listenAddr     string
	customerClient customer.Client
}

func (s Server) Run() error {
	app := fiber.New(fiber.Config{AppName: s.serviceName})

	app.Use(otelfiber.Middleware())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		customers, err := s.customerClient.GetCustomers(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get customers"})
		}

		return c.Status(fiber.StatusOK).JSON(customers)
	})

	return app.Listen(s.listenAddr)
}

func NewServer(ctx context.Context, serviceName, listenAddr string) Server {
	httpClient := http.NewClient()
	customerClient := customer.NewClient(httpClient)

	return Server{
		serviceName:    serviceName,
		listenAddr:     listenAddr,
		customerClient: customerClient,
	}
}
