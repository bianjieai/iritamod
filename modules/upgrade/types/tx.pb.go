// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iritamod/upgrade/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// MsgUpgradeSoftware - struct for upgrade software
type MsgUpgradeSoftware struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The height at which the upgrade must be performed.
	// Only used if Time is not set.
	Height int64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Info     string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Operator string `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgUpgradeSoftware) Reset()         { *m = MsgUpgradeSoftware{} }
func (m *MsgUpgradeSoftware) String() string { return proto.CompactTextString(m) }
func (*MsgUpgradeSoftware) ProtoMessage()    {}
func (*MsgUpgradeSoftware) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51a115646a2cd7c, []int{0}
}
func (m *MsgUpgradeSoftware) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpgradeSoftware) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpgradeSoftware.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpgradeSoftware) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpgradeSoftware.Merge(m, src)
}
func (m *MsgUpgradeSoftware) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpgradeSoftware) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpgradeSoftware.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpgradeSoftware proto.InternalMessageInfo

func (m *MsgUpgradeSoftware) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgUpgradeSoftware) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MsgUpgradeSoftware) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *MsgUpgradeSoftware) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

// MsgUpgradeSoftwareResponse defines the Msg/UpgradeSoftware response type.
type MsgUpgradeSoftwareResponse struct {
}

func (m *MsgUpgradeSoftwareResponse) Reset()         { *m = MsgUpgradeSoftwareResponse{} }
func (m *MsgUpgradeSoftwareResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpgradeSoftwareResponse) ProtoMessage()    {}
func (*MsgUpgradeSoftwareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51a115646a2cd7c, []int{1}
}
func (m *MsgUpgradeSoftwareResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpgradeSoftwareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpgradeSoftwareResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpgradeSoftwareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpgradeSoftwareResponse.Merge(m, src)
}
func (m *MsgUpgradeSoftwareResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpgradeSoftwareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpgradeSoftwareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpgradeSoftwareResponse proto.InternalMessageInfo

// MsgCancelUpgrade - struct for cancel software upgrade
type MsgCancelUpgrade struct {
	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgCancelUpgrade) Reset()         { *m = MsgCancelUpgrade{} }
func (m *MsgCancelUpgrade) String() string { return proto.CompactTextString(m) }
func (*MsgCancelUpgrade) ProtoMessage()    {}
func (*MsgCancelUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51a115646a2cd7c, []int{2}
}
func (m *MsgCancelUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelUpgrade.Merge(m, src)
}
func (m *MsgCancelUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelUpgrade proto.InternalMessageInfo

func (m *MsgCancelUpgrade) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

// MsgCancelUpgradeResponse defines the Msg/CancelUpgrade response type.
type MsgCancelUpgradeResponse struct {
}

func (m *MsgCancelUpgradeResponse) Reset()         { *m = MsgCancelUpgradeResponse{} }
func (m *MsgCancelUpgradeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCancelUpgradeResponse) ProtoMessage()    {}
func (*MsgCancelUpgradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51a115646a2cd7c, []int{3}
}
func (m *MsgCancelUpgradeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelUpgradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelUpgradeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelUpgradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelUpgradeResponse.Merge(m, src)
}
func (m *MsgCancelUpgradeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelUpgradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelUpgradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelUpgradeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpgradeSoftware)(nil), "iritamod.upgrade.MsgUpgradeSoftware")
	proto.RegisterType((*MsgUpgradeSoftwareResponse)(nil), "iritamod.upgrade.MsgUpgradeSoftwareResponse")
	proto.RegisterType((*MsgCancelUpgrade)(nil), "iritamod.upgrade.MsgCancelUpgrade")
	proto.RegisterType((*MsgCancelUpgradeResponse)(nil), "iritamod.upgrade.MsgCancelUpgradeResponse")
}

func init() { proto.RegisterFile("iritamod/upgrade/tx.proto", fileDescriptor_c51a115646a2cd7c) }

var fileDescriptor_c51a115646a2cd7c = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4a, 0x3b, 0x31,
	0x14, 0xc5, 0x9b, 0x7f, 0x4b, 0xf9, 0x1b, 0x10, 0x4b, 0x10, 0x19, 0x83, 0x4c, 0xcb, 0xe0, 0xa2,
	0x8a, 0x4c, 0x40, 0x77, 0x2e, 0xed, 0xba, 0x9b, 0x8a, 0x1b, 0x37, 0x92, 0xda, 0xdb, 0x34, 0xd2,
	0x99, 0x1b, 0x92, 0xd4, 0x8f, 0xb7, 0xf0, 0x11, 0x7c, 0x1c, 0x37, 0x42, 0x97, 0x2e, 0xa5, 0xdd,
	0xf8, 0x18, 0xd2, 0x99, 0x4e, 0x65, 0x5a, 0x41, 0x77, 0xf7, 0xe6, 0xfc, 0x72, 0xce, 0xcd, 0x07,
	0xdd, 0xd7, 0x56, 0x7b, 0x99, 0xe0, 0x40, 0x4c, 0x8c, 0xb2, 0x72, 0x00, 0xc2, 0x3f, 0xc6, 0xc6,
	0xa2, 0x47, 0xd6, 0x28, 0xa4, 0x78, 0x29, 0xf1, 0x5d, 0x85, 0x0a, 0x33, 0x51, 0x2c, 0xaa, 0x9c,
	0xe3, 0x4d, 0x85, 0xa8, 0xc6, 0x20, 0xb2, 0xae, 0x3f, 0x19, 0x0a, 0xaf, 0x13, 0x70, 0x5e, 0x26,
	0x26, 0x07, 0xa2, 0x7b, 0xca, 0xba, 0x4e, 0x5d, 0xe5, 0x26, 0x97, 0x38, 0xf4, 0x0f, 0xd2, 0x02,
	0x63, 0xb4, 0x96, 0xca, 0x04, 0x02, 0xd2, 0x22, 0xed, 0xad, 0x5e, 0x56, 0xb3, 0x3d, 0x5a, 0x1f,
	0x81, 0x56, 0x23, 0x1f, 0xfc, 0x6b, 0x91, 0x76, 0xb5, 0xb7, 0xec, 0x16, 0xac, 0x4e, 0x87, 0x18,
	0x54, 0x73, 0x76, 0x51, 0x33, 0x4e, 0xff, 0xa3, 0x01, 0x2b, 0x3d, 0xda, 0xa0, 0x96, 0xad, 0xaf,
	0xfa, 0xf3, 0xda, 0xe7, 0x4b, 0x93, 0x44, 0x07, 0x94, 0x6f, 0xe6, 0xf6, 0xc0, 0x19, 0x4c, 0x1d,
	0x44, 0x31, 0x6d, 0x74, 0x9d, 0xea, 0xc8, 0xf4, 0x16, 0xc6, 0x4b, 0xa6, 0xe4, 0x49, 0xca, 0x9e,
	0x11, 0xa7, 0xc1, 0x3a, 0x5f, 0x78, 0x9d, 0xbe, 0x11, 0x5a, 0xed, 0x3a, 0xc5, 0x80, 0xee, 0xac,
	0x1f, 0xf3, 0x30, 0x5e, 0xbf, 0xc6, 0x78, 0x73, 0x28, 0x7e, 0xf2, 0x17, 0xaa, 0x88, 0x63, 0x37,
	0x74, 0xbb, 0x3c, 0x77, 0xf4, 0xe3, 0xf6, 0x12, 0xc3, 0x8f, 0x7f, 0x67, 0x8a, 0x80, 0x8b, 0xce,
	0xeb, 0x2c, 0x24, 0xd3, 0x59, 0x48, 0x3e, 0x66, 0x21, 0x79, 0x9e, 0x87, 0x95, 0xe9, 0x3c, 0xac,
	0xbc, 0xcf, 0xc3, 0xca, 0xf5, 0xd1, 0xca, 0xa4, 0xaf, 0x65, 0x7a, 0xa7, 0x21, 0x96, 0x5a, 0x24,
	0x38, 0x98, 0x8c, 0xc1, 0x7d, 0x7f, 0xa1, 0x27, 0x03, 0xae, 0x5f, 0xcf, 0x5e, 0xff, 0xec, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0xa2, 0x6d, 0x89, 0x75, 0x63, 0x02, 0x00, 0x00,
}

func (this *MsgUpgradeSoftware) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUpgradeSoftware)
	if !ok {
		that2, ok := that.(MsgUpgradeSoftware)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.Info != that1.Info {
		return false
	}
	if this.Operator != that1.Operator {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateClient defines a method for creating a light client.
	UpgradeSoftware(ctx context.Context, in *MsgUpgradeSoftware, opts ...grpc.CallOption) (*MsgUpgradeSoftwareResponse, error)
	// CancelUpgrade defines a method for updating a light client state.
	CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpgradeSoftware(ctx context.Context, in *MsgUpgradeSoftware, opts ...grpc.CallOption) (*MsgUpgradeSoftwareResponse, error) {
	out := new(MsgUpgradeSoftwareResponse)
	err := c.cc.Invoke(ctx, "/iritamod.upgrade.Msg/UpgradeSoftware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error) {
	out := new(MsgCancelUpgradeResponse)
	err := c.cc.Invoke(ctx, "/iritamod.upgrade.Msg/CancelUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateClient defines a method for creating a light client.
	UpgradeSoftware(context.Context, *MsgUpgradeSoftware) (*MsgUpgradeSoftwareResponse, error)
	// CancelUpgrade defines a method for updating a light client state.
	CancelUpgrade(context.Context, *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpgradeSoftware(ctx context.Context, req *MsgUpgradeSoftware) (*MsgUpgradeSoftwareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeSoftware not implemented")
}
func (*UnimplementedMsgServer) CancelUpgrade(ctx context.Context, req *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUpgrade not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpgradeSoftware_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpgradeSoftware)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpgradeSoftware(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iritamod.upgrade.Msg/UpgradeSoftware",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpgradeSoftware(ctx, req.(*MsgUpgradeSoftware))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iritamod.upgrade.Msg/CancelUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelUpgrade(ctx, req.(*MsgCancelUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iritamod.upgrade.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpgradeSoftware",
			Handler:    _Msg_UpgradeSoftware_Handler,
		},
		{
			MethodName: "CancelUpgrade",
			Handler:    _Msg_CancelUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iritamod/upgrade/tx.proto",
}

func (m *MsgUpgradeSoftware) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpgradeSoftware) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpgradeSoftware) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Height != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpgradeSoftwareResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpgradeSoftwareResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpgradeSoftwareResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCancelUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCancelUpgradeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelUpgradeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelUpgradeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpgradeSoftware) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTx(uint64(m.Height))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpgradeSoftwareResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCancelUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCancelUpgradeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpgradeSoftware) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpgradeSoftware: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpgradeSoftware: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpgradeSoftwareResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpgradeSoftwareResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpgradeSoftwareResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCancelUpgrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCancelUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCancelUpgradeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCancelUpgradeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelUpgradeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
