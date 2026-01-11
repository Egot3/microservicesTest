module github.com/Egot3/microservicesTest/product-service

go 1.21

require (
  google.golang.org/grpc v1.60.1
  google.golang.org/protobuf v1.32.0
)

replace github.com/Egot3/microservicesTest/proto/gen => ../proto/gen

require (
  golang.org/x/net v0.20.0 // indirect
  golang.org/x/sys v0.16.0 // indirect
  golang.org/x/text v0.14.0 // indirect
  google.golang.org/grpc v1.60.1
    google.golang.org/protobuf v1.32.0
  google.golang.org/genproto/googleapis/rpc v0.0.0-20240116215550-a9fa1716bcac // indirect
)