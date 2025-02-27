package log

import (
	"context"
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
		l, err := zap.NewProduction()
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
