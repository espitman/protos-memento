#!/bin/bash 

protoc --go_out=plugins=grpc:. movie/movie.proto \
&& protoc-go-inject-tag -input=./movie/movie.pb.go \
&& git add . && git commit -m 'movie update' && git push