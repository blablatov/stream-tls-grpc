## Service and Client - Go Implementation

## Building and Running Service

In order to build, Go to ``Go`` module directory location `stream-tls-grpc/tls-service` and execute the following
 shell command:
```
go build -v 
```

In order to run, Go to ``Go`` module directory location `stream-tls-grpc/tls-service` and execute the following
shell command:

```
./tls-server
```

## Building and Running Client   

In order to build, Go to ``Go`` module directory location `stream-tls-grpc/tls-client` and execute the following shell command:
```
go build -v 
```

In order to run, Go to ``Go`` module directory location `stream-tls-grpc/tls-client` and execute the following shell command:

```
./client
```

## Additional Information

### Generate Server and Client side code   
Go to ``Go`` module directory location `stream-tls-grpc/tls-proto` and execute the following shell commands:    
``` 
protoc product_info.proto --go_out=./ --go-grpc_out=./
protoc product_info.proto --go-grpc_out=require_unimplemented_servers=false:.
``` 
