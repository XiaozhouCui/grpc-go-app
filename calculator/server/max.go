package main

import (
	"io"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func (*Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	// init a slice to store the numbers received from client
	var nums = []int64{}
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

		nums = append(nums, req.Number)

		for _, num := range nums {
			if num > max {
				max = num
			}
		}

		// send back the result as a stream
		err = stream.Send(&pb.MaxResponse{
			Result: max,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}

}
