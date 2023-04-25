package main

import (
	"context"
	"log"
	"time"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg function was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 9},
		{Number: 2},
		{Number: 3},
		{Number: 7},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg RPC: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg response: %v\n", res.Result)
}
