package main

import (
	"context"
	"gringotts-bank/pkg/log"
	"gringotts-bank/pkg/tracing"
	"gringotts-bank/service/frontend"

	"go.uber.org/zap"
)

const service = "frontend"
const version = "1.0.0"
const listenAddr = ":8080"

func main() {
	ctx := context.Background()
	logger := log.Logger(ctx)

	shutDownFunc, err := tracing.Init(ctx, service, version)
	if err != nil {
		logger.Fatal("failed to initialize tracer", zap.Error(err))
	}
	defer func() {
		if err := shutDownFunc(ctx); err != nil {
			logger.Fatal("failed to shutdown tracer", zap.Error(err))
		}
	}()

	server := frontend.NewServer(ctx, service, listenAddr)
	if err := server.Run(); err != nil {
		logger.Fatal("server failed to start", zap.Error(err))
	}
}
