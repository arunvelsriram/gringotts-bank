package payment

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

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	app.Get("/customers/:id/transactions", func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		logger := log.Logger(ctx)
		customerId := c.Params("id")
		var transactions Transactions

		logger.Info("getting customer transactions from db", zap.String("customer_id", customerId))

		result := s.db.WithContext(ctx).
			Where("customer_id = ?", customerId).
			Find(&transactions)
		if result.Error != nil {
			logger.Error("db query failed", zap.Error(result.Error))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get transactions of customer"})
		}

		logger.Info("fetched customer transactions from db", zap.Int("transactionsCount", len(transactions)))

		return c.Status(fiber.StatusOK).JSON(transactions)
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
