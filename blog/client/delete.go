package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")
	// create a request
	req := &pb.BlogId{Id: id}
	// call gRPC endpoint to delete a blog
	_, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}
	log.Println("Blog was deleted")
}
