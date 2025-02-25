package main

import (
	"context"
	"gringotts-bank/pkg/tracing"
	"gringotts-bank/service/customer"
	"log"
)

const service = "customer"
const version = "5.0.0"
const listenAddr = ":8081"
const dbConnUrl = "postgresql://postgres:postgres@localhost:15432/postgres?sslmode=disable"

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

	server, err := customer.NewServer(ctx, service, listenAddr, dbConnUrl)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
