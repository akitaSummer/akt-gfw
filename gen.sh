#!/bin/sh

if [ ! -d "./network/protocol/gen/" ];then
    mkdir ./network/protocol/gen/
fi

protoc \
    --validate_out="lang=go:./network/protocol/gen/" \
    --go_out=./network/protocol/gen/ \
    ./network/protocol/proto/*.proto 