package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	req := &pb.GreetRequest{
		FirstName: "Joe",
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v\n", err)
	}
	log.Printf("Response from Greet RPC: %v\n", res.Result)
}
