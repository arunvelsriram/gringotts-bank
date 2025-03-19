package middleware

import (
	"gringotts-bank/pkg/log"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func TraceBaggagePopulator() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		logger := log.Logger(ctx)

		logger.Info("baggage creation middleware - START")

		customerId := c.Get("X-Customer-Id", "")
		if customerId == "" {
			logger.Info("customer id not present in header")
		} else {
			customerIdBag, err := baggage.NewMember("customer.id", customerId)
			if err != nil {
				logger.Info("baggage member creation failed", zap.Error(err))
				return err
			}

			bag, err := baggage.New(customerIdBag)
			if err != nil {
				logger.Info("baggage creation failed", zap.Error(err))
				return err
			}

			ctxWithBaggage := baggage.ContextWithBaggage(ctx, bag)
			c.SetUserContext(ctxWithBaggage)

			logger.Info("baggage creation middleware - END")
		}

		return c.Next()
	}
}

func BaggageToSpanAttributes() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		logger := log.Logger(ctx)

		logger.Info("baggage to span attributes middleware - START")

		reqBaggage := baggage.FromContext(ctx)
		span := trace.SpanFromContext(ctx)
		if !span.SpanContext().IsValid() {
			logger.Warn("span is invalid")
		}

		customerId := reqBaggage.Member("customer.id").Value()
		if customerId == "" {
			logger.Info("customer id not present in baggage", zap.String("customer_id", customerId))
		} else {

			span.SetAttributes(attribute.String("customer.id", customerId))
		}

		logger.Info("baggage to span attributes middleware - END")

		return c.Next()
	}
}
