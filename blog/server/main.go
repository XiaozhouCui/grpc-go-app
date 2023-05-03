package main

import (
	"context"
	"log"
	"net"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

// declare a MongoDB collection
var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	// create a mongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017"))

	if err != nil {
		log.Fatalf("Failed to create mongo client: %v\n", err)
	}

	// create a mongoDB connection
	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("Failed to connect to mongo server: %v\n", err)
	}

	// get a handle for your collection
	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
