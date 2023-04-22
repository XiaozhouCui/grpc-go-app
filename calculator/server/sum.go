package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Sum function was invoked")
	a := req.FirstNumber
	b := req.SecondNumber

	result := a + b

	res := &pb.SumResponse{
		SumResult: result,
	}

	return res, nil
}
