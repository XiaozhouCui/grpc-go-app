package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", req)

	result := req.FirstNumber + req.SecondNumber

	res := &pb.SumResponse{
		SumResult: result,
	}

	return res, nil
}
