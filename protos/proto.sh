#!/usr/bin/env bash

GOOGLE_APIPATH=../vendor/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

protoc --proto_path=.:$GOOGLE_APIPATH \
    --gofast_out=plugins=grpc:../src/feedpb \
    --grpc-gateway_out=logtostderr=true:../src/feedpb \
    ./*.proto
