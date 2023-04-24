package main

import (
	"context"
	"io"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	req := &pb.PrimesRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes RPC: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Response from Primes: %v\n", res.Result)
	}
}
