package main

import (
	"context"
	"gringotts-bank/api/customer"
	"gringotts-bank/pkg/tracing"
	"log"
)

const service = "customer"
const version = "5.0.0"
const listenAddr = ":8081"

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

	server := customer.NewServer(service, listenAddr)
	if err := server.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
