# 用GRPC实现通信
# 用Consul实现服务发现

import sys
sys.path.append('/root/go/src/github.com/jstang9527/buoy')
import time
import consul
# from threading import Thread
from concurrent import futures
from micro.proto.grpc.pocsuite_pb2 import VerifyResponse
from micro.proto.grpc.pocsuite_pb2_grpc import PocScanServicer, add_PocScanServicer_to_server, grpc


class Verifier(PocScanServicer):
    def Communication(self, request, context):
        print(request.Cmd)
        print(request.cmd)
        return VerifyResponse(msg="hello grpc consul {}".format(request.cmd))


class Consul:
    def __init__(self, host, port):
        # 连接注册中心
        self._consul = consul.Consul(host=host, port=port)

    def RegisterService(self, name, host, port, tags=None):
        # 注册服务
        tags = tags or []
        self._consul.agent.service.register(name, service_id=name, address=host, port=port, tags=tags, check=consul.Check().tcp(host, port, "5s", "30s", "30s"))

    def GetService(self, name):
        services = self._consul.agent.services()
        service = services.get(name)
        if not service:
            return None, None
        addr = "{0}:{1}".format(service['Address'], service['Port'])
        return service, addr

    # def Run(self):
    def ListenAndServer(self):
        # 2.注册服务
        self.RegisterService("test_python3", "172.31.50.249", 50051)
        # 1.运行守护程序
        print("listen server at [::]:50051")
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        add_PocScanServicer_to_server(Verifier, server)
        server.add_insecure_port("[::]:50051")
        server.start()
        try:
            while True:
                time.sleep(100)
        except KeyboardInterrupt:
            print("stop")
            server.stop(0)


if __name__ == "__main__":
    reg = Consul(host="172.31.50.249", port="8500")
    reg.ListenAndServer()
    check = consul.Check().tcp("172.31.50.249", 50051, "5s", "30s", "30s")
    print(check)
    # res = reg.GetService("maple")
    # print(res)
