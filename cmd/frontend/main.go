package main

import (
	"context"
	"gringotts-bank/api/frontend"
	"gringotts-bank/pkg/tracing"
	"log"
)

const service = "frontend"
const version = "1.0.0"
const listenAddr = ":8080"

func main() {
	ctx := context.Background()

	shutDownFunc, err := tracing.Init(ctx, service, version)
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer func() {
		if err := shutDownFunc(ctx); err != nil {
			log.Fatalf("failed to shutdown tracer: %v", err)
		}
	}()

	server := frontend.NewServer(listenAddr)
	if err := server.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
