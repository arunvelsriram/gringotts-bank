package recommendation

import (
	"context"
	"gringotts-bank/pkg/http"
	"gringotts-bank/pkg/log"
	"gringotts-bank/service/customer"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Server struct {
	serviceName    string
	listenAddr     string
	customerClient customer.Client
	rDb            *redis.Client
}

func (s Server) Run() error {
	app := fiber.New(fiber.Config{AppName: s.serviceName})

	app.Use(otelfiber.Middleware())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	app.Get("/recommendations", func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		logger := log.Logger(ctx)
		id := c.Query("customerId")
		var customer Customer

		err := s.customerClient.GetCustomer(ctx, id, &customer)
		if err != nil {
			logger.Error("failed to get customer", zap.Error(err), zap.String("customer_id", id))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get recommendations"})
		}

		logger.Info("feteched customer", zap.Int("customer_id", customer.ID))

		// x := s.rDb.LRange(ctx, "upi", 0, -1)
		// v, err := x.Result()
		// if err != nil {
		// 	logger.Error("failed", zap.Error(err))

		// 	return c.SendStatus(fiber.StatusInternalServerError)
		// }

		// logger.Info("successful", zap.String("val", strings.Join(v, "----")))
		return c.SendStatus(fiber.StatusOK)
	})

	return app.Listen(s.listenAddr)
}

func NewServer(ctx context.Context, serviceName, listenAddr, redisAddr string) (*Server, error) {
	rDb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	if err := redisotel.InstrumentTracing(rDb); err != nil {
		return nil, err
	}

	httpClient := http.NewClient()
	customerClient := customer.NewClient(httpClient)

	return &Server{
		serviceName:    serviceName,
		listenAddr:     listenAddr,
		customerClient: customerClient,
		rDb:            rDb,
	}, nil
}
