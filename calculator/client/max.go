package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax function was invoked")

	// create a stream by invoking the client
	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 2},
		{Number: 7},
		{Number: 3},
		{Number: 9},
	}

	// create a channel to receive a signal to block
	waitc := make(chan struct{})

	// first go-routine to send requests stream
	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// second go-routine to receive the response stream from server
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving response: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		// close the channel
		close(waitc)
	}()

	// block until close(waitc) is called
	<-waitc
}
