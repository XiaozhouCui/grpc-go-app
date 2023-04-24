package main

import (
	"context"
	"io"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Joe",
	}

	// call the GreetManyTimes RPC endpoint (server streaming)
	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v\n", err)
	}

	// loop over the stream
	for {
		// receive the stream in format of GreetResponse
		msg, err := stream.Recv()

		// if the stream is closed, break the loop
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Response from GreetManyTimes RPC: %v\n", msg.Result)
	}
}
