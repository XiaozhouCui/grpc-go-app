package main

import (
	"log"
	"net"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

// gRPC server object
type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	// grpc server needs an instance for the GreetService, need &Server{} to implement rpc endpoints
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
