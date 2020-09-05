# grpc-microservice 

## 使用教程

### 基础rpc服务

[基础rpc服务](https://github.com/reaperhero/grpc-microservice/tree/master/demo01)
```
protoc --go_out=plugins=grpc:../services Prod.proto
```


### grpc gateway


[http rpc代理](https://github.com/reaperhero/grpc-microservice/tree/master/demo02)

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

### proto数组

[数组返回](https://github.com/reaperhero/grpc-microservice/tree/master/demo03)
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


```
go run rpcServer.go 

go run rpcClient.go 
[prod_stock:12  prod_stock:13  prod_stock:14 ]
 

go run httpServer.go # rpc端口的代理服务

http://localhost:8001/v1/prods/4  
{
    "prodres": [
        {
            "prod_stock": 12
        },
        {
            "prod_stock": 13
        },
        {
            "prod_stock": 14
        }
    ]
}
```


### 枚举使用

[枚举](https://github.com/reaperhero/grpc-microservice/tree/master/demo04)

```shell script
protoc --go_out=plugins=grpc:../services Prod.proto
go run rpcServer.go


go run rpcClient.go

```


### 外部proto

[外部proto文件](https://github.com/reaperhero/grpc-microservice/tree/master/demo05)
```
protoc --go_out=plugins=grpc:../services Prod.proto
protoc --go_out=plugins=grpc:../services Models.proto

go run rpcServer.go


go run rpcClient.go
prod_id:101 prod_name:"testname" prod_price:20.5 
```

