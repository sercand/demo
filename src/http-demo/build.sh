#!/bin/bash

TAG_NAME=sercand/demo-node-js:1.0

docker build --rm -t ${TAG_NAME} .

if [ "$1" = "push" ]; then
    docker push ${TAG_NAME}
fi
