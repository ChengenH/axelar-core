// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/multisig/v1beta1/events.proto

package types

import (
	fmt "fmt"
	github_com_axelarnetwork_axelar_core_x_multisig_exported "github.com/axelarnetwork/axelar-core/x/multisig/exported"
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

type KeygenStarted struct {
	Module       string                                                         `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	KeyID        github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3,casttype=github.com/axelarnetwork/axelar-core/x/multisig/exported.KeyID" json:"key_id,omitempty"`
	Participants []github_com_cosmos_cosmos_sdk_types.ValAddress                `protobuf:"bytes,3,rep,name=participants,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participants,omitempty"`
}

func (m *KeygenStarted) Reset()         { *m = KeygenStarted{} }
func (m *KeygenStarted) String() string { return proto.CompactTextString(m) }
func (*KeygenStarted) ProtoMessage()    {}
func (*KeygenStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b18b0391cba3fc, []int{0}
}
func (m *KeygenStarted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeygenStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeygenStarted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeygenStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeygenStarted.Merge(m, src)
}
func (m *KeygenStarted) XXX_Size() int {
	return m.Size()
}
func (m *KeygenStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_KeygenStarted.DiscardUnknown(m)
}

var xxx_messageInfo_KeygenStarted proto.InternalMessageInfo

func (m *KeygenStarted) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *KeygenStarted) GetKeyID() github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *KeygenStarted) GetParticipants() []github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Participants
	}
	return nil
}

type KeygenCompleted struct {
	Module string                                                         `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	KeyID  github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3,casttype=github.com/axelarnetwork/axelar-core/x/multisig/exported.KeyID" json:"key_id,omitempty"`
}

func (m *KeygenCompleted) Reset()         { *m = KeygenCompleted{} }
func (m *KeygenCompleted) String() string { return proto.CompactTextString(m) }
func (*KeygenCompleted) ProtoMessage()    {}
func (*KeygenCompleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b18b0391cba3fc, []int{1}
}
func (m *KeygenCompleted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeygenCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeygenCompleted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeygenCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeygenCompleted.Merge(m, src)
}
func (m *KeygenCompleted) XXX_Size() int {
	return m.Size()
}
func (m *KeygenCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_KeygenCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_KeygenCompleted proto.InternalMessageInfo

func (m *KeygenCompleted) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *KeygenCompleted) GetKeyID() github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID {
	if m != nil {
		return m.KeyID
	}
	return ""
}

type KeygenExpired struct {
	Module string                                                         `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	KeyID  github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3,casttype=github.com/axelarnetwork/axelar-core/x/multisig/exported.KeyID" json:"key_id,omitempty"`
}

func (m *KeygenExpired) Reset()         { *m = KeygenExpired{} }
func (m *KeygenExpired) String() string { return proto.CompactTextString(m) }
func (*KeygenExpired) ProtoMessage()    {}
func (*KeygenExpired) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b18b0391cba3fc, []int{2}
}
func (m *KeygenExpired) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeygenExpired) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeygenExpired.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeygenExpired) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeygenExpired.Merge(m, src)
}
func (m *KeygenExpired) XXX_Size() int {
	return m.Size()
}
func (m *KeygenExpired) XXX_DiscardUnknown() {
	xxx_messageInfo_KeygenExpired.DiscardUnknown(m)
}

var xxx_messageInfo_KeygenExpired proto.InternalMessageInfo

func (m *KeygenExpired) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *KeygenExpired) GetKeyID() github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID {
	if m != nil {
		return m.KeyID
	}
	return ""
}

type PubKeySubmitted struct {
	Module      string                                                         `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	KeyID       github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3,casttype=github.com/axelarnetwork/axelar-core/x/multisig/exported.KeyID" json:"key_id,omitempty"`
	Participant github_com_cosmos_cosmos_sdk_types.ValAddress                  `protobuf:"bytes,3,opt,name=participant,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participant,omitempty"`
	PubKey      PublicKey                                                      `protobuf:"bytes,4,opt,name=pub_key,json=pubKey,proto3,casttype=PublicKey" json:"pub_key,omitempty"`
}

func (m *PubKeySubmitted) Reset()         { *m = PubKeySubmitted{} }
func (m *PubKeySubmitted) String() string { return proto.CompactTextString(m) }
func (*PubKeySubmitted) ProtoMessage()    {}
func (*PubKeySubmitted) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b18b0391cba3fc, []int{3}
}
func (m *PubKeySubmitted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKeySubmitted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKeySubmitted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKeySubmitted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeySubmitted.Merge(m, src)
}
func (m *PubKeySubmitted) XXX_Size() int {
	return m.Size()
}
func (m *PubKeySubmitted) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeySubmitted.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeySubmitted proto.InternalMessageInfo

func (m *PubKeySubmitted) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *PubKeySubmitted) GetKeyID() github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *PubKeySubmitted) GetParticipant() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Participant
	}
	return nil
}

func (m *PubKeySubmitted) GetPubKey() PublicKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func init() {
	proto.RegisterType((*KeygenStarted)(nil), "axelar.multisig.v1beta1.KeygenStarted")
	proto.RegisterType((*KeygenCompleted)(nil), "axelar.multisig.v1beta1.KeygenCompleted")
	proto.RegisterType((*KeygenExpired)(nil), "axelar.multisig.v1beta1.KeygenExpired")
	proto.RegisterType((*PubKeySubmitted)(nil), "axelar.multisig.v1beta1.PubKeySubmitted")
}

func init() {
	proto.RegisterFile("axelar/multisig/v1beta1/events.proto", fileDescriptor_36b18b0391cba3fc)
}

var fileDescriptor_36b18b0391cba3fc = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0xb7, 0xde, 0xc8, 0x1d, 0x6f, 0x29, 0x04, 0xd1, 0xe0, 0x22, 0x29, 0x45, 0xa4,
	0x9b, 0x26, 0x14, 0x71, 0x2b, 0x58, 0x75, 0x51, 0xb2, 0xb0, 0xa4, 0xe8, 0xc2, 0x4d, 0xc9, 0x24,
	0x87, 0x38, 0xe4, 0xcf, 0x0c, 0x33, 0x93, 0x9a, 0x79, 0x05, 0x45, 0xf0, 0xb1, 0x5c, 0x76, 0xe9,
	0x2a, 0x48, 0xfa, 0x16, 0x5d, 0x49, 0x93, 0xa8, 0x75, 0xa9, 0x9b, 0xae, 0x92, 0x33, 0xfc, 0xf8,
	0xce, 0xf7, 0x9d, 0x39, 0x83, 0x1f, 0x87, 0x15, 0x64, 0xa1, 0xf0, 0xf2, 0x32, 0x53, 0x54, 0xd2,
	0xc4, 0xdb, 0x2d, 0x08, 0xa8, 0x70, 0xe1, 0xc1, 0x0e, 0x0a, 0x25, 0x5d, 0x2e, 0x98, 0x62, 0xe6,
	0xc3, 0x8e, 0x72, 0x7f, 0x51, 0x6e, 0x4f, 0x3d, 0xba, 0x9f, 0xb0, 0x84, 0xb5, 0x8c, 0x77, 0xfa,
	0xeb, 0xf0, 0x69, 0x83, 0xf0, 0xc8, 0x07, 0x9d, 0x40, 0xb1, 0x51, 0xa1, 0x50, 0x10, 0x9b, 0x0f,
	0xb0, 0x91, 0xb3, 0xb8, 0xcc, 0xc0, 0x42, 0x13, 0x34, 0xbb, 0x09, 0xfa, 0xca, 0x24, 0xd8, 0x48,
	0x41, 0x6f, 0x69, 0x6c, 0x5d, 0x9d, 0xce, 0x97, 0x7e, 0x53, 0x3b, 0xd7, 0x3e, 0xe8, 0xd5, 0xab,
	0x63, 0xed, 0x3c, 0x4f, 0xa8, 0xfa, 0x50, 0x12, 0x37, 0x62, 0xb9, 0xd7, 0x19, 0x28, 0x40, 0x7d,
	0x64, 0x22, 0xed, 0xab, 0x79, 0xc4, 0x04, 0x78, 0xd5, 0x1f, 0xef, 0x50, 0x71, 0x76, 0x6a, 0xe7,
	0xb6, 0x0a, 0xc1, 0x75, 0x0a, 0x7a, 0x15, 0x9b, 0x6f, 0xf1, 0x2d, 0x0f, 0x85, 0xa2, 0x11, 0xe5,
	0x61, 0xa1, 0xa4, 0x35, 0x9c, 0x0c, 0x67, 0xb7, 0xcb, 0xc5, 0xb1, 0x76, 0xe6, 0x67, 0x0d, 0x22,
	0x26, 0x73, 0x26, 0xfb, 0xcf, 0x5c, 0xc6, 0xa9, 0xa7, 0x34, 0x07, 0xe9, 0xbe, 0x0b, 0xb3, 0x17,
	0x71, 0x2c, 0x40, 0xca, 0xe0, 0x2f, 0x99, 0xe9, 0x17, 0x84, 0xc7, 0x5d, 0xc8, 0x97, 0x2c, 0xe7,
	0x19, 0x5c, 0x38, 0xe6, 0xf4, 0xf3, 0xef, 0xa1, 0xbf, 0xae, 0x38, 0x15, 0x17, 0x76, 0xf3, 0xe9,
	0x0a, 0x8f, 0xd7, 0x25, 0xf1, 0x41, 0x6f, 0x4a, 0x92, 0x53, 0x75, 0xe9, 0x25, 0xd8, 0xe0, 0x7b,
	0x67, 0xb7, 0x67, 0x0d, 0x27, 0xe8, 0xff, 0x76, 0xe0, 0x5c, 0xc5, 0x7c, 0x82, 0xef, 0xf2, 0x92,
	0x6c, 0x53, 0xd0, 0xd6, 0x9d, 0x56, 0x70, 0x74, 0xac, 0x9d, 0x9b, 0x75, 0x49, 0x32, 0x1a, 0xf9,
	0xa0, 0x03, 0x83, 0xb7, 0x13, 0x58, 0xbe, 0xf9, 0xd6, 0xd8, 0x68, 0xdf, 0xd8, 0xe8, 0x47, 0x63,
	0xa3, 0xaf, 0x07, 0x7b, 0xb0, 0x3f, 0xd8, 0x83, 0xef, 0x07, 0x7b, 0xf0, 0xfe, 0xd9, 0xbf, 0xa6,
	0x6b, 0x0d, 0x11, 0xa3, 0x7d, 0x67, 0x4f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xfd, 0x99,
	0x18, 0xbe, 0x03, 0x00, 0x00,
}

func (m *KeygenStarted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeygenStarted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenStarted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Participants[iNdEx])
			copy(dAtA[i:], m.Participants[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Participants[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeygenCompleted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeygenCompleted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenCompleted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeygenExpired) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeygenExpired) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenExpired) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PubKeySubmitted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKeySubmitted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKeySubmitted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeygenStarted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Participants) > 0 {
		for _, b := range m.Participants {
			l = len(b)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *KeygenCompleted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *KeygenExpired) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *PubKeySubmitted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeygenStarted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: KeygenStarted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeygenStarted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, make([]byte, postIndex-iNdEx))
			copy(m.Participants[len(m.Participants)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *KeygenCompleted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: KeygenCompleted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeygenCompleted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *KeygenExpired) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: KeygenExpired: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeygenExpired: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *PubKeySubmitted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: PubKeySubmitted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKeySubmitted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participant = append(m.Participant[:0], dAtA[iNdEx:postIndex]...)
			if m.Participant == nil {
				m.Participant = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)