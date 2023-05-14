package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(_ *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlog was invoked")

	ctx := context.Background()
	// collection.Find returns a cursor for our (empty) query
	cur, err := collection.Find(ctx, primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}
	// make sure to close the cursor at the end
	defer cur.Close(ctx)

	// iterate over the cursor and send the blogs over stream
	for cur.Next(ctx) {
		// create an empty struct
		data := &BlogItem{}
		// decode the data
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v\n", err),
			)
		}
		// convert the blog item into a protobuf blog message
		stream.Send(documentToBlog(data))
	}

	// check if the cursor encountered any errors while iterating
	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}

	return nil
}
