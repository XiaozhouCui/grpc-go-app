package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt function was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		// check if the error is a gRPC error
		e, ok := status.FromError(err)

		if ok {
			log.Printf("gRPC Error message from server: %v\n", e.Message())
			log.Printf("gRPC Error code from server: %v\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Response from Sqrt: %v\n", res.Result)
}
