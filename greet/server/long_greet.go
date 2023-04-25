package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	// init a string to store the result
	res := ""

	// infinie loop: keep receiving client stream until EOF
	for {
		req, err := stream.Recv()

		// check if client stream has ended
		if err == io.EOF {
			// if yes, send back the result
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		// in each req in stream, concat the req.FirstName
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
