package main

import (
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func (s *Server) Primes(req *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", req)

	// divide the number n by 2 as many times as possible
	n := req.Number
	k := int64(2) // divider

	for n > 1 {
		// while n is divisible by k
		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			n = n / k
		} else {
			// increase k by 1
			k = k + 1
		}
	}

	return nil
}
