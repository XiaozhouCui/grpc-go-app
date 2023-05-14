package main

import (
	"context"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, req *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v\n", req)

	// parse the ID string to a MongoDB ObjectID
	oid, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}
	// create the data struct to update MongoDB
	data := &BlogItem{
		AuthorID: req.AuthorId,
		Title:    req.Title,
		Content:  req.Content,
	}
	// update the blog in MongoDB
	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)
	// check if the update was successful
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}
	// check if the blog was found
	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Could not find blog with specified ID",
		)
	}

	return &emptypb.Empty{}, nil
}
