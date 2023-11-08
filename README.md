[![Go](https://github.com/blablatov/stream-tls-grpc/actions/workflows/stream-tls-grpc.yml/badge.svg)](https://github.com/blablatov/stream-tls-grpc/actions/workflows/stream-tls-grpc.yml)
### Building and Running gPRC service

In order to build, Go to ``Go`` module directory location `stream-tls-grpc/tls-service` and execute the following
 shell command:
```
go build -v 
./tls-service
```   

### Building and Running gRPC client   

In order to build, Go to ``Go`` module directory location `stream-tls-grpc/tls-client` and execute the following shell command:
```
go build -v 
./tls-client
```  


### Generates Server and Client side code via proto-file     
Go to ``Go`` module directory location `stream-tls-grpc/tls-proto` and execute the following shell commands:    
``` 
protoc product_info.proto --go_out=./ --go-grpc_out=./
protoc product_info.proto --go-grpc_out=require_unimplemented_servers=false:.
``` 
