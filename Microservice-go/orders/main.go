package main

import (
	"context"
	common "github.com/DeTiveNC/commons"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("Failed to listen ", err)
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			log.Fatal("Failed to close ", err)
		}
	}(l)

	store := NewStore()
	service := NewService(store)
	NewGrpcHandler(grpcServer, service)

	service.CreateOrder(context.Background())

	log.Println("Starting server on ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Failed to start ", err)
	}
}
