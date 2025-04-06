package main

import (
	"context"
	"log"
	"net"

	pb "grpc_demo/proto/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDemoServiceServer
}

func (s *server) Process(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Output: "Go processed: " + req.Input}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")  // 监听所有IPv4/IPv6地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDemoServiceServer(s, &server{})
	log.Printf("Go gRPC server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}