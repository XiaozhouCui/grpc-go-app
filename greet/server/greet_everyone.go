package main

import (
	"io"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

func (*Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked")

	// infinite loop: keep receiving client stream until EOF
	for {
		req, err := stream.Recv()

		// stop if client stream has ended
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"

		// send back the result as a stream
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
