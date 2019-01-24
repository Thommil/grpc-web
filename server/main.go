package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	service "github.com/thommil/grpc-web/service"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) Simple(ctx context.Context, in *service.SimpleRequest) (*service.SimpleResponse, error) {
	log.Printf("Received: %v", in.Name)
	return &service.SimpleResponse{Message: "Hello : " + in.Name}, nil
}

func (s *server) Push(in *service.SimpleRequest, out service.Service_PushServer) error {
	for i := range []int{0, 1, 2, 3, 4} {
		time.Sleep(1000 * time.Millisecond)
		out.SendMsg(&service.SimpleResponse{Message: in.Name + strconv.Itoa(i)})
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Starting server ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
