# gRPC 跨语言调用示例

本示例演示 Python 和 Golang 服务端如何通过 gRPC 提供相同接口，以及客户端如何调用不同语言的服务端。

## 文件树结构
```
grpc-demo/
├── proto/                # Proto文件目录
│   └── common.proto
├── go-server/            # Golang主服务器
│   ├── main.go
│   ├── go.mod
│   └── proto/
├── py-service/           # Python接口服务
│   ├── service.py
│   └── proto/
└── go-service/           # Golang接口服务
    ├── main.go
    ├── go.mod
    └── proto/
```

# 依赖安装
```
# Python服务
pip install grpcio-tools

# Go服务和主服务器

# go有两种依赖配置方式，一种是依赖vendor的打包配置，一种是在本地环境配置
# 打包配置：初始化模块并下载依赖 (如果尚未执行)
go mod init
go mod tidy
go mod vendor

# 本地环境：安装protoc插件 (如果需要在全局使用)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc/credentials/insecure
```

# 运行命令生成代码

```
# 在./proto/gen下生成go的代码依赖
protoc -I=proto --go_out=proto/gen --go_opt=paths=source_relative --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative proto/common.proto

# 在./proto/python_gen下生成python的代码依赖
python -m grpc_tools.protoc -I=proto --python_out=py-service --grpc_python_out=py-service proto/common.proto
```


# 启动顺序

1. 先把服务拉起来
Python接口服务：python py-service/service.py
Go接口服务：go run go-service/main.go

2. 再让主服务器调用服务
主服务器：go run go-server/main.go
最终效果
主服务器会同时调用两个服务并输出：
```
开始调用服务...
2025/04/07 00:25:21 正在向 Python 服务 发送请求，内容: "测试请求"
2025/04/07 00:25:21 收到 Python 服务 的响应: "Python processed: 测试请求"
2025/04/07 00:25:21 ========== Python 服务 调用完成 ==========
2025/04/07 00:25:21 正在向 Go 服务 发送请求，内容: "测试请求"
2025/04/07 00:25:21 收到 Go 服务 的响应: "Go processed: 测试请求"
2025/04/07 00:25:21 ========== Go 服务 调用完成 ==========
```