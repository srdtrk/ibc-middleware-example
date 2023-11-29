// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmosregistry/example/v1/types.proto

package example

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters of the module.
type Params struct {
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d1cf6348eb2ada, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// Metadata defines the middleware specific metadata encoded into the channel
// version bytestring See ICS004:
// https://github.com/cosmos/ibc/tree/master/spec/core/ics-004-channel-and-packet-semantics#Versioning
// If the middleware is a stateless middleware, you need not include any
// metadata, and can omit this message.
type Metadata struct {
	// example_version defines the middleware version
	ExampleVersion string `protobuf:"bytes,1,opt,name=example_version,json=exampleVersion,proto3" json:"example_version,omitempty"`
	// app_version defines the underlying application version, which may or may
	// not be a JSON encoded bytestring
	AppVersion string `protobuf:"bytes,2,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d1cf6348eb2ada, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetExampleVersion() string {
	if m != nil {
		return m.ExampleVersion
	}
	return ""
}

func (m *Metadata) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

// Counter defines a counter object.
// It is used only for genesis purposes. Collections does not need to use it.
type CallbackCounter struct {
	// on_recv_packet counts the number of times OnRecvPacket is called in the
	// underlying channel.
	OnRecvPacket uint64 `protobuf:"varint,1,opt,name=on_recv_packet,json=onRecvPacket,proto3" json:"on_recv_packet,omitempty"`
	// on_acknowledgement_packet counts the number of times
	// OnAcknowledgementPacket is called in the underlying channel.
	OnAcknowledgementPacket uint64 `protobuf:"varint,2,opt,name=on_acknowledgement_packet,json=onAcknowledgementPacket,proto3" json:"on_acknowledgement_packet,omitempty"`
	// on_timeout_packet counts the number of times OnTimeoutPacket is called in
	// the underlying channel.
	OnTimeoutPacket uint64 `protobuf:"varint,3,opt,name=on_timeout_packet,json=onTimeoutPacket,proto3" json:"on_timeout_packet,omitempty"`
	// send_packet counts the number of times SendPacket is called in the
	// underlying channel.
	SendPacket uint64 `protobuf:"varint,4,opt,name=send_packet,json=sendPacket,proto3" json:"send_packet,omitempty"`
	// channel_id defines the channel that is associated with the callback counts.
	ChannelId string `protobuf:"bytes,5,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (m *CallbackCounter) Reset()         { *m = CallbackCounter{} }
func (m *CallbackCounter) String() string { return proto.CompactTextString(m) }
func (*CallbackCounter) ProtoMessage()    {}
func (*CallbackCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d1cf6348eb2ada, []int{2}
}
func (m *CallbackCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CallbackCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CallbackCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CallbackCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackCounter.Merge(m, src)
}
func (m *CallbackCounter) XXX_Size() int {
	return m.Size()
}
func (m *CallbackCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackCounter.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackCounter proto.InternalMessageInfo

func (m *CallbackCounter) GetOnRecvPacket() uint64 {
	if m != nil {
		return m.OnRecvPacket
	}
	return 0
}

func (m *CallbackCounter) GetOnAcknowledgementPacket() uint64 {
	if m != nil {
		return m.OnAcknowledgementPacket
	}
	return 0
}

func (m *CallbackCounter) GetOnTimeoutPacket() uint64 {
	if m != nil {
		return m.OnTimeoutPacket
	}
	return 0
}

func (m *CallbackCounter) GetSendPacket() uint64 {
	if m != nil {
		return m.SendPacket
	}
	return 0
}

func (m *CallbackCounter) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// MiddlewareEnabledChannel defines a channel that is enabled with this
// middleware.
type MiddlewareEnabledChannel struct {
	// unique port identifier
	PortId string `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// unique channel identifier
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (m *MiddlewareEnabledChannel) Reset()         { *m = MiddlewareEnabledChannel{} }
func (m *MiddlewareEnabledChannel) String() string { return proto.CompactTextString(m) }
func (*MiddlewareEnabledChannel) ProtoMessage()    {}
func (*MiddlewareEnabledChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d1cf6348eb2ada, []int{3}
}
func (m *MiddlewareEnabledChannel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MiddlewareEnabledChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MiddlewareEnabledChannel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MiddlewareEnabledChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiddlewareEnabledChannel.Merge(m, src)
}
func (m *MiddlewareEnabledChannel) XXX_Size() int {
	return m.Size()
}
func (m *MiddlewareEnabledChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_MiddlewareEnabledChannel.DiscardUnknown(m)
}

var xxx_messageInfo_MiddlewareEnabledChannel proto.InternalMessageInfo

func (m *MiddlewareEnabledChannel) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *MiddlewareEnabledChannel) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	// counter defines the counter object.
	Counters []CallbackCounter `protobuf:"bytes,1,rep,name=counters,proto3" json:"counters"`
	// list of middleware enabled channels
	MiddlewareEnabledChannels []MiddlewareEnabledChannel `protobuf:"bytes,2,rep,name=middleware_enabled_channels,json=middlewareEnabledChannels,proto3" json:"middleware_enabled_channels"`
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d1cf6348eb2ada, []int{4}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCounters() []CallbackCounter {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *GenesisState) GetMiddlewareEnabledChannels() []MiddlewareEnabledChannel {
	if m != nil {
		return m.MiddlewareEnabledChannels
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*Params)(nil), "cosmosregistry.example.v1.Params")
	proto.RegisterType((*Metadata)(nil), "cosmosregistry.example.v1.Metadata")
	proto.RegisterType((*CallbackCounter)(nil), "cosmosregistry.example.v1.CallbackCounter")
	proto.RegisterType((*MiddlewareEnabledChannel)(nil), "cosmosregistry.example.v1.MiddlewareEnabledChannel")
	proto.RegisterType((*GenesisState)(nil), "cosmosregistry.example.v1.GenesisState")
}

func init() {
	proto.RegisterFile("cosmosregistry/example/v1/types.proto", fileDescriptor_b9d1cf6348eb2ada)
}

var fileDescriptor_b9d1cf6348eb2ada = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x3f, 0x4f, 0xdb, 0x40,
	0x18, 0xc6, 0x63, 0x43, 0x53, 0x78, 0x83, 0x88, 0xb0, 0x2a, 0x91, 0x50, 0xe1, 0x80, 0xd5, 0x3f,
	0x28, 0xad, 0x6c, 0x01, 0x9d, 0xd2, 0xa9, 0xa4, 0x55, 0xc5, 0x80, 0x44, 0x53, 0xd4, 0xa1, 0x8b,
	0x75, 0xb1, 0x5f, 0x05, 0x2b, 0xf6, 0xdd, 0xc9, 0x77, 0x09, 0x65, 0xe9, 0x5c, 0x55, 0x1d, 0xfa,
	0x2d, 0xda, 0x91, 0x8f, 0xc1, 0xc8, 0xd8, 0xa9, 0xaa, 0x92, 0x81, 0xaf, 0x51, 0xf9, 0xee, 0x88,
	0x20, 0x92, 0x59, 0x2c, 0xfb, 0xbd, 0xdf, 0xf3, 0x3e, 0xcf, 0xbd, 0xbe, 0x83, 0xa7, 0x11, 0x13,
	0x19, 0x13, 0x39, 0x0e, 0x12, 0x21, 0xf3, 0xf3, 0x00, 0xbf, 0x90, 0x8c, 0xa7, 0x18, 0x8c, 0x77,
	0x03, 0x79, 0xce, 0x51, 0xf8, 0x3c, 0x67, 0x92, 0x39, 0xcd, 0xbb, 0x98, 0x6f, 0x30, 0x7f, 0xbc,
	0xbb, 0xf1, 0x68, 0xc0, 0x06, 0x4c, 0x51, 0x41, 0xf1, 0xa6, 0x05, 0x1b, 0x6b, 0x24, 0x4b, 0x28,
	0x0b, 0xd4, 0x53, 0x97, 0xbc, 0x97, 0x50, 0x3d, 0x26, 0x39, 0xc9, 0x44, 0xc7, 0xfb, 0x7e, 0x7d,
	0xd1, 0xde, 0x2c, 0x71, 0xd6, 0x8c, 0x77, 0x02, 0x4b, 0x47, 0x28, 0x49, 0x4c, 0x24, 0x71, 0x9e,
	0x43, 0xdd, 0xac, 0x86, 0x63, 0xcc, 0x45, 0xc2, 0x68, 0xc3, 0xda, 0xb2, 0x76, 0x96, 0x7b, 0xab,
	0xa6, 0xfc, 0x49, 0x57, 0x9d, 0x16, 0xd4, 0x08, 0xe7, 0x33, 0xc8, 0x56, 0x10, 0x10, 0xce, 0x0d,
	0xe0, 0xfd, 0xb0, 0xa1, 0xde, 0x25, 0x69, 0xda, 0x27, 0xd1, 0xb0, 0xcb, 0x46, 0x54, 0x62, 0xee,
	0x3c, 0x81, 0x55, 0x46, 0xc3, 0x1c, 0xa3, 0x71, 0xc8, 0x49, 0x34, 0x44, 0xa9, 0x9a, 0x2f, 0xf6,
	0x56, 0x18, 0xed, 0x61, 0x34, 0x3e, 0x56, 0x35, 0xa7, 0x03, 0x4d, 0x46, 0x43, 0x12, 0x0d, 0x29,
	0x3b, 0x4b, 0x31, 0x1e, 0x60, 0x86, 0x54, 0xde, 0x08, 0x6c, 0x25, 0x58, 0x67, 0xf4, 0xcd, 0xdd,
	0x75, 0xa3, 0x6d, 0xc3, 0x1a, 0xa3, 0xa1, 0x4c, 0x32, 0x64, 0xa3, 0x99, 0x66, 0x41, 0x69, 0xea,
	0x8c, 0x9e, 0xe8, 0xba, 0x61, 0x5b, 0x50, 0x13, 0x48, 0xe3, 0x1b, 0x6a, 0x51, 0x51, 0x50, 0x94,
	0x0c, 0xb0, 0x09, 0x10, 0x9d, 0x12, 0x4a, 0x31, 0x0d, 0x93, 0xb8, 0xf1, 0x40, 0x6d, 0x71, 0xd9,
	0x54, 0x0e, 0xe3, 0xce, 0x8b, 0x62, 0xb6, 0xcf, 0x4a, 0x66, 0x3b, 0xb7, 0x75, 0xef, 0x9b, 0x05,
	0x8d, 0xa3, 0x24, 0x8e, 0x53, 0x3c, 0x23, 0x39, 0xbe, 0xa3, 0xa4, 0x9f, 0x62, 0xdc, 0xd5, 0xbd,
	0x9c, 0x75, 0x78, 0xc8, 0x59, 0x2e, 0x0b, 0x17, 0x3d, 0xed, 0x6a, 0xf1, 0x79, 0x18, 0xcf, 0x25,
	0xb0, 0xe7, 0x13, 0xbc, 0x2a, 0x12, 0x04, 0x25, 0x09, 0xca, 0xdc, 0xbc, 0x5f, 0x36, 0xac, 0xbc,
	0x47, 0x8a, 0x22, 0x11, 0x1f, 0x25, 0x91, 0xe8, 0x7c, 0x80, 0xa5, 0x48, 0xc7, 0x14, 0x0d, 0x6b,
	0x6b, 0x61, 0xa7, 0xb6, 0xd7, 0xf6, 0x4b, 0x4f, 0xa1, 0x3f, 0xb7, 0xb3, 0x83, 0xe5, 0xcb, 0xbf,
	0xad, 0xca, 0xef, 0xeb, 0x8b, 0xb6, 0xd5, 0x9b, 0xb5, 0x71, 0xbe, 0xc2, 0xe3, 0x6c, 0xe6, 0x1f,
	0xa2, 0x0e, 0x10, 0x9a, 0xe4, 0xa2, 0x61, 0x2b, 0x97, 0xfd, 0x7b, 0x5c, 0xca, 0xd2, 0xdf, 0xb6,
	0x6b, 0x66, 0x25, 0x90, 0x70, 0xde, 0x42, 0x95, 0xab, 0xd3, 0xad, 0x7e, 0x7e, 0x6d, 0x6f, 0xfb,
	0x1e, 0x2b, 0x7d, 0x0d, 0x6e, 0x37, 0x36, 0xda, 0x83, 0xd7, 0x97, 0x13, 0xd7, 0xba, 0x9a, 0xb8,
	0xd6, 0xbf, 0x89, 0x6b, 0xfd, 0x9c, 0xba, 0x95, 0xab, 0xa9, 0x5b, 0xf9, 0x33, 0x75, 0x2b, 0x9f,
	0xb7, 0x07, 0x89, 0x3c, 0x1d, 0xf5, 0xfd, 0x88, 0x65, 0x25, 0xf3, 0xef, 0x57, 0xd5, 0x5d, 0xdc,
	0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x01, 0x9d, 0x84, 0x63, 0xf8, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AppVersion) > 0 {
		i -= len(m.AppVersion)
		copy(dAtA[i:], m.AppVersion)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AppVersion)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ExampleVersion) > 0 {
		i -= len(m.ExampleVersion)
		copy(dAtA[i:], m.ExampleVersion)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ExampleVersion)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CallbackCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallbackCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CallbackCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SendPacket != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SendPacket))
		i--
		dAtA[i] = 0x20
	}
	if m.OnTimeoutPacket != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.OnTimeoutPacket))
		i--
		dAtA[i] = 0x18
	}
	if m.OnAcknowledgementPacket != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.OnAcknowledgementPacket))
		i--
		dAtA[i] = 0x10
	}
	if m.OnRecvPacket != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.OnRecvPacket))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MiddlewareEnabledChannel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MiddlewareEnabledChannel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MiddlewareEnabledChannel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.MiddlewareEnabledChannels) > 0 {
		for iNdEx := len(m.MiddlewareEnabledChannels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MiddlewareEnabledChannels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Counters) > 0 {
		for iNdEx := len(m.Counters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Counters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExampleVersion)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AppVersion)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *CallbackCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OnRecvPacket != 0 {
		n += 1 + sovTypes(uint64(m.OnRecvPacket))
	}
	if m.OnAcknowledgementPacket != 0 {
		n += 1 + sovTypes(uint64(m.OnAcknowledgementPacket))
	}
	if m.OnTimeoutPacket != 0 {
		n += 1 + sovTypes(uint64(m.OnTimeoutPacket))
	}
	if m.SendPacket != 0 {
		n += 1 + sovTypes(uint64(m.SendPacket))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *MiddlewareEnabledChannel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Counters) > 0 {
		for _, e := range m.Counters {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.MiddlewareEnabledChannels) > 0 {
		for _, e := range m.MiddlewareEnabledChannels {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExampleVersion", wireType)
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
			m.ExampleVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppVersion", wireType)
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
			m.AppVersion = string(dAtA[iNdEx:postIndex])
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
func (m *CallbackCounter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CallbackCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallbackCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnRecvPacket", wireType)
			}
			m.OnRecvPacket = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OnRecvPacket |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnAcknowledgementPacket", wireType)
			}
			m.OnAcknowledgementPacket = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OnAcknowledgementPacket |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnTimeoutPacket", wireType)
			}
			m.OnTimeoutPacket = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OnTimeoutPacket |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendPacket", wireType)
			}
			m.SendPacket = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendPacket |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
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
			m.ChannelId = string(dAtA[iNdEx:postIndex])
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
func (m *MiddlewareEnabledChannel) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MiddlewareEnabledChannel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MiddlewareEnabledChannel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
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
			m.ChannelId = string(dAtA[iNdEx:postIndex])
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Counters = append(m.Counters, CallbackCounter{})
			if err := m.Counters[len(m.Counters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MiddlewareEnabledChannels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MiddlewareEnabledChannels = append(m.MiddlewareEnabledChannels, MiddlewareEnabledChannel{})
			if err := m.MiddlewareEnabledChannels[len(m.MiddlewareEnabledChannels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
