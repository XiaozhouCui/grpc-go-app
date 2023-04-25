package main

import (
	"io"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func (*Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var max int64 = 0

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

		// send back the result only if the req.Number is greater than the current max
		if req.Number > max {
			max = req.Number
			err := stream.Send(&pb.MaxResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}

}
