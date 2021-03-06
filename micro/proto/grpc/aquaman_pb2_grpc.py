# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from micro.proto.grpc import aquaman_pb2 as aquaman__pb2


class DomainStub(object):
    """定义域名类服务
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Resolv = channel.unary_unary(
                '/grpc.Domain/Resolv',
                request_serializer=aquaman__pb2.ResolvRequest.SerializeToString,
                response_deserializer=aquaman__pb2.ResolvResponse.FromString,
                )
        self.Analysis = channel.unary_unary(
                '/grpc.Domain/Analysis',
                request_serializer=aquaman__pb2.AnlsRequest.SerializeToString,
                response_deserializer=aquaman__pb2.AnlsResponse.FromString,
                )


class DomainServicer(object):
    """定义域名类服务
    """

    def Resolv(self, request, context):
        """定义域名解析方法, 资产需要解析IP存表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Analysis(self, request, context):
        """定义域名查询、子域爆破方法
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DomainServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Resolv': grpc.unary_unary_rpc_method_handler(
                    servicer.Resolv,
                    request_deserializer=aquaman__pb2.ResolvRequest.FromString,
                    response_serializer=aquaman__pb2.ResolvResponse.SerializeToString,
            ),
            'Analysis': grpc.unary_unary_rpc_method_handler(
                    servicer.Analysis,
                    request_deserializer=aquaman__pb2.AnlsRequest.FromString,
                    response_serializer=aquaman__pb2.AnlsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.Domain', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Domain(object):
    """定义域名类服务
    """

    @staticmethod
    def Resolv(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.Domain/Resolv',
            aquaman__pb2.ResolvRequest.SerializeToString,
            aquaman__pb2.ResolvResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Analysis(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.Domain/Analysis',
            aquaman__pb2.AnlsRequest.SerializeToString,
            aquaman__pb2.AnlsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)


class HostStub(object):
    """定义主机信息识别服务
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Location = channel.unary_unary(
                '/grpc.Host/Location',
                request_serializer=aquaman__pb2.LocRequest.SerializeToString,
                response_deserializer=aquaman__pb2.LocResponse.FromString,
                )
        self.Alive = channel.unary_unary(
                '/grpc.Host/Alive',
                request_serializer=aquaman__pb2.AlvRequest.SerializeToString,
                response_deserializer=aquaman__pb2.AlvResponse.FromString,
                )
        self.Detail = channel.unary_unary(
                '/grpc.Host/Detail',
                request_serializer=aquaman__pb2.DetlRequest.SerializeToString,
                response_deserializer=aquaman__pb2.DetlResponse.FromString,
                )


class HostServicer(object):
    """定义主机信息识别服务
    """

    def Location(self, request, context):
        """定义IP信息查询方法
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Alive(self, request, context):
        """定义主机存活方法
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Detail(self, request, context):
        """定义服务识别方法
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_HostServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Location': grpc.unary_unary_rpc_method_handler(
                    servicer.Location,
                    request_deserializer=aquaman__pb2.LocRequest.FromString,
                    response_serializer=aquaman__pb2.LocResponse.SerializeToString,
            ),
            'Alive': grpc.unary_unary_rpc_method_handler(
                    servicer.Alive,
                    request_deserializer=aquaman__pb2.AlvRequest.FromString,
                    response_serializer=aquaman__pb2.AlvResponse.SerializeToString,
            ),
            'Detail': grpc.unary_unary_rpc_method_handler(
                    servicer.Detail,
                    request_deserializer=aquaman__pb2.DetlRequest.FromString,
                    response_serializer=aquaman__pb2.DetlResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.Host', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Host(object):
    """定义主机信息识别服务
    """

    @staticmethod
    def Location(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.Host/Location',
            aquaman__pb2.LocRequest.SerializeToString,
            aquaman__pb2.LocResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Alive(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.Host/Alive',
            aquaman__pb2.AlvRequest.SerializeToString,
            aquaman__pb2.AlvResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Detail(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.Host/Detail',
            aquaman__pb2.DetlRequest.SerializeToString,
            aquaman__pb2.DetlResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)


class WebScrapStub(object):
    """定义Web扫描服务
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Spider = channel.unary_unary(
                '/grpc.WebScrap/Spider',
                request_serializer=aquaman__pb2.SpiRequest.SerializeToString,
                response_deserializer=aquaman__pb2.SpiResponse.FromString,
                )


class WebScrapServicer(object):
    """定义Web扫描服务
    """

    def Spider(self, request, context):
        """定义Web爬虫
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_WebScrapServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Spider': grpc.unary_unary_rpc_method_handler(
                    servicer.Spider,
                    request_deserializer=aquaman__pb2.SpiRequest.FromString,
                    response_serializer=aquaman__pb2.SpiResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.WebScrap', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class WebScrap(object):
    """定义Web扫描服务
    """

    @staticmethod
    def Spider(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.WebScrap/Spider',
            aquaman__pb2.SpiRequest.SerializeToString,
            aquaman__pb2.SpiResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)


class VulStub(object):
    """定义扫描服务
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Verify = channel.unary_unary(
                '/grpc.Vul/Verify',
                request_serializer=aquaman__pb2.PocRequest.SerializeToString,
                response_deserializer=aquaman__pb2.PocResponse.FromString,
                )
        self.Hydra = channel.unary_unary(
                '/grpc.Vul/Hydra',
                request_serializer=aquaman__pb2.AuthRequest.SerializeToString,
                response_deserializer=aquaman__pb2.AuthResponse.FromString,
                )
        self.Trap = channel.unary_unary(
                '/grpc.Vul/Trap',
                request_serializer=aquaman__pb2.TrapRequest.SerializeToString,
                response_deserializer=aquaman__pb2.TrapResponse.FromString,
                )


class VulServicer(object):
    """定义扫描服务
    """

    def Verify(self, request, context):
        """定义漏洞验证方法
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Hydra(self, request, context):
        """定义权限爆破方法
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Trap(self, request, context):
        """定义蜜罐识别方法
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_VulServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Verify': grpc.unary_unary_rpc_method_handler(
                    servicer.Verify,
                    request_deserializer=aquaman__pb2.PocRequest.FromString,
                    response_serializer=aquaman__pb2.PocResponse.SerializeToString,
            ),
            'Hydra': grpc.unary_unary_rpc_method_handler(
                    servicer.Hydra,
                    request_deserializer=aquaman__pb2.AuthRequest.FromString,
                    response_serializer=aquaman__pb2.AuthResponse.SerializeToString,
            ),
            'Trap': grpc.unary_unary_rpc_method_handler(
                    servicer.Trap,
                    request_deserializer=aquaman__pb2.TrapRequest.FromString,
                    response_serializer=aquaman__pb2.TrapResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.Vul', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Vul(object):
    """定义扫描服务
    """

    @staticmethod
    def Verify(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.Vul/Verify',
            aquaman__pb2.PocRequest.SerializeToString,
            aquaman__pb2.PocResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Hydra(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.Vul/Hydra',
            aquaman__pb2.AuthRequest.SerializeToString,
            aquaman__pb2.AuthResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Trap(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.Vul/Trap',
            aquaman__pb2.TrapRequest.SerializeToString,
            aquaman__pb2.TrapResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
