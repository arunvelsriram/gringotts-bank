package customer

import (
	"context"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

type Server struct {
	serviceName string
	listenAddr  string
	db          *gorm.DB
}

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Customers []Customer

func (s Server) Run() error {
	app := fiber.New(fiber.Config{AppName: s.serviceName})

	app.Use(otelfiber.Middleware())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	app.Get("/customers", func(c *fiber.Ctx) error {
		var customers Customers
		ctx := c.UserContext()

		result := s.db.WithContext(ctx).Find(&customers)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get customers"})
		}

		return c.Status(fiber.StatusOK).JSON(customers)
	})

	return app.Listen(s.listenAddr)
}

func NewServer(ctx context.Context, serviceName, listenAddr, dbConnUrl string) (*Server, error) {
	db, err := gorm.Open(postgres.Open(dbConnUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, err
	}

	return &Server{
		serviceName: serviceName,
		listenAddr:  listenAddr,
		db:          db,
	}, nil
}
