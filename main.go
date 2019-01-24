package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	service "github.com/thommil/grpc-web/service"
)

const (
	port = ":8080"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Simple(ctx context.Context, in *service.SimpleRequest) (*service.SimpleResponse, error) {
	log.Printf("Received: %v", in.Name)
	return &service.SimpleResponse{Greeting: "Hello : " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterGreeterServer(s, &server{})
	service.RegisterServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
