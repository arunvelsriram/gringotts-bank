package recommendation

import (
	"context"
	"encoding/json"
	"gringotts-bank/pkg/http"
	"gringotts-bank/pkg/log"
	"gringotts-bank/pkg/redis"
	"gringotts-bank/service/customer"
	"gringotts-bank/service/payment"
	"strings"
	"time"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	redispkg "github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type OfferVariant int

const (
	SeniorCitizen OfferVariant = iota
	UPI
	Loans
	PremiumCC
	SafeInvestor
	RiskInvestor
)

func (o OfferVariant) String() string {
	return []string{"seniorcitizen", "upi", "loans", "premiumcc", "safeinvestor", "riskinvestor"}[o]
}

type OfferVariants []OfferVariant

func (offerVariants OfferVariants) String() string {
	s := []string{}

	for _, offerVariant := range offerVariants {
		s = append(s, offerVariant.String())
	}

	return strings.Join(s, ",")
}

type Server struct {
	serviceName    string
	listenAddr     string
	customerClient customer.Client
	paymentClient  payment.Client
	rDb            *redispkg.Client
}

func (s Server) Run() error {
	app := fiber.New(fiber.Config{AppName: s.serviceName})

	app.Use(otelfiber.Middleware())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	app.Get("/customers/:id/recommendations", func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		logger := log.Logger(ctx)

		id := c.Params("id")
		if id == "" {
			logger.Error("invalid customer id", zap.String("customer_id", id))
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customer id is invalid"})
		}

		var customer Customer
		logger.Info("fetching customer", zap.String("customer_id", id))
		err := s.customerClient.GetCustomer(ctx, id, &customer)
		if err != nil {
			logger.Error("failed to fetch customer", zap.Error(err), zap.String("customer_id", id))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get recommendations"})
		}

		var transactions Transactions
		logger.Info("fetching customer transactions", zap.String("customer_id", id))
		err = s.paymentClient.GetCustomerTransactions(ctx, id, &transactions)
		if err != nil {
			logger.Error("failed to fetch customer transactions", zap.String("customer_id", id))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get recommendations"})
		}

		offerVariants := s.computeOffers(ctx, customer, transactions)
		recommendations, err := s.getRecommendations(ctx, offerVariants)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get recommendations"})
		}

		logger.Info("computed recommendations", zap.Int("recommendation_count", len(recommendations)))

		return c.Status(fiber.StatusOK).JSON(recommendations)
	})

	return app.Listen(s.listenAddr)
}

func (s Server) computeOffers(ctx context.Context, customer Customer, transactions Transactions) OfferVariants {
	tp := otel.GetTracerProvider()
	tracer := tp.Tracer("gringotts-bank-manual")
	ctx, span := tracer.Start(ctx, "compute-offers")
	defer span.End()

	logger := log.Logger(ctx)

	var offerVariants OfferVariants
	if customer.Age > 65 {
		offerVariants = append(offerVariants, SeniorCitizen)
		logger.Info("senior citien customer", zap.String("offer_variants", offerVariants.String()))
	}

	if (customer.Age >= 24 && customer.Age <= 30) && (transactions.MonthlyTransactionAmount() > 500000) {
		offerVariants = append(offerVariants, RiskInvestor, PremiumCC)
		logger.Info("high risk customer with more spends", zap.String("offer_variants", offerVariants.String()))
	}

	if (customer.Age >= 30 && customer.Age <= 58) && (transactions.MonthlyTransactionAmount() > 500000) {
		offerVariants = append(offerVariants, SafeInvestor, PremiumCC)
		logger.Info("low risk customer with more spends", zap.String("offer_variants", offerVariants.String()))
	}

	if transactions.MonthlyUpiTransactionCount() > 10 {
		offerVariants = append(offerVariants, UPI)
		logger.Info("customer with high upi transactions", zap.String("offer_variants", offerVariants.String()))
	}

	// Intentional Delay
	if customer.Name == "Hagrid" {
		time.Sleep(5 * time.Second)
	}

	return offerVariants
}

func (s Server) getRecommendations(ctx context.Context, offerVariants OfferVariants) (Recommendations, error) {
	logger := log.Logger(ctx)

	var recommendations Recommendations
	for _, offerVariant := range offerVariants {
		offerVariantStr := offerVariant.String()

		cmd := s.rDb.LRange(ctx, offerVariantStr, 0, -1)
		offers, err := cmd.Result()
		if err != nil {
			logger.Error("unable to get offer metadata from redis",
				zap.String("offer_variant", offerVariantStr), zap.Error(err))
			return nil, err
		}

		logger.Info("fetched offers for offer variant from redis",
			zap.String("offer_variant", offerVariantStr), zap.Int("offer_count", len(offers)))

		for _, offer := range offers {
			var recommendation Recommendation
			err := json.Unmarshal([]byte(offer), &recommendation)
			if err != nil {
				logger.Error("failed to parse offer json",
					zap.String("offer_variant", offerVariantStr), zap.Error(err), zap.String("offer", offer))
				return nil, err
			}
			recommendations = append(recommendations, recommendation)
		}
	}

	return recommendations, nil
}

func NewServer(ctx context.Context, serviceName, listenAddr, redisAddr string) (*Server, error) {
	httpClient := http.NewClient()

	rDb, err := redis.NewClient(redisAddr)
	if err != nil {
		return nil, err
	}

	return &Server{
		serviceName:    serviceName,
		listenAddr:     listenAddr,
		customerClient: customer.NewClient(httpClient),
		paymentClient:  payment.NewClient(httpClient),
		rDb:            rDb,
	}, nil
}
