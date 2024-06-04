package main

import (
	"context"
	pb "github.com/DeTiveNC/commons/api"
)

type OrderService interface {
	CreateOrder(context.Context) error
	ValidateOder(context.Context, *pb.CreateOrderRequest) error
}

type OrderStore interface {
	Create(context.Context) error
}
