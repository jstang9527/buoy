// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arith.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ArithRequest 请求结构
type ArithRequest struct {
	Num1                 int32    `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2                 int32    `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithRequest) Reset()         { *m = ArithRequest{} }
func (m *ArithRequest) String() string { return proto.CompactTextString(m) }
func (*ArithRequest) ProtoMessage()    {}
func (*ArithRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_211289c5d02710c5, []int{0}
}

func (m *ArithRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithRequest.Unmarshal(m, b)
}
func (m *ArithRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithRequest.Marshal(b, m, deterministic)
}
func (m *ArithRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithRequest.Merge(m, src)
}
func (m *ArithRequest) XXX_Size() int {
	return xxx_messageInfo_ArithRequest.Size(m)
}
func (m *ArithRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArithRequest proto.InternalMessageInfo

func (m *ArithRequest) GetNum1() int32 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *ArithRequest) GetNum2() int32 {
	if m != nil {
		return m.Num2
	}
	return 0
}

// ArithResponse 响应结构
type ArithResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithResponse) Reset()         { *m = ArithResponse{} }
func (m *ArithResponse) String() string { return proto.CompactTextString(m) }
func (*ArithResponse) ProtoMessage()    {}
func (*ArithResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_211289c5d02710c5, []int{1}
}

func (m *ArithResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithResponse.Unmarshal(m, b)
}
func (m *ArithResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithResponse.Marshal(b, m, deterministic)
}
func (m *ArithResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithResponse.Merge(m, src)
}
func (m *ArithResponse) XXX_Size() int {
	return xxx_messageInfo_ArithResponse.Size(m)
}
func (m *ArithResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArithResponse proto.InternalMessageInfo

func (m *ArithResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*ArithRequest)(nil), "grpc.ArithRequest")
	proto.RegisterType((*ArithResponse)(nil), "grpc.ArithResponse")
}

func init() { proto.RegisterFile("arith.proto", fileDescriptor_211289c5d02710c5) }

var fileDescriptor_211289c5d02710c5 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0xca, 0x2c,
	0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2f, 0x2a, 0x48, 0x56, 0x32, 0xe3,
	0xe2, 0x71, 0x04, 0x09, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4,
	0x95, 0xe6, 0x1a, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0xd9, 0x50, 0x31, 0x23, 0x09,
	0x26, 0xb8, 0x98, 0x91, 0x92, 0x3a, 0x17, 0x2f, 0x54, 0x5f, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa,
	0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0x54, 0x2b, 0x94, 0x67, 0x54, 0xc6,
	0xc5, 0x0a, 0x56, 0x28, 0x64, 0xca, 0xc5, 0x11, 0x91, 0x99, 0x98, 0x97, 0xee, 0x95, 0x99, 0x28,
	0x24, 0xa4, 0x07, 0xb2, 0x5c, 0x0f, 0xd9, 0x66, 0x29, 0x61, 0x14, 0x31, 0x88, 0xa9, 0x4a, 0x0c,
	0x42, 0x66, 0x5c, 0x9c, 0x30, 0x6d, 0x79, 0x24, 0xe8, 0x73, 0x62, 0x8b, 0x02, 0x7b, 0x30, 0x89,
	0x0d, 0xec, 0x5b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xb6, 0xd2, 0x6a, 0xfc, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArithClient is the client API for Arith service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArithClient interface {
	// 定义相加方法
	XiangJia(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error)
	// 定义相减方法
	XiangJian(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error)
}

type arithClient struct {
	cc *grpc.ClientConn
}

func NewArithClient(cc *grpc.ClientConn) ArithClient {
	return &arithClient{cc}
}

func (c *arithClient) XiangJia(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error) {
	out := new(ArithResponse)
	err := c.cc.Invoke(ctx, "/grpc.Arith/XiangJia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithClient) XiangJian(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error) {
	out := new(ArithResponse)
	err := c.cc.Invoke(ctx, "/grpc.Arith/XiangJian", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithServer is the server API for Arith service.
type ArithServer interface {
	// 定义相加方法
	XiangJia(context.Context, *ArithRequest) (*ArithResponse, error)
	// 定义相减方法
	XiangJian(context.Context, *ArithRequest) (*ArithResponse, error)
}

// UnimplementedArithServer can be embedded to have forward compatible implementations.
type UnimplementedArithServer struct {
}

func (*UnimplementedArithServer) XiangJia(ctx context.Context, req *ArithRequest) (*ArithResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XiangJia not implemented")
}
func (*UnimplementedArithServer) XiangJian(ctx context.Context, req *ArithRequest) (*ArithResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XiangJian not implemented")
}

func RegisterArithServer(s *grpc.Server, srv ArithServer) {
	s.RegisterService(&_Arith_serviceDesc, srv)
}

func _Arith_XiangJia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithServer).XiangJia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Arith/XiangJia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithServer).XiangJia(ctx, req.(*ArithRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arith_XiangJian_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithServer).XiangJian(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Arith/XiangJian",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithServer).XiangJian(ctx, req.(*ArithRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Arith_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Arith",
	HandlerType: (*ArithServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "XiangJia",
			Handler:    _Arith_XiangJia_Handler,
		},
		{
			MethodName: "XiangJian",
			Handler:    _Arith_XiangJian_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arith.proto",
}