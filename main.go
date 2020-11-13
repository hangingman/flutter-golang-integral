package main

import (
	"log"
	"net"

	"github.com/hangingman/flutter-golang-integral/pb"
	"github.com/hangingman/flutter-golang-integral/service"
	"google.golang.org/grpc"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	srv := grpc.NewServer()
	pb.RegisterEchoServiceServer(srv, &service.EchoService{})
	log.Printf("start server on port%s\n", port)

	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
