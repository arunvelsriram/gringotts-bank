package main

import (
	"context"
	"gringotts-bank/pkg/log"
	"gringotts-bank/pkg/tracing"

	"gringotts-bank/service/recommendation"

	"go.uber.org/zap"
)

const service = "recommendation"
const version = "1.0.0"
const listenAddr = ":8081"
const redisAddr = "localhost:16379"

func main() {
	ctx := context.Background()
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

	server, err := recommendation.NewServer(ctx, service, listenAddr, redisAddr)
	if err != nil {
		logger.Fatal("failed to create server", zap.Error(err))
	}

	if err := server.Run(); err != nil {
		logger.Fatal("server failed to start", zap.Error(err))
	}
}
