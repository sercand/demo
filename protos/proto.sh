#!/usr/bin/env bash

GOOGLE_APIPATH=../vendor/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

protoc --proto_path=.:$GOOGLE_APIPATH \
    --gofast_out=plugins=grpc:../src/feedpb \
    --java_out=../src/feedpb \
    --cpp_out=../src/feedpb \
    --csharp_out=../src/feedpb \
    --python_out=../src/feedpb \
    --grpc-gateway_out=logtostderr=true:../src/feedpb \
    ./*.proto


python -m grpc.tools.protoc \
    --proto_path=. \
    --python_out=../src/feed-python \
    --grpc_python_out=../src/feed-python \
    ./common.proto ./color.proto ./feedprovider.proto