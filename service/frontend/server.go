package frontend

import (
	"context"
	"gringotts-bank/pkg/http"
	"gringotts-bank/pkg/log"
	"gringotts-bank/pkg/middleware"
	"gringotts-bank/service/customer"
	"gringotts-bank/service/recommendation"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	serviceName          string
	listenAddr           string
	customerClient       customer.Client
	recommendationClient recommendation.Client
}

func (s Server) Run() error {
	app := fiber.New(fiber.Config{AppName: s.serviceName})

	app.Use(otelfiber.Middleware())
	app.Use(middleware.TraceBaggagePopulator())
	app.Use(middleware.BaggageToSpanAttributes())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	app.Get("/recommendations", func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		logger := log.Logger(ctx)

		customerId := c.Query("customerId")
		if customerId == "" {
			logger.Error("invalid customer id", zap.String("customer_id", customerId))
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid customer id"})
		}

		var customer Customer
		err := s.customerClient.GetCustomer(ctx, customerId, &customer)
		if err != nil {
			logger.Error("fetching customer failed", zap.Error(err), zap.String("customer_id", customerId))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get recommendations"})
		}

		var recommendations Recommendations
		err = s.recommendationClient.GetRecommendations(ctx, customerId, &recommendations)
		if err != nil {
			logger.Error("fetching recommendations failed", zap.Error(err), zap.String("customer_id", customerId))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get recommendations"})
		}

		res := RecommendationsResponse{
			CustomerId:      customer.ID,
			CustomerName:    customer.Name,
			CustomerAge:     customer.Age,
			Recommendations: recommendations,
		}

		return c.Status(fiber.StatusOK).JSON(res)
	})

	app.Static("/", "./service/frontend/web")

	return app.Listen(s.listenAddr)
}

func NewServer(ctx context.Context, serviceName, listenAddr string) Server {
	httpClient := http.NewClient()

	return Server{
		serviceName:          serviceName,
		listenAddr:           listenAddr,
		customerClient:       customer.NewClient(httpClient),
		recommendationClient: recommendation.NewClient(httpClient),
	}
}
