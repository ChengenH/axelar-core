// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/snapshot/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("axelar/snapshot/v1beta1/service.proto", fileDescriptor_795649fbc059ac6e)
}
func init() {
	golang_proto.RegisterFile("axelar/snapshot/v1beta1/service.proto", fileDescriptor_795649fbc059ac6e)
}

var fileDescriptor_795649fbc059ac6e = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x1d, 0x0f, 0x1d, 0x16, 0x22, 0x58, 0x82, 0x40, 0x62, 0x10, 0xc1, 0x8b, 0xe2, 0x8c,
	0x16, 0x75, 0xf0, 0x18, 0x5d, 0x83, 0xd0, 0x5b, 0x97, 0x18, 0xd7, 0xc7, 0xb8, 0x64, 0xfb, 0xa6,
	0x99, 0xa7, 0xad, 0x44, 0x97, 0x3e, 0x41, 0xd0, 0x27, 0xe8, 0xd2, 0xe7, 0xe8, 0xd8, 0x51, 0xe8,
	0x52, 0xb7, 0x70, 0xfb, 0x20, 0xa1, 0xbb, 0x52, 0x1a, 0x0b, 0x9e, 0x66, 0x86, 0xf9, 0xfd, 0xdf,
	0xff, 0x07, 0xcf, 0xab, 0xaa, 0x18, 0x86, 0xca, 0x4a, 0x17, 0x29, 0xe3, 0x06, 0x48, 0x72, 0xdc,
	0xea, 0x01, 0xa9, 0x96, 0x74, 0x60, 0xc7, 0x61, 0x00, 0xc2, 0x58, 0x24, 0xf4, 0xf7, 0x52, 0x4c,
	0x2c, 0x31, 0x91, 0x61, 0xa5, 0x5d, 0x8d, 0x1a, 0x17, 0x8c, 0x9c, 0xdf, 0x52, 0xbc, 0xb4, 0xaf,
	0x11, 0xf5, 0x10, 0xa4, 0x32, 0xa1, 0x54, 0x51, 0x84, 0xa4, 0x28, 0xc4, 0xc8, 0x65, 0xbf, 0xe5,
	0xbc, 0x4e, 0x8a, 0x53, 0xe2, 0xe0, 0xb3, 0xe8, 0x79, 0x67, 0x4e, 0x77, 0x53, 0x07, 0xff, 0x85,
	0x79, 0xdb, 0x1d, 0xd0, 0xa1, 0x23, 0xb0, 0xe7, 0x16, 0xe3, 0x89, 0xdf, 0x10, 0x39, 0x42, 0x62,
	0x85, 0xeb, 0xc0, 0xcd, 0x08, 0x1c, 0x95, 0xc4, 0xa6, 0xb8, 0x33, 0x18, 0x39, 0xa8, 0x1c, 0x3f,
	0xbc, 0x7f, 0x3f, 0x15, 0x9b, 0x95, 0xba, 0x5c, 0x57, 0xb5, 0x7f, 0x79, 0x79, 0x67, 0xe6, 0xc7,
	0xa5, 0xea, 0xf7, 0xed, 0x7d, 0x9b, 0xd5, 0xfc, 0x67, 0xe6, 0xed, 0x9c, 0x82, 0x0a, 0x28, 0x1c,
	0x2b, 0x82, 0x54, 0x55, 0xe6, 0x76, 0xaf, 0x91, 0x4b, 0xd9, 0xe6, 0xe6, 0x81, 0x4c, 0xb7, 0xbe,
	0xd0, 0xad, 0x56, 0xca, 0xff, 0x74, 0xfb, 0xab, 0x89, 0x36, 0xab, 0x9d, 0x74, 0xdf, 0x66, 0x9c,
	0x4d, 0x67, 0x9c, 0x7d, 0xcd, 0x38, 0x7b, 0x4c, 0x78, 0xe1, 0x35, 0xe1, 0x6c, 0x9a, 0xf0, 0xc2,
	0x47, 0xc2, 0x0b, 0x17, 0x47, 0x3a, 0xa4, 0xc1, 0xa8, 0x27, 0x02, 0xbc, 0xce, 0x86, 0x45, 0x40,
	0xb7, 0x68, 0xaf, 0xb2, 0x57, 0x23, 0x40, 0x0b, 0x32, 0xfe, 0x6d, 0xa0, 0x89, 0x01, 0xd7, 0xdb,
	0x5a, 0xec, 0xed, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xdb, 0xb0, 0x19, 0x4f, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	// RegisterProxy defines a method for registering a proxy account that can act
	// in a validator account's stead.
	RegisterProxy(ctx context.Context, in *RegisterProxyRequest, opts ...grpc.CallOption) (*RegisterProxyResponse, error)
	// DeactivateProxy defines a method for deregistering a proxy account.
	DeactivateProxy(ctx context.Context, in *DeactivateProxyRequest, opts ...grpc.CallOption) (*DeactivateProxyResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) RegisterProxy(ctx context.Context, in *RegisterProxyRequest, opts ...grpc.CallOption) (*RegisterProxyResponse, error) {
	out := new(RegisterProxyResponse)
	err := c.cc.Invoke(ctx, "/axelar.snapshot.v1beta1.MsgService/RegisterProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) DeactivateProxy(ctx context.Context, in *DeactivateProxyRequest, opts ...grpc.CallOption) (*DeactivateProxyResponse, error) {
	out := new(DeactivateProxyResponse)
	err := c.cc.Invoke(ctx, "/axelar.snapshot.v1beta1.MsgService/DeactivateProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// RegisterProxy defines a method for registering a proxy account that can act
	// in a validator account's stead.
	RegisterProxy(context.Context, *RegisterProxyRequest) (*RegisterProxyResponse, error)
	// DeactivateProxy defines a method for deregistering a proxy account.
	DeactivateProxy(context.Context, *DeactivateProxyRequest) (*DeactivateProxyResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) RegisterProxy(ctx context.Context, req *RegisterProxyRequest) (*RegisterProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProxy not implemented")
}
func (*UnimplementedMsgServiceServer) DeactivateProxy(ctx context.Context, req *DeactivateProxyRequest) (*DeactivateProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateProxy not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_RegisterProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.snapshot.v1beta1.MsgService/RegisterProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterProxy(ctx, req.(*RegisterProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_DeactivateProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).DeactivateProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.snapshot.v1beta1.MsgService/DeactivateProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).DeactivateProxy(ctx, req.(*DeactivateProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.snapshot.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterProxy",
			Handler:    _MsgService_RegisterProxy_Handler,
		},
		{
			MethodName: "DeactivateProxy",
			Handler:    _MsgService_DeactivateProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/snapshot/v1beta1/service.proto",
}
