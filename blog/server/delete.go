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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, req *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v\n", req)
	// parse the ID string to a MongoDB ObjectID
	oid, err := primitive.ObjectIDFromHex(req.Id)
	// check if the parsing went well
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}
	// delete the blog in MongoDB
	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	// check if the delete was successful
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v\n", err),
		)
	}
	// check if the blog was found
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	}
	// return an empty struct
	return &emptypb.Empty{}, nil
}
