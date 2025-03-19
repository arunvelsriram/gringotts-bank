package log

import (
	"context"
	"gringotts-bank/pkg/contextutil"
	"sync"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   sync.Once
	logger *otelzap.Logger
)

func Logger(ctx context.Context) otelzap.LoggerWithCtx {
	once.Do(func() {
		serviceName, ok := ctx.Value(contextutil.ServiceNameKey).(string)
		if !ok {
			panic("service name not present in context")
		}

		l, err := zap.NewProduction(zap.Fields(zap.String("service", serviceName)))
		if err != nil {
			panic(err)
		}

		logger = otelzap.New(l,
			otelzap.WithMinLevel(zapcore.InfoLevel),
			otelzap.WithTraceIDField(true),
		)
	})

	return logger.Ctx(ctx)
}
