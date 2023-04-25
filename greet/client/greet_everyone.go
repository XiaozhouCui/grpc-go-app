package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	// init a stream by calling the LongGreet RPC
	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream %v\n", err)
	}

	// init a slice of requests
	reqs := []*pb.GreetRequest{
		{FirstName: "Joe"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	// Create a go-channel with 2 go-routines running simultaneously
	waitc := make(chan struct{})

	// first go-routine to send requests stream
	go func() {
		// function to send a bunch of messages
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		// close the stream
		stream.CloseSend()
	}() // need to invoke the function

	// second go-routine to receive the response stream from server
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		// close the channel when break from the loop
		close(waitc)
	}()

	// wait channel to wait for the go-routines to finish
	<-waitc
}
