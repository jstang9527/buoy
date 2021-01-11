import sys
sys.path.append("/root/go/src/github.com/jstang9527/buoy")
import re
import time
from concurrent import futures
import argparse
from argparse import RawTextHelpFormatter
from micro.proto.grpc.pocsuite_pb2 import PocResponse
from micro.proto.grpc.pocsuite_pb2_grpc import add_PocScanServicer_to_server, PocScanServicer, grpc
from micro.handler.grpc_registry import GRPCServiceBase
from handler import Pocsuite


class PocScanner(PocScanServicer):
    def Verify(self, request, context):
        result = {'VerifyInfo': None, 'ExploitInfo': None, 'WebshellInfo': None, 'TrojanInfo': None}
        print("Recv Data From Client, Data: ", re.sub(r"[\n\r\t]", ",", "{}".format(request)))

        # request.command
        p = Pocsuite(request.target, request.poc_plugins, request.asset_id)  # target: str, pocs: list, asset_id: str = None, command: str = None
        resp = p.Verify(request.exploit)                           # exploit: bool

        if not resp:
            return PocResponse()
        if 'ExploitInfo' in resp.keys():
            result['ExploitInfo'] = resp['ExploitInfo']
        if 'WebshellInfo' in resp.keys():
            result['WebshellInfo'] = resp['WebshellInfo']
        if 'TrojanInfo' in resp.keys():
            result['TrojanInfo'] = resp['TrojanInfo']
        print(result)
        return PocResponse(VerifyInfo=resp['VerifyInfo'], ExploitInfo=result['ExploitInfo'], WebshellInfo=result['WebshellInfo'], TrojanInfo=result['TrojanInfo'])


class PocScanServer(GRPCServiceBase):
    def __init__(self, registry_host, registry_port, server_name, server_addr, server_port=None):
        super(PocScanServer, self).__init__(registry_host, registry_port, server_addr, server_port)
        self.server_name = server_name

    def ListenAndServer(self):
        if not self.server_port:
            print("Faild generate service port2.")
            return
        # 1.运行守护程序
        print("listen server on [::]:{}".format(self.server_port))
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        add_PocScanServicer_to_server(PocScanner(), server)
        server.add_insecure_port("[::]:{}".format(self.server_port))
        server.start()
        # 2.注册服务
        # self.RegisterService("test_python3", "172.31.50.249", self.server_port)
        self.RegisterService(self.server_name, self.server_addr, self.server_port)
        try:
            while True:
                time.sleep(100)
        except KeyboardInterrupt:
            print("stop")
            server.stop(0)


def user_args():
    print("\033[33m\t\t---------- PocScannerServer ----------\033[0m")
    description = "\033[32mProvide Service of Airth by GRPC.\033[0m"
    example = "\n\nexample:\n\t[-] 向注册中心(127.0.0.1:8500)注册主机(192.168.1.11:50051)服务:\n\t" \
        "python3 server.py --registry_host 127.0.0.1 --registry_port 8500 --server_name pocsuite --server_addr 192.168.1.11 --server_port 50051\n\t"
    description += example
    parser = argparse.ArgumentParser(description=description, prog='python2 main.py', formatter_class=RawTextHelpFormatter)                        # description参数可以用于插入描述脚本用途的信息，可以为空
    parser.add_argument('--registry_host', required=True, help='\t输入注册中心IP(必填)')               # 添加--verbose标签，标签别名可以为-v，这里action的意思是当读取的参数中出现--verbose/-v的时候
    parser.add_argument('--registry_port', required=True, help='\t输入注册中心端口(必填)')
    parser.add_argument('--server_name', required=True, help='\t输入服务名(必填)')
    parser.add_argument('--server_addr', required=True, help='\t输入绑定本机IP(必填)')
    parser.add_argument('--server_port', type=int, required=False, help='\t输入服务端口(选填), 默认随机端口')

    args = parser.parse_args(sys.argv[1:])                                             # 将变量以标签-值的字典形式存入args字典
    return{'registry_host': args.registry_host, 'registry_port': args.registry_port, 'server_name': args.server_name, 'server_addr': args.server_addr, 'server_port': args.server_port}


if __name__ == "__main__":
    args = user_args()
    x = PocScanServer(args['registry_host'], args['registry_port'], args['server_name'], args['server_addr'], args['server_port'])
    x.ListenAndServer()
