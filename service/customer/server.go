package customer

import (
	"context"
	"gringotts-bank/pkg/log"
	"gringotts-bank/pkg/middleware"
	"gringotts-bank/pkg/postgres"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
	app.Use(middleware.BaggageToSpanAttributes())
	// app.Use(middleware.DumpHeaders())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	app.Get("/customers", func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		logger := log.Logger(ctx)
		var customers Customers

		result := s.db.WithContext(ctx).Find(&customers)
		if result.Error != nil {
			logger.Error("db query failed", zap.Error(result.Error))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get customers"})
		}

		logger.Info("fetched all customers from db", zap.Int("customers", len(customers)))

		return c.Status(fiber.StatusOK).JSON(customers)
	})

	app.Get("/customers/:id", func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		logger := log.Logger(ctx)
		id := c.Params("id")
		var customer Customer

		result := s.db.WithContext(ctx).First(&customer, id)
		if result.Error != nil {
			logger.Error("failed to get customer from db", zap.Error(result.Error), zap.String("customer_id", id))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get customers"})
		}

		logger.Info("fetched customer from db", zap.Int("customer_id", customer.ID))

		return c.Status(fiber.StatusOK).JSON(customer)
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
