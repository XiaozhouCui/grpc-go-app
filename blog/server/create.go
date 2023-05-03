package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, req *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked with %v\n", req)

	// convert the data to MongoDB format
	data := BlogItem{
		AuthorID: req.AuthorId,
		Title:    req.Title,
		Content:  req.Content,
	}

	// insert the data into MongoDB
	res, err := collection.InsertOne(ctx, data)

	// check if the insertion was successful
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v\n", err))
	}

	// get the inserted ID and cast into an ObjectID
	oid, ok := res.InsertedID.(primitive.ObjectID)

	// check if the cast was successful
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to OID: %v\n", err))
	}

	// return the response with the inserted ID
	return &pb.BlogId{Id: oid.Hex()}, nil
}
