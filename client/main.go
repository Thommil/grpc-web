package main

import (
	"context"
	"io"
	"log"
	"os"

	service "github.com/thommil/grpc-web/service"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := service.NewServiceClient(conn)

	// Contact the server and print out its response.
	name := "DEFAULT"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := c.Simple(ctx, &service.SimpleRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.Message)

	stream, err := c.Push(ctx, &service.SimpleRequest{Name: name})

	if err != nil {
		log.Fatalf("could not stream greet: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		log.Printf("Stream response: %s", message)
	}
}
