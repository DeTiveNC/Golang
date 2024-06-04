package main

import (
	"context"
	common "github.com/DeTiveNC/commons"
	pb "github.com/DeTiveNC/commons/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:2000"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(ctx, orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to dial to order service: ", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Failed to close connection: ", err)
		}
	}(conn)

	log.Println("Connected to order service at ", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	h := NewHandler(c)
	h.RegisterRoutes(mux)

	log.Printf("Starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start ", err)
	}
}
