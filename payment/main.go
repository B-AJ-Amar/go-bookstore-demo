package main

import (
	"log"
	"net"

	config "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/config"
	paymentpb "github.com/B-AJ-Amar/go-bookstore-demo/proto/paymentpb"

	"google.golang.org/grpc"
)

func main() {
	
	config.InitDB()

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    paymentpb.RegisterPaymentServiceServer(grpcServer, &PaymentServer{})

    log.Println("Payment gRPC server running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
