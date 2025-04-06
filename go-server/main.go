package main

import (
	"context"
	"log"
	
	// 统一导入生成的 Go 客户端代码
	pb "grpc_demo/proto/gen" 
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func callService(addr string, client pb.DemoServiceClient, input string) {
	// 打印请求信息
	log.Printf("正在向 %s 发送请求，内容: %q", addr, input)
	
	// 调用服务
	res, err := client.Process(context.Background(), &pb.Request{Input: input})
	
	// 处理响应
	if err != nil {
		log.Printf("调用 %s 失败: %v", addr, err)
		return
	}
	log.Printf("收到 %s 的响应: %q", addr, res.Output)
	log.Printf("========== %s 调用完成 ==========", addr)
}

func main() {
	log.Println("启动 gRPC 客户端...")
	
	// 连接 Python 服务
	log.Println("正在连接 Python 服务 (localhost:50051)...")
	pyConn, err := grpc.Dial(
		"localhost:50051", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接 Python 服务失败: %v", err)
	}
	defer pyConn.Close()
	pyClient := pb.NewDemoServiceClient(pyConn)
	log.Println("Python 服务连接成功")

	// 连接 Go 服务
	log.Println("正在连接 Go 服务 (localhost:50052)...")
	goConn, err := grpc.Dial(
		"localhost:50052", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接 Go 服务失败: %v", err)
	}
	defer goConn.Close()
	goClient := pb.NewDemoServiceClient(goConn)
	log.Println("Go 服务连接成功")

	// 调用服务
	log.Println("\n开始调用服务...")
	callService("Python 服务", pyClient, "测试请求")
	callService("Go 服务", goClient, "测试请求")
	
	log.Println("所有服务调用完成")
}