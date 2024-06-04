package main

import (
	"context"
	pb "github.com/DeTiveNC/commons/api"
	"google.golang.org/grpc"
	"log"
)

type GrpcHandler struct {
	pb.UnimplementedOrderServiceServer

	service OrderService
}

func NewGrpcHandler(grpc *grpc.Server, service OrderService) {
	handler := &GrpcHandler{
		service: service,
	}
	pb.RegisterOrderServiceServer(grpc, handler)
}

func (h *GrpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New Order received! Order %v", p)
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
