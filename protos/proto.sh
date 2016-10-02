#!/usr/bin/env bash

protoc --proto_path=. \
    --gofast_out=plugins=grpc:../src/feedpb \
    --
    ./*.proto
