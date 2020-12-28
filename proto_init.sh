#!/bin/bash
echo 'generate proto model'
protoc --proto_path=$GOPATH/src:. --go_out=. micro/proto/model/*.proto 
echo 'generate proto rpcapi'
protoc --proto_path=$GOPATH/src:. --micro_out=. micro/proto/rpcapi/*.proto