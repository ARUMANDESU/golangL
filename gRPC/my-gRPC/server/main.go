package main

import (
	"context"
	"github.com/arumandesu/grpcL/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, _ := net.Listen("tcp", "0.0.0.0:50051")
	s := grpc.NewServer()
	gen.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) AddProduct(ctx context.Context, in *gen.Product) (*gen.ProductID, error) {
	// бизнес-логика
}

// удаленный метод для получения товара
func (s *server) GetProduct(ctx context.Context, in *gen.ProductID) (*gen.Product, error) {
	// бизнес-логика
}
