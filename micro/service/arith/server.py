import re
import time
from concurrent import futures
import sys
sys.path.append("/root/go/src/github.com/jstang9527/buoy")
from micro.proto.grpc.arith_pb2 import ArithResponse
from micro.proto.grpc.arith_pb2_grpc import ArithServicer, add_ArithServicer_to_server, grpc
from micro.handler.grpc_registry import GRPCServiceBase


class Arither(ArithServicer):
    def XiangJia(self, request, context):
        data = re.sub(r"[\n\r\t]", ", ", "{}".format(request))
        print("Recv Data From Client, Data: ", data)
        return ArithResponse(result=request.num1 + request.num2)

    def XiangJian(self, request, context):
        print("Recv Data From Client, Data: ", request)
        return ArithResponse(result=request.num1 - request.num2)


class ArithServer(GRPCServiceBase):
    def __init__(self, registry_host, registry_port, server_addr=None, server_port=None):
        super(ArithServer, self).__init__(registry_host, registry_port, server_addr, server_port)

    def ListenAndServer(self):
        if not self.server_port:
            print("Faild generate service port2.")
            return
        # 1.运行守护程序
        print("listen server on [::]:{}".format(self.server_port))
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        add_ArithServicer_to_server(Arither(), server)
        server.add_insecure_port("[::]:{}".format(self.server_port))
        server.start()
        # 1.注册服务
        self.RegisterService("test_python3", "172.31.50.249", self.server_port)
        try:
            while True:
                time.sleep(100)
        except KeyboardInterrupt:
            print("stop")
            server.stop(0)


if __name__ == "__main__":
    # reg = ArithServer("172.31.50.249", "8500")
    # reg.ListenAndServer()
    # check = consul.Check().tcp("172.31.50.249", 50051, "5s", "30s", "30s")
    # print(check)
    x = ArithServer("172.31.50.249", "8500")
    x.ListenAndServer()
