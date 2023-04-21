package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

// implement Greet rpc endpoints

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
