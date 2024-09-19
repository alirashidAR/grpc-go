// server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-go/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	return &proto.HelloResponse{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
