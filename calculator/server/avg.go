package main

import (
	"io"
	"log"

	pb "github.com/XiaozhouCui/grpc-go-app/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	var count int64 = 0
	var sum int64 = 0
	var res float64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// need to return if EOF
			return stream.SendAndClose(&pb.AvgResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Received req: %v\n", req)

		sum += req.Number
		count++
		res = float64(sum) / float64(count)
	}
}
