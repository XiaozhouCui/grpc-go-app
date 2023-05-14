package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")
	// create a request
	req := &pb.BlogId{Id: id}
	// call gRPC endpoint to read a blog
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
