package main

import (
	"context"
	"github.com/janniksinz/gRPC-go/grpc-server/currency"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myCurrencyServer struct {
	currency.UnimplementedCurrencyServer
}

func (s myCurrencyServer) GetRate(ctx context.Context, req *currency.RateRequest) (*currency.RateResponse, error) {
	log.Printf("Received: %s", req)
	return nil, nil
}

func main() {
	// Start gRPC server
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myCurrencyServer{}
	currency.RegisterCurrencyServer(serverRegistrar, service) // takes ServiceRegistrar and CurrencyServer

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
