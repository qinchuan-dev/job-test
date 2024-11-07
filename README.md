# 欢迎使用job-test
    
编译

```bash
cd proto
./gen-buf.sh
cd ../
go build
```


---
## 开发

该工程使用golang开发， 依赖grpc 和 buf； 

```bash
 
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  
  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
  go install github.com/golang/protobuf/protoc-gen-go@latest
  go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

  npm install -g swagger-combine

```

## http服务启动


查看启动命令支持的参数
```bash
./job-test start --help
Run the job-test process

Usage:
job-test start [flags]

Flags:
--gatewayPort string   the http gateway port (default ":8081")
--grpcPort string      The grpc port (default ":50051")
-h, --help             help for start

--pgur     string     Postgres url 
--redisurl string     Redis url 

```

启动http服务, 使用默认端口
```bash
#
./job-test start 
2023/10/21 19:01:14 Serving gRPC on localhost:50051
2023/10/21 19:01:14 Serving gRPC-Gateway on http://0.0.0.0:8081
```

启动http 服务, 自定义端口
```bash
# 使用自定义端口
./job-test start --gatewayPort :10000 --grpcPort :20000
2023/10/21 19:02:22 Serving gRPC on localhost:20000
2023/10/21 19:02:22 Serving gRPC-Gateway on http://0.0.0.0:10000

```
