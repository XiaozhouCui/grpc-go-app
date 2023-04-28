package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("doGreetWithDeadline function was invoked!")
	req := &pb.GreetRequest{
		FirstName: "Joe",
	}
	// create a timeout context, if the server doesn't respond within ? seconds, it will timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// cancel the context to avoid memory leak
	defer cancel()

	// Call GreetWithDeadline RPC
	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		// Check if the error is a gRPC error
		respErr, ok := status.FromError(err)
		if ok {
			if respErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded!")
				// cancel the request if the timeout was hit
				return
			} else {
				fmt.Printf("Unexpected grpc error: %v\n", respErr)
			}
		} else {
			log.Fatalf("Non-grpc error while calling GreetWithDeadline RPC: %v", err)
		}
	}
	log.Printf("Response from GreetWithDeadline: %v\n", res.Result)
}
