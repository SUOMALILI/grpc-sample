package main

import (
	"context"
	"log"
	"net"
	
	pb "github.com/yourname/grpc-demo/go-service/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDemoServiceServer
}

func (s *server) Process(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Output: "Go processed: " + req.Input}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDemoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}