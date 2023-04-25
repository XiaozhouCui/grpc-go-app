package main

import (
	"context"
	"log"
	"time"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	// init a slice of requests
	reqs := []*pb.GreetRequest{
		{FirstName: "Joe"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	// init a stream by calling the LongGreet RPC
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	// iterate over the slice of requests and send each req to the stream
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		// sleep for 1 second
		time.Sleep(1 * time.Second)
	}

	// close the stream and receive the response
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet %v\n", err)
	}

	log.Printf("LongGreet Response: %v\n", res.Result)
}
