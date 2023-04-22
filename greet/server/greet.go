package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

// implement Greet RPC endpoint

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", req)
	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
