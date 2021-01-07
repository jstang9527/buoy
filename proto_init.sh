#!/bin/bash
echo 'generate proto model'
protoc --proto_path=$GOPATH/src:. --go_out=. micro/proto/model/*.proto 
echo 'generate proto rpcapi'
protoc --proto_path=$GOPATH/src:. --micro_out=. micro/proto/rpcapi/*.proto

echo 'generate proto python model'
python3 -m grpc_tools.protoc -I micro/proto/model/ --python_out=micro/proto/model/ pocsuite.proto
echo 'generate proto python rpcapi'
python3 -m grpc_tools.protoc -I /root/go/src/ --grpc_python_out=micro/proto/rpcapi/ github.com/jstang9527/buoy/micro/proto/rpcapi/pocsuite.proto
if [ $? -ne 0 ];then
    python3 -m grpc_tools.protoc -I ../../../ --grpc_python_out=micro/proto/rpcapi/ github.com/jstang9527/buoy/micro/proto/rpcapi/pocsuite.proto
fi

python3 -m grpc_tools.protoc -I micro/proto/grpc/ --grpc_python_out=micro/proto/grpc/  --python_out=micro/proto/model/ pocsuite.proto

new:
python3 -m grpc_tools.protoc -I . --grpc_python_out=.  --python_out=. arith.proto
protoc -I . --go_out=plugins=grpc:. arith.proto