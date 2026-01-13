#!/bin/bash

mkdir -p gen

protoc --proto_path=. --go_out=gen --go-grpc_out=gen order.proto
protoc --proto_path=. --go_out=gen --go-grpc_out=gen product.proto