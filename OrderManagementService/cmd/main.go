package main

import (
	"OrderManagementService/controllers/rpc"
	"OrderManagementService/proto"
	"OrderManagementService/repository"
	"OrderManagementService/services"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	orderRepo := repository.NewOrderRepository()
	orderService := services.NewOrderService(orderRepo)
	orderController := rpc.NewOrderController(orderService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterOrderManagementServer(s, orderController)

	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
