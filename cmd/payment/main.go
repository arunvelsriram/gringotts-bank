package main

import (
	"context"
	"gringotts-bank/pkg/contextutil"
	"gringotts-bank/pkg/log"
	"gringotts-bank/pkg/tracing"
	"gringotts-bank/service/payment"

	"go.uber.org/zap"
)

const service = "payment"
const version = "10.0.0"
const listenAddr = ":8083"
const dbConnUrl = "postgresql://postgres:postgres@localhost:25432/postgres?sslmode=disable"

func main() {
	ctx := context.WithValue(context.Background(), contextutil.ServiceNameKey, service)
	logger := log.Logger(ctx)

	shutDownFunc, err := tracing.Init(ctx, service, version)
	if err != nil {
		logger.Info("failed to initialize tracer", zap.Error(err))
	}
	defer func() {
		if err := shutDownFunc(ctx); err != nil {
			logger.Fatal("failed to shutdown tracer", zap.Error(err))
		}
	}()

	server, err := payment.NewServer(ctx, service, listenAddr, dbConnUrl)
	if err != nil {
		logger.Fatal("failed to create server", zap.Error(err))
	}

	if err := server.Run(); err != nil {
		logger.Fatal("server failed to start", zap.Error(err))
	}
}
