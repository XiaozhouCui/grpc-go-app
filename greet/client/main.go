package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	// close the connection when the function exits
	defer conn.Close()

	// create a greet service client so that we can call Greet RPC endpoint
	c := pb.NewGreetServiceClient(conn)

	// call the Greet RPC endpoint
	doGreet(c)
}
