#!/bin/bash 

protoc --go_out=plugins=grpc:. person/person.proto \
&& protoc-go-inject-tag -input=./person/person.pb.go \
&& git add . && git commit -m 'person update' && git push