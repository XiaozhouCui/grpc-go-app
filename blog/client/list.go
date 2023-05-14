package main

import (
	"context"
	"io"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("---listBlogs was invoked---")

	// call gRPC endpoint to list all blogs
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs RPC: %v\n", err)
	}
	// infinite loop: iterate over the stream to get the data
	for {
		res, err := stream.Recv()
		// if we reach the end of the stream, we break the loop
		if err == io.EOF {
			break
		}
		// if we get an error while reading the stream, we log it
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Println(res)
	}
}
