GO gRPC
-------
This is an example of implementing gRPC using Golang.  
This includes a stream server, stream client, and gRPC client implementation.

**Generate new proto**  
`protoc --go_out=proto/message --go_opt=module=mbooke/message --go-grpc_out=proto/message --go-grpc_opt=module=mbooke/message proto/message.proto`  

`protoc --go_out=proto/location --go_opt=module=mbooke/location --go-grpc_out=proto/location --go-grpc_opt=module=mbooke/location proto/location.proto`

**Run Go**
- run main  
  `go run .`
  
- run stream client  
  `go run grpcclient/streamclient/main.go`
  
- run client  
  `go run grpcclient/client/main.go`

**Build Linux**  
  `env GOOS=linux GOARCH=amd64 go build`
