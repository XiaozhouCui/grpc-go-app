package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, req *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v\n", req)

	// convert string ID (from proto) to MongoDB's ObjectID
	oid, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}
	// create an empty struct
	data := &BlogItem{}
	// filter to get the blog with the provided ID
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)
	// decode the result into data
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v\n", err),
		)
	}
	// convert data into Blog message
	return documentToBlog(data), nil
}
