// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/v1beta1/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// NetworkInfo describes information about a network
type NetworkInfo struct {
	Name string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=id,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"id"`
}

func (m *NetworkInfo) Reset()         { *m = NetworkInfo{} }
func (m *NetworkInfo) String() string { return proto.CompactTextString(m) }
func (*NetworkInfo) ProtoMessage()    {}
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{0}
}
func (m *NetworkInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfo.Merge(m, src)
}
func (m *NetworkInfo) XXX_Size() int {
	return m.Size()
}
func (m *NetworkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfo proto.InternalMessageInfo

// BurnerInfo describes information required to burn token at an burner address
// that is deposited by an user
type BurnerInfo struct {
	TokenAddress     Address `protobuf:"bytes,1,opt,name=token_address,json=tokenAddress,proto3,customtype=Address" json:"token_address"`
	DestinationChain string  `protobuf:"bytes,2,opt,name=destination_chain,json=destinationChain,proto3" json:"destination_chain,omitempty"`
	Symbol           string  `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Salt             Hash    `protobuf:"bytes,4,opt,name=salt,proto3,customtype=Hash" json:"salt"`
}

func (m *BurnerInfo) Reset()         { *m = BurnerInfo{} }
func (m *BurnerInfo) String() string { return proto.CompactTextString(m) }
func (*BurnerInfo) ProtoMessage()    {}
func (*BurnerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{1}
}
func (m *BurnerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurnerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurnerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurnerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnerInfo.Merge(m, src)
}
func (m *BurnerInfo) XXX_Size() int {
	return m.Size()
}
func (m *BurnerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BurnerInfo proto.InternalMessageInfo

// ERC20Deposit contains information for an ERC20 deposit
type ERC20Deposit struct {
	TxID             Hash                                    `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,customtype=Hash" json:"tx_id"`
	Amount           github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"amount"`
	DestinationChain string                                  `protobuf:"bytes,3,opt,name=destination_chain,json=destinationChain,proto3" json:"destination_chain,omitempty"`
	BurnerAddress    Address                                 `protobuf:"bytes,4,opt,name=burner_address,json=burnerAddress,proto3,customtype=Address" json:"burner_address"`
}

func (m *ERC20Deposit) Reset()         { *m = ERC20Deposit{} }
func (m *ERC20Deposit) String() string { return proto.CompactTextString(m) }
func (*ERC20Deposit) ProtoMessage()    {}
func (*ERC20Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{2}
}
func (m *ERC20Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ERC20Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ERC20Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ERC20Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20Deposit.Merge(m, src)
}
func (m *ERC20Deposit) XXX_Size() int {
	return m.Size()
}
func (m *ERC20Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20Deposit proto.InternalMessageInfo

// ERC20TokenDeployment describes information about an ERC20 token
type ERC20TokenDeployment struct {
	Asset        string  `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	TokenAddress Address `protobuf:"bytes,2,opt,name=token_address,json=tokenAddress,proto3,customtype=Address" json:"token_address"`
}

func (m *ERC20TokenDeployment) Reset()         { *m = ERC20TokenDeployment{} }
func (m *ERC20TokenDeployment) String() string { return proto.CompactTextString(m) }
func (*ERC20TokenDeployment) ProtoMessage()    {}
func (*ERC20TokenDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{3}
}
func (m *ERC20TokenDeployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ERC20TokenDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ERC20TokenDeployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ERC20TokenDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20TokenDeployment.Merge(m, src)
}
func (m *ERC20TokenDeployment) XXX_Size() int {
	return m.Size()
}
func (m *ERC20TokenDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20TokenDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20TokenDeployment proto.InternalMessageInfo

// TransferOwnership contains information for a transfer ownership
type TransferOwnership struct {
	TxID      Hash   `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,customtype=Hash" json:"tx_id"`
	NextKeyID string `protobuf:"bytes,3,opt,name=next_key_id,json=nextKeyId,proto3" json:"next_key_id,omitempty"`
}

func (m *TransferOwnership) Reset()         { *m = TransferOwnership{} }
func (m *TransferOwnership) String() string { return proto.CompactTextString(m) }
func (*TransferOwnership) ProtoMessage()    {}
func (*TransferOwnership) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{4}
}
func (m *TransferOwnership) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferOwnership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferOwnership.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferOwnership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferOwnership.Merge(m, src)
}
func (m *TransferOwnership) XXX_Size() int {
	return m.Size()
}
func (m *TransferOwnership) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferOwnership.DiscardUnknown(m)
}

var xxx_messageInfo_TransferOwnership proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NetworkInfo)(nil), "evm.v1beta1.NetworkInfo")
	proto.RegisterType((*BurnerInfo)(nil), "evm.v1beta1.BurnerInfo")
	proto.RegisterType((*ERC20Deposit)(nil), "evm.v1beta1.ERC20Deposit")
	proto.RegisterType((*ERC20TokenDeployment)(nil), "evm.v1beta1.ERC20TokenDeployment")
	proto.RegisterType((*TransferOwnership)(nil), "evm.v1beta1.TransferOwnership")
}

func init() { proto.RegisterFile("evm/v1beta1/types.proto", fileDescriptor_af2cf809b4baed32) }

var fileDescriptor_af2cf809b4baed32 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x8f, 0xd2, 0x40,
	0x14, 0xa6, 0x6c, 0x17, 0xc3, 0x00, 0xea, 0x4e, 0x88, 0x12, 0x0f, 0x85, 0x70, 0xd0, 0x35, 0x06,
	0xba, 0xab, 0xc6, 0xa3, 0x89, 0x5d, 0x8c, 0x36, 0x26, 0x98, 0x34, 0x78, 0xf1, 0x42, 0xa6, 0xf4,
	0x2d, 0x4c, 0xa0, 0x33, 0xcd, 0xcc, 0xc0, 0xb6, 0xff, 0xc2, 0xbf, 0xe1, 0x3f, 0xe1, 0xb8, 0x47,
	0xe3, 0x81, 0x68, 0x39, 0xf8, 0x37, 0x4c, 0xa7, 0x13, 0x43, 0xe2, 0x1a, 0xf5, 0xd4, 0xf7, 0xbd,
	0xf7, 0xbd, 0xbe, 0xf7, 0xbe, 0x37, 0x0f, 0xdd, 0x87, 0x4d, 0xec, 0x6e, 0xce, 0x43, 0x50, 0xe4,
	0xdc, 0x55, 0x59, 0x02, 0x72, 0x98, 0x08, 0xae, 0x38, 0x6e, 0xc0, 0x26, 0x1e, 0x9a, 0xc0, 0x83,
	0xf6, 0x9c, 0xcf, 0xb9, 0xf6, 0xbb, 0x85, 0x55, 0x52, 0xfa, 0x04, 0x35, 0xc6, 0xa0, 0xae, 0xb8,
	0x58, 0xfa, 0xec, 0x92, 0x63, 0x8c, 0x6c, 0x46, 0x62, 0xe8, 0x58, 0x3d, 0xeb, 0xb4, 0x1e, 0x68,
	0x1b, 0xbf, 0x44, 0x55, 0x1a, 0x75, 0xaa, 0x3d, 0xeb, 0xb4, 0xe9, 0x0d, 0xb7, 0xbb, 0x6e, 0xe5,
	0xeb, 0xae, 0xfb, 0x70, 0x4e, 0xd5, 0x62, 0x1d, 0x0e, 0x67, 0x3c, 0x76, 0x67, 0x5c, 0xc6, 0x5c,
	0x9a, 0xcf, 0x40, 0x46, 0x4b, 0xd3, 0x83, 0xcf, 0x54, 0x50, 0xa5, 0x51, 0xff, 0xb3, 0x85, 0x90,
	0xb7, 0x16, 0x0c, 0x84, 0x2e, 0xf1, 0x1c, 0xb5, 0x14, 0x5f, 0x02, 0x9b, 0x92, 0x28, 0x12, 0x20,
	0xa5, 0xae, 0xd5, 0xf4, 0xee, 0x98, 0x3f, 0xdf, 0x7a, 0x55, 0xba, 0x83, 0xa6, 0x66, 0x19, 0x84,
	0x9f, 0xa0, 0x93, 0x08, 0xa4, 0xa2, 0x8c, 0x28, 0xca, 0xd9, 0x74, 0xb6, 0x20, 0x94, 0xe9, 0x9e,
	0xea, 0xc1, 0xdd, 0x83, 0xc0, 0x45, 0xe1, 0xc7, 0xf7, 0x50, 0x4d, 0x66, 0x71, 0xc8, 0x57, 0x9d,
	0x23, 0xcd, 0x30, 0x08, 0xf7, 0x90, 0x2d, 0xc9, 0x4a, 0x75, 0x6c, 0x5d, 0xb1, 0x69, 0x2a, 0xda,
	0x6f, 0x89, 0x5c, 0x04, 0x3a, 0xd2, 0xff, 0x61, 0xa1, 0xe6, 0xeb, 0xe0, 0xe2, 0xe9, 0xd9, 0x08,
	0x12, 0x2e, 0xa9, 0xc2, 0x8f, 0xd1, 0xb1, 0x4a, 0xa7, 0x34, 0x32, 0x5d, 0xb6, 0x0f, 0x73, 0xf2,
	0x5d, 0xd7, 0x9e, 0xa4, 0xfe, 0x28, 0xb0, 0x55, 0xea, 0x47, 0xf8, 0x0d, 0xaa, 0x91, 0x98, 0xaf,
	0x99, 0x32, 0x5a, 0xb9, 0x86, 0xfb, 0xe8, 0x1f, 0xb4, 0xfa, 0x40, 0x99, 0x0a, 0x4c, 0xfa, 0xcd,
	0xb3, 0x1e, 0xfd, 0x61, 0xd6, 0x17, 0xe8, 0x76, 0xa8, 0xc5, 0xfd, 0xa5, 0xa7, 0x7d, 0xb3, 0x9e,
	0xad, 0x92, 0x66, 0x60, 0x3f, 0x44, 0x6d, 0x3d, 0xe8, 0xa4, 0x50, 0x79, 0x04, 0xc9, 0x8a, 0x67,
	0x31, 0x30, 0x85, 0xdb, 0xe8, 0x98, 0x48, 0x09, 0xca, 0x3c, 0x81, 0x12, 0xfc, 0xbe, 0x34, 0x2d,
	0xfd, 0x5f, 0x96, 0xd6, 0x8f, 0xd1, 0xc9, 0x44, 0x10, 0x26, 0x2f, 0x41, 0xbc, 0xbf, 0x62, 0x20,
	0xe4, 0x82, 0x26, 0xff, 0xa3, 0xe8, 0x00, 0x35, 0x18, 0xa4, 0x6a, 0xba, 0x84, 0xac, 0x48, 0xd0,
	0x12, 0x78, 0xad, 0x7c, 0xd7, 0xad, 0x8f, 0x21, 0x55, 0xef, 0x20, 0xf3, 0x47, 0x41, 0x9d, 0x19,
	0x33, 0xf2, 0xc6, 0xdb, 0xef, 0x4e, 0x65, 0x9b, 0x3b, 0xd6, 0x75, 0xee, 0x58, 0xdf, 0x72, 0xc7,
	0xfa, 0xb4, 0x77, 0x2a, 0xd7, 0x7b, 0xa7, 0xf2, 0x65, 0xef, 0x54, 0x3e, 0x9e, 0x1d, 0xac, 0x81,
	0xa4, 0xb0, 0x22, 0x82, 0x95, 0x0f, 0xdf, 0xa0, 0xc1, 0x8c, 0x0b, 0x70, 0x53, 0xb7, 0x38, 0x26,
	0xbd, 0x94, 0xb0, 0xa6, 0x4f, 0xe4, 0xd9, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xae, 0xc4,
	0x0c, 0x60, 0x03, 0x00, 0x00,
}

func (m *NetworkInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Id.Size()
		i -= size
		if _, err := m.Id.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BurnerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurnerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurnerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Salt.Size()
		i -= size
		if _, err := m.Salt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestinationChain) > 0 {
		i -= len(m.DestinationChain)
		copy(dAtA[i:], m.DestinationChain)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DestinationChain)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.TokenAddress.Size()
		i -= size
		if _, err := m.TokenAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ERC20Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ERC20Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ERC20Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BurnerAddress.Size()
		i -= size
		if _, err := m.BurnerAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.DestinationChain) > 0 {
		i -= len(m.DestinationChain)
		copy(dAtA[i:], m.DestinationChain)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DestinationChain)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.TxID.Size()
		i -= size
		if _, err := m.TxID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ERC20TokenDeployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ERC20TokenDeployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ERC20TokenDeployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenAddress.Size()
		i -= size
		if _, err := m.TokenAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TransferOwnership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferOwnership) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferOwnership) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextKeyID) > 0 {
		i -= len(m.NextKeyID)
		copy(dAtA[i:], m.NextKeyID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NextKeyID)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.TxID.Size()
		i -= size
		if _, err := m.TxID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NetworkInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Id.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *BurnerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenAddress.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.DestinationChain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Salt.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *ERC20Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TxID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.DestinationChain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.BurnerAddress.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *ERC20TokenDeployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.TokenAddress.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *TransferOwnership) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TxID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.NextKeyID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BurnerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BurnerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurnerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Salt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ERC20Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ERC20Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ERC20Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnerAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ERC20TokenDeployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ERC20TokenDeployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ERC20TokenDeployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TransferOwnership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TransferOwnership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferOwnership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
