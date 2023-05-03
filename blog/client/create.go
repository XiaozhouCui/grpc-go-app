package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Joe",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	// call gRPC endpoint to create a blog
	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
