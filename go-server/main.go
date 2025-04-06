package main

import (
	"context"
	"log"
	
	pb_go "github.com/yourname/grpc-demo/go-service/proto"
	pb_py "github.com/yourname/grpc-demo/go-server/proto"
	"google.golang.org/grpc"
)

func callService(addr string, client pb_go.DemoServiceClient, input string) {
	res, err := client.Process(context.Background(), &pb_go.Request{Input: input})
	if err != nil {
		log.Printf("Error calling %s: %v", addr, err)
		return
	}
	log.Printf("%s response: %s", addr, res.Output)
}

func main() {
	// 连接Python服务
	pyConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer pyConn.Close()
	pyClient := pb_py.NewDemoServiceClient(pyConn)

	// 连接Go服务
	goConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer goConn.Close()
	goClient := pb_go.NewDemoServiceClient(goConn)

	// 调用两个服务
	callService("Python", pyClient, "test request")
	callService("Go", goClient, "test request")
}