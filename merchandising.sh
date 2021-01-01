#!/bin/bash 

protoc --go_out=plugins=grpc:. merchandising/merchandising.proto \
&& protoc-go-inject-tag -input=./merchandising/merchandising.pb.go \
&& git add . && git commit -m 'merchandising update' && git push