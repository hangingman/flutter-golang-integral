package main

import (
	"log"
	"net"

	"github.com/hangingman/flutter-golang-integral/pb"
	"github.com/hangingman/flutter-golang-integral/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, &service.EchoService{})
	log.Printf("start server on port%s\n", port)

	// Register reflection service on gRPC server.
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
