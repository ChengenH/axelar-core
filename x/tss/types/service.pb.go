// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/tss/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/axelarnetwork/axelar-core/x/snapshot/types"
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

func init() { proto.RegisterFile("axelar/tss/v1beta1/service.proto", fileDescriptor_03d0b5f2955aa837) }
func init() {
	golang_proto.RegisterFile("axelar/tss/v1beta1/service.proto", fileDescriptor_03d0b5f2955aa837)
}

var fileDescriptor_03d0b5f2955aa837 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x4d, 0x4f, 0x13, 0x41,
	0x18, 0xc7, 0xbb, 0x1c, 0x30, 0x8c, 0x7a, 0x99, 0x40, 0x24, 0x05, 0x57, 0x5c, 0x04, 0x5f, 0x62,
	0x77, 0x79, 0xf1, 0xe4, 0x4d, 0xa2, 0x89, 0x86, 0x40, 0x90, 0x1a, 0x0f, 0x5c, 0xc8, 0xec, 0xfa,
	0x30, 0x9d, 0x50, 0x76, 0xca, 0xcc, 0x2c, 0x76, 0x21, 0xc6, 0x48, 0xbc, 0x78, 0x30, 0xd1, 0x98,
	0xf8, 0x15, 0xfc, 0x1a, 0x1e, 0x3d, 0x92, 0x78, 0xf1, 0x68, 0xa8, 0x1f, 0xc4, 0xcc, 0x76, 0x76,
	0xd9, 0xd6, 0xb1, 0x85, 0x5b, 0xdb, 0xe7, 0xf7, 0xcc, 0xff, 0xb7, 0xcf, 0xcc, 0x4e, 0xd1, 0x0c,
	0x69, 0x43, 0x93, 0x88, 0x40, 0x49, 0x19, 0x1c, 0x2c, 0x86, 0xa0, 0xc8, 0x62, 0x20, 0x41, 0x1c,
	0xb0, 0x08, 0xfc, 0x96, 0xe0, 0x8a, 0x63, 0xdc, 0x25, 0x7c, 0x25, 0xa5, 0x6f, 0x88, 0xea, 0x38,
	0xe5, 0x94, 0x67, 0xe5, 0x40, 0x7f, 0xea, 0x92, 0xd5, 0x69, 0xca, 0x39, 0x6d, 0x42, 0x40, 0x5a,
	0x2c, 0x20, 0x71, 0xcc, 0x15, 0x51, 0x8c, 0xc7, 0xd2, 0x54, 0xf3, 0x24, 0x19, 0x93, 0x96, 0x6c,
	0x70, 0x55, 0xc4, 0xa9, 0xb6, 0x21, 0xa6, 0x2c, 0x2e, 0x45, 0xd1, 0xb5, 0x14, 0xf7, 0x13, 0x10,
	0x69, 0xb7, 0xbe, 0xf4, 0x19, 0x21, 0xb4, 0x26, 0x69, 0xbd, 0xeb, 0x8e, 0xbf, 0x39, 0x68, 0x7c,
	0x13, 0x28, 0x93, 0x0a, 0xc4, 0x93, 0xb6, 0x02, 0x11, 0x93, 0xe6, 0x2a, 0xa4, 0x12, 0x07, 0xfe,
	0xbf, 0xcf, 0xe3, 0xdb, 0xc8, 0x4d, 0xd8, 0x4f, 0x40, 0xaa, 0xea, 0xc2, 0xf9, 0x1b, 0x64, 0x8b,
	0xc7, 0x12, 0xbc, 0xfb, 0xc7, 0x3f, 0xff, 0x7c, 0x19, 0x99, 0xf7, 0x6e, 0x06, 0x25, 0x67, 0x61,
	0x3a, 0x6a, 0x60, 0x5a, 0x6a, 0xbb, 0x90, 0x3e, 0x74, 0xee, 0xe1, 0x43, 0x34, 0xf6, 0x14, 0x88,
	0x50, 0x2b, 0x40, 0x14, 0xbe, 0x65, 0x0b, 0x2b, 0xca, 0xb9, 0xd2, 0xdc, 0x10, 0xca, 0x78, 0xcc,
	0x64, 0x1e, 0x55, 0x6f, 0xa2, 0xec, 0xd1, 0xd0, 0x58, 0x08, 0x44, 0xe9, 0xec, 0x63, 0x07, 0x5d,
	0xae, 0x2b, 0x22, 0xd4, 0x2a, 0xa4, 0x14, 0x62, 0x3c, 0x6f, 0x5b, 0xb8, 0x04, 0xe4, 0x02, 0xb7,
	0x87, 0x72, 0x46, 0xc1, 0xcb, 0x14, 0xa6, 0xbd, 0x6b, 0x65, 0x05, 0x79, 0x06, 0x6a, 0x89, 0x77,
	0x0e, 0x1a, 0xdf, 0x10, 0x3c, 0x02, 0x29, 0xbb, 0x3f, 0xbe, 0x10, 0x64, 0x67, 0x87, 0x45, 0xf6,
	0xad, 0xb2, 0x91, 0x03, 0xb7, 0xca, 0xde, 0x60, 0xfc, 0x46, 0x33, 0xbf, 0x0a, 0x7e, 0x8b, 0xc6,
	0x36, 0xf5, 0x79, 0x85, 0x55, 0x48, 0xed, 0x9b, 0x50, 0x94, 0x07, 0x6e, 0x42, 0x89, 0x32, 0x09,
	0x73, 0x59, 0xc2, 0x0d, 0xaf, 0x5a, 0x9e, 0x00, 0x91, 0x92, 0xd1, 0x38, 0x38, 0x8a, 0x1a, 0x84,
	0xc5, 0x6f, 0xf4, 0x10, 0x22, 0x84, 0x5e, 0x72, 0x05, 0x1b, 0x49, 0xa8, 0x0d, 0xac, 0x6b, 0x9f,
	0xd5, 0x73, 0x85, 0xf9, 0x61, 0x58, 0xdf, 0x53, 0x1e, 0x21, 0x6c, 0xa6, 0x51, 0x67, 0xb4, 0x18,
	0x73, 0x6d, 0xc0, 0xd4, 0x4a, 0x5c, 0x1e, 0xea, 0x9f, 0x17, 0xef, 0x0b, 0xdf, 0x42, 0x97, 0xb4,
	0x5a, 0x9d, 0x51, 0xec, 0xfd, 0xcf, 0xbb, 0xce, 0x68, 0x1e, 0x33, 0x3b, 0x90, 0xe9, 0x5b, 0xfb,
	0xbd, 0x83, 0x26, 0xea, 0x49, 0xb8, 0xc7, 0xd4, 0x5a, 0xd2, 0x54, 0x4c, 0x32, 0xda, 0x9d, 0x80,
	0xc4, 0xd6, 0x23, 0x61, 0x45, 0xf3, 0xe0, 0xc5, 0x0b, 0x74, 0xf4, 0x69, 0x7c, 0x74, 0xd0, 0x64,
	0x2f, 0xa9, 0x07, 0x42, 0x54, 0x22, 0x40, 0xe2, 0xe5, 0xe1, 0xeb, 0x9e, 0xd1, 0xb9, 0xcc, 0x83,
	0x8b, 0x35, 0xf5, 0xfa, 0x2c, 0x7d, 0x1d, 0x41, 0x57, 0x9e, 0xeb, 0x3b, 0x32, 0xbf, 0x15, 0x0f,
	0xd1, 0xd8, 0x3a, 0xb4, 0xf5, 0xbb, 0xf7, 0xec, 0xb1, 0xfd, 0x98, 0x17, 0xe5, 0x81, 0xc7, 0xbc,
	0x44, 0xf5, 0xde, 0x35, 0x78, 0xb2, 0xe7, 0x82, 0x8e, 0xa1, 0xad, 0xb6, 0x77, 0x21, 0xdd, 0x66,
	0xaf, 0xf0, 0x07, 0x07, 0x5d, 0x7d, 0x94, 0x9d, 0x7b, 0x12, 0x36, 0xb3, 0xf7, 0xec, 0x8e, 0x6d,
	0xe9, 0x1e, 0x24, 0x97, 0xb8, 0x7b, 0x0e, 0xd2, 0x88, 0xcc, 0x66, 0x22, 0xd7, 0xf1, 0x54, 0x8f,
	0x08, 0x29, 0x58, 0xad, 0xb3, 0xb2, 0xfe, 0xe3, 0xd4, 0x75, 0x4e, 0x4e, 0x5d, 0xe7, 0xf7, 0xa9,
	0xeb, 0x7c, 0xea, 0xb8, 0x95, 0xef, 0x1d, 0xd7, 0x39, 0xe9, 0xb8, 0x95, 0x5f, 0x1d, 0xb7, 0xb2,
	0xb5, 0x40, 0x99, 0x6a, 0x24, 0xa1, 0x1f, 0xf1, 0x3d, 0xf3, 0xd2, 0xc6, 0xa0, 0x5e, 0x73, 0xb1,
	0x6b, 0xbe, 0xd5, 0x22, 0x2e, 0x20, 0x68, 0x67, 0x01, 0x2a, 0x6d, 0x81, 0x0c, 0x47, 0xb3, 0xff,
	0xa0, 0xe5, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xbe, 0x7d, 0x38, 0x4e, 0x07, 0x00, 0x00,
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
	RegisterExternalKeys(ctx context.Context, in *RegisterExternalKeysRequest, opts ...grpc.CallOption) (*RegisterExternalKeysResponse, error)
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error)
	StartKeygen(ctx context.Context, in *StartKeygenRequest, opts ...grpc.CallOption) (*StartKeygenResponse, error)
	ProcessKeygenTraffic(ctx context.Context, in *ProcessKeygenTrafficRequest, opts ...grpc.CallOption) (*ProcessKeygenTrafficResponse, error)
	RotateKey(ctx context.Context, in *RotateKeyRequest, opts ...grpc.CallOption) (*RotateKeyResponse, error)
	VotePubKey(ctx context.Context, in *VotePubKeyRequest, opts ...grpc.CallOption) (*VotePubKeyResponse, error)
	ProcessSignTraffic(ctx context.Context, in *ProcessSignTrafficRequest, opts ...grpc.CallOption) (*ProcessSignTrafficResponse, error)
	VoteSig(ctx context.Context, in *VoteSigRequest, opts ...grpc.CallOption) (*VoteSigResponse, error)
	SubmitMultisigPubKeys(ctx context.Context, in *SubmitMultisigPubKeysRequest, opts ...grpc.CallOption) (*SubmitMultisigPubKeysResponse, error)
	SubmitMultisigSignatures(ctx context.Context, in *SubmitMultisigSignaturesRequest, opts ...grpc.CallOption) (*SubmitMultisigSignaturesResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) RegisterExternalKeys(ctx context.Context, in *RegisterExternalKeysRequest, opts ...grpc.CallOption) (*RegisterExternalKeysResponse, error) {
	out := new(RegisterExternalKeysResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/RegisterExternalKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error) {
	out := new(HeartBeatResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) StartKeygen(ctx context.Context, in *StartKeygenRequest, opts ...grpc.CallOption) (*StartKeygenResponse, error) {
	out := new(StartKeygenResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/StartKeygen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ProcessKeygenTraffic(ctx context.Context, in *ProcessKeygenTrafficRequest, opts ...grpc.CallOption) (*ProcessKeygenTrafficResponse, error) {
	out := new(ProcessKeygenTrafficResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/ProcessKeygenTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RotateKey(ctx context.Context, in *RotateKeyRequest, opts ...grpc.CallOption) (*RotateKeyResponse, error) {
	out := new(RotateKeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/RotateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VotePubKey(ctx context.Context, in *VotePubKeyRequest, opts ...grpc.CallOption) (*VotePubKeyResponse, error) {
	out := new(VotePubKeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/VotePubKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ProcessSignTraffic(ctx context.Context, in *ProcessSignTrafficRequest, opts ...grpc.CallOption) (*ProcessSignTrafficResponse, error) {
	out := new(ProcessSignTrafficResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/ProcessSignTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VoteSig(ctx context.Context, in *VoteSigRequest, opts ...grpc.CallOption) (*VoteSigResponse, error) {
	out := new(VoteSigResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/VoteSig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SubmitMultisigPubKeys(ctx context.Context, in *SubmitMultisigPubKeysRequest, opts ...grpc.CallOption) (*SubmitMultisigPubKeysResponse, error) {
	out := new(SubmitMultisigPubKeysResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/SubmitMultisigPubKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SubmitMultisigSignatures(ctx context.Context, in *SubmitMultisigSignaturesRequest, opts ...grpc.CallOption) (*SubmitMultisigSignaturesResponse, error) {
	out := new(SubmitMultisigSignaturesResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/SubmitMultisigSignatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	RegisterExternalKeys(context.Context, *RegisterExternalKeysRequest) (*RegisterExternalKeysResponse, error)
	HeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatResponse, error)
	StartKeygen(context.Context, *StartKeygenRequest) (*StartKeygenResponse, error)
	ProcessKeygenTraffic(context.Context, *ProcessKeygenTrafficRequest) (*ProcessKeygenTrafficResponse, error)
	RotateKey(context.Context, *RotateKeyRequest) (*RotateKeyResponse, error)
	VotePubKey(context.Context, *VotePubKeyRequest) (*VotePubKeyResponse, error)
	ProcessSignTraffic(context.Context, *ProcessSignTrafficRequest) (*ProcessSignTrafficResponse, error)
	VoteSig(context.Context, *VoteSigRequest) (*VoteSigResponse, error)
	SubmitMultisigPubKeys(context.Context, *SubmitMultisigPubKeysRequest) (*SubmitMultisigPubKeysResponse, error)
	SubmitMultisigSignatures(context.Context, *SubmitMultisigSignaturesRequest) (*SubmitMultisigSignaturesResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) RegisterExternalKeys(ctx context.Context, req *RegisterExternalKeysRequest) (*RegisterExternalKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterExternalKeys not implemented")
}
func (*UnimplementedMsgServiceServer) HeartBeat(ctx context.Context, req *HeartBeatRequest) (*HeartBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (*UnimplementedMsgServiceServer) StartKeygen(ctx context.Context, req *StartKeygenRequest) (*StartKeygenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartKeygen not implemented")
}
func (*UnimplementedMsgServiceServer) ProcessKeygenTraffic(ctx context.Context, req *ProcessKeygenTrafficRequest) (*ProcessKeygenTrafficResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessKeygenTraffic not implemented")
}
func (*UnimplementedMsgServiceServer) RotateKey(ctx context.Context, req *RotateKeyRequest) (*RotateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateKey not implemented")
}
func (*UnimplementedMsgServiceServer) VotePubKey(ctx context.Context, req *VotePubKeyRequest) (*VotePubKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VotePubKey not implemented")
}
func (*UnimplementedMsgServiceServer) ProcessSignTraffic(ctx context.Context, req *ProcessSignTrafficRequest) (*ProcessSignTrafficResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessSignTraffic not implemented")
}
func (*UnimplementedMsgServiceServer) VoteSig(ctx context.Context, req *VoteSigRequest) (*VoteSigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteSig not implemented")
}
func (*UnimplementedMsgServiceServer) SubmitMultisigPubKeys(ctx context.Context, req *SubmitMultisigPubKeysRequest) (*SubmitMultisigPubKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitMultisigPubKeys not implemented")
}
func (*UnimplementedMsgServiceServer) SubmitMultisigSignatures(ctx context.Context, req *SubmitMultisigSignaturesRequest) (*SubmitMultisigSignaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitMultisigSignatures not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_RegisterExternalKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterExternalKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterExternalKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/RegisterExternalKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterExternalKeys(ctx, req.(*RegisterExternalKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_StartKeygen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartKeygenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).StartKeygen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/StartKeygen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).StartKeygen(ctx, req.(*StartKeygenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ProcessKeygenTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessKeygenTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ProcessKeygenTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/ProcessKeygenTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ProcessKeygenTraffic(ctx, req.(*ProcessKeygenTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RotateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RotateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/RotateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RotateKey(ctx, req.(*RotateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VotePubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VotePubKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VotePubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/VotePubKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VotePubKey(ctx, req.(*VotePubKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ProcessSignTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessSignTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ProcessSignTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/ProcessSignTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ProcessSignTraffic(ctx, req.(*ProcessSignTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VoteSig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteSigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VoteSig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/VoteSig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VoteSig(ctx, req.(*VoteSigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SubmitMultisigPubKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitMultisigPubKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SubmitMultisigPubKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/SubmitMultisigPubKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SubmitMultisigPubKeys(ctx, req.(*SubmitMultisigPubKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SubmitMultisigSignatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitMultisigSignaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SubmitMultisigSignatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/SubmitMultisigSignatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SubmitMultisigSignatures(ctx, req.(*SubmitMultisigSignaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.tss.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterExternalKeys",
			Handler:    _MsgService_RegisterExternalKeys_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _MsgService_HeartBeat_Handler,
		},
		{
			MethodName: "StartKeygen",
			Handler:    _MsgService_StartKeygen_Handler,
		},
		{
			MethodName: "ProcessKeygenTraffic",
			Handler:    _MsgService_ProcessKeygenTraffic_Handler,
		},
		{
			MethodName: "RotateKey",
			Handler:    _MsgService_RotateKey_Handler,
		},
		{
			MethodName: "VotePubKey",
			Handler:    _MsgService_VotePubKey_Handler,
		},
		{
			MethodName: "ProcessSignTraffic",
			Handler:    _MsgService_ProcessSignTraffic_Handler,
		},
		{
			MethodName: "VoteSig",
			Handler:    _MsgService_VoteSig_Handler,
		},
		{
			MethodName: "SubmitMultisigPubKeys",
			Handler:    _MsgService_SubmitMultisigPubKeys_Handler,
		},
		{
			MethodName: "SubmitMultisigSignatures",
			Handler:    _MsgService_SubmitMultisigSignatures_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/tss/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// NextKeyID returns the key ID assigned for the next rotation on a given
	// chain and for the given key role
	NextKeyID(ctx context.Context, in *NextKeyIDRequest, opts ...grpc.CallOption) (*NextKeyIDResponse, error)
	// AssignableKey returns true if there is no assigned key for the
	// next rotation on a given chain, and false otherwise
	AssignableKey(ctx context.Context, in *AssignableKeyRequest, opts ...grpc.CallOption) (*AssignableKeyResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) NextKeyID(ctx context.Context, in *NextKeyIDRequest, opts ...grpc.CallOption) (*NextKeyIDResponse, error) {
	out := new(NextKeyIDResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.QueryService/NextKeyID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) AssignableKey(ctx context.Context, in *AssignableKeyRequest, opts ...grpc.CallOption) (*AssignableKeyResponse, error) {
	out := new(AssignableKeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.QueryService/AssignableKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// NextKeyID returns the key ID assigned for the next rotation on a given
	// chain and for the given key role
	NextKeyID(context.Context, *NextKeyIDRequest) (*NextKeyIDResponse, error)
	// AssignableKey returns true if there is no assigned key for the
	// next rotation on a given chain, and false otherwise
	AssignableKey(context.Context, *AssignableKeyRequest) (*AssignableKeyResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) NextKeyID(ctx context.Context, req *NextKeyIDRequest) (*NextKeyIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextKeyID not implemented")
}
func (*UnimplementedQueryServiceServer) AssignableKey(ctx context.Context, req *AssignableKeyRequest) (*AssignableKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignableKey not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_NextKeyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextKeyIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).NextKeyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.QueryService/NextKeyID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).NextKeyID(ctx, req.(*NextKeyIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_AssignableKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignableKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).AssignableKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.QueryService/AssignableKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).AssignableKey(ctx, req.(*AssignableKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.tss.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextKeyID",
			Handler:    _QueryService_NextKeyID_Handler,
		},
		{
			MethodName: "AssignableKey",
			Handler:    _QueryService_AssignableKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/tss/v1beta1/service.proto",
}
