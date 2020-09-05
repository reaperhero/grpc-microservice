# 

- grpc 搭建

生成Prod.pb.go
```
protoc --go_out=plugins=grpc:../services Prod.proto
```


- grpc gateway

1. 安装依赖

```
https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/protoc-3.13.0-osx-x86_64.zip
mv protoc3/bin/* /usr/local/bin/
mv protoc3/include/* /usr/local/include/
```

2. 生成pb.go

```
1、Prod.pb.go
protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:../services  \
		*.proto

2、Prod.pb.gw.go
protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:../services \
		*.proto
```

3. 启动

```
go run rpcServer.go # server
go run rpcClient.go # client

go run httpServer.go # rpc端口的代理服务
http://localhost:8001/v1/prod/100
```




```
