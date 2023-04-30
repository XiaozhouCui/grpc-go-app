package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/XiaozhouCui/grpc-go-app/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	// Toggle SSL/TLS
	tls := true

	// create dial option
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	// close the connection when the function exits
	defer conn.Close()

	// create a greet service client so that we can call Greet RPC endpoint
	c := pb.NewGreetServiceClient(conn)

	// call the Greet RPC endpoints
	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5*time.Second)
	// doGreetWithDeadline(c, 1*time.Second)
}
