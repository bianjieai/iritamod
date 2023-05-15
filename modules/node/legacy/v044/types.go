package v044

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
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

// Request defines a standard for validator. The validator will participate the
// blockchain consensus, power determines the probability of proposing a new block.
type OldValidator struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pubkey      string `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Certificate string `protobuf:"bytes,4,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Power       int64  `protobuf:"varint,5,opt,name=power,proto3" json:"power,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Jailed      bool   `protobuf:"varint,7,opt,name=jailed,proto3" json:"jailed,omitempty"`
	Operator    string `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *OldValidator) Reset()         { *m = OldValidator{} }
func (m *OldValidator) String() string { return proto.CompactTextString(m) }
func (*OldValidator) ProtoMessage()    {}
func (*OldValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{0}
}
func (m *OldValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OldValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OldValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *OldValidator) XXX_Size() int {
	return m.Size()
}
func (m *OldValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

// OldHistoricalInfo contains the historical information that gets stored at
// each height.
type OldHistoricalInfo struct {
	Header types.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
	Valset []OldValidator `protobuf:"bytes,2,rep,name=valset,proto3" json:"valset"`
}

func (m *OldHistoricalInfo) Reset()         { *m = OldHistoricalInfo{} }
func (m *OldHistoricalInfo) String() string { return proto.CompactTextString(m) }
func (*OldHistoricalInfo) ProtoMessage()    {}
func (*OldHistoricalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{1}
}
func (m *OldHistoricalInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OldHistoricalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OldHistoricalInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OldHistoricalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldHistoricalInfo.Merge(m, src)
}
func (m *OldHistoricalInfo) XXX_Size() int {
	return m.Size()
}
func (m *OldHistoricalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OldHistoricalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OldHistoricalInfo proto.InternalMessageInfo

// OldNode defines a struct to represent a node identity
type OldNode struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Certificate string `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (m *OldNode) Reset()         { *m = OldNode{} }
func (m *OldNode) String() string { return proto.CompactTextString(m) }
func (*OldNode) ProtoMessage()    {}
func (*OldNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{2}
}
func (m *OldNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OldNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OldNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *OldNode) XXX_Size() int {
	return m.Size()
}
func (m *OldNode) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

// Params defines the parameters for the node module.
type Params struct {
	HistoricalEntries uint32 `protobuf:"varint,1,opt,name=historical_entries,json=historicalEntries,proto3" json:"historical_entries,omitempty" yaml:"historical_entries"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{3}
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

func init() {
	proto.RegisterType((*OldValidator)(nil), "iritamod.node.OldValidator")
	proto.RegisterType((*OldHistoricalInfo)(nil), "iritamod.node.OldHistoricalInfo")
	proto.RegisterType((*OldNode)(nil), "iritamod.node.OldNode")
	proto.RegisterType((*Params)(nil), "iritamod.node.Params")
}

func init() { proto.RegisterFile("node/node.proto", fileDescriptor_a18530e439628818) }

var fileDescriptor_a18530e439628818 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0xbb, 0x69, 0xdc, 0x4e, 0x69, 0xc5, 0xa1, 0xc8, 0xb8, 0x68, 0x36, 0xe4, 0xb4,
	0xa7, 0x44, 0x14, 0x3c, 0xf4, 0xb8, 0x20, 0x54, 0x90, 0x22, 0x39, 0x78, 0x10, 0x41, 0x66, 0x33,
	0xaf, 0xbb, 0xaf, 0x26, 0x99, 0x30, 0x99, 0x55, 0xf6, 0xe6, 0x47, 0xf0, 0x23, 0xf8, 0x71, 0xf6,
	0xd8, 0xa3, 0x20, 0x14, 0xdd, 0xbd, 0x78, 0xf6, 0x13, 0x48, 0x66, 0xd2, 0x6d, 0x75, 0x2f, 0xbd,
	0x0c, 0xf3, 0xfe, 0xef, 0xff, 0xcf, 0xe4, 0xfd, 0x78, 0xf4, 0x7e, 0xa5, 0x24, 0xa4, 0xed, 0x91,
	0xd4, 0x5a, 0x19, 0xc5, 0x0e, 0x51, 0xa3, 0x11, 0xa5, 0x92, 0x49, 0x2b, 0x0e, 0x8f, 0x67, 0x6a,
	0xa6, 0x6c, 0x27, 0x6d, 0x6f, 0xce, 0x34, 0x7c, 0x6c, 0xa0, 0x92, 0xa0, 0x4b, 0xac, 0x4c, 0x6a,
	0x96, 0x35, 0x34, 0xee, 0x74, 0xdd, 0xf8, 0x07, 0xa1, 0xfb, 0x6f, 0x45, 0x81, 0x52, 0x18, 0xa5,
	0xd9, 0x11, 0xed, 0xa1, 0xe4, 0x24, 0x22, 0xe3, 0xfd, 0xac, 0x87, 0x92, 0x31, 0xea, 0x57, 0xa2,
	0x04, 0xde, 0xb3, 0x8a, 0xbd, 0xb3, 0x87, 0x34, 0xa8, 0x17, 0xd3, 0x8f, 0xb0, 0xe4, 0x7d, 0xab,
	0x76, 0x15, 0x8b, 0xe8, 0x41, 0x0e, 0xda, 0xe0, 0x39, 0xe6, 0xc2, 0x00, 0xf7, 0x6d, 0xf3, 0xb6,
	0xc4, 0x8e, 0xe9, 0x5e, 0xad, 0x3e, 0x83, 0xe6, 0x7b, 0x11, 0x19, 0xf7, 0x33, 0x57, 0xb4, 0x39,
	0x09, 0x4d, 0xae, 0xb1, 0x36, 0xa8, 0x2a, 0x1e, 0xb8, 0xdc, 0x2d, 0xa9, 0x7d, 0xf1, 0x42, 0x60,
	0x01, 0x92, 0xdf, 0x8b, 0xc8, 0x78, 0x90, 0x75, 0x15, 0x1b, 0xd2, 0x81, 0xaa, 0x41, 0xb7, 0x7f,
	0xce, 0x07, 0x36, 0xb6, 0xad, 0x4f, 0xfc, 0xdf, 0xdf, 0x46, 0x24, 0xfe, 0x42, 0xe8, 0xd1, 0x29,
	0x36, 0x46, 0x69, 0xcc, 0x45, 0xf1, 0xaa, 0x3a, 0x57, 0xec, 0x05, 0x0d, 0xe6, 0x20, 0x24, 0x68,
	0x3b, 0xe6, 0xc1, 0x33, 0x9e, 0xdc, 0xf0, 0x49, 0x1c, 0x99, 0x53, 0xdb, 0x9f, 0xf8, 0xab, 0xab,
	0x91, 0x97, 0x75, 0xee, 0x36, 0xf7, 0x49, 0x14, 0x0d, 0x18, 0xde, 0x8b, 0xfa, 0x36, 0xf7, 0x0f,
	0xfc, 0x64, 0x0b, 0xf1, 0x3a, 0xe7, 0xdc, 0x71, 0x46, 0xfd, 0x33, 0x25, 0xe1, 0x4e, 0x68, 0xff,
	0x43, 0xd8, 0xdf, 0x41, 0xd8, 0x8d, 0xf5, 0x9e, 0x06, 0x6f, 0x84, 0x16, 0x65, 0xc3, 0x5e, 0x53,
	0x36, 0xdf, 0xce, 0xf7, 0x01, 0x2a, 0xa3, 0x11, 0x1a, 0xfb, 0xca, 0xe1, 0xe4, 0xc9, 0x9f, 0xab,
	0xd1, 0xa3, 0xa5, 0x28, 0x8b, 0x93, 0x78, 0xd7, 0x13, 0x67, 0x0f, 0x6e, 0xc4, 0x97, 0x4e, 0x73,
	0x5f, 0x9f, 0x9c, 0xad, 0x7e, 0x85, 0xde, 0x6a, 0x1d, 0x92, 0xcb, 0x75, 0x48, 0x7e, 0xae, 0x43,
	0xf2, 0x75, 0x13, 0x7a, 0x97, 0x9b, 0xd0, 0xfb, 0xbe, 0x09, 0xbd, 0x77, 0x4f, 0x67, 0x68, 0xe6,
	0x8b, 0x69, 0x92, 0xab, 0x32, 0x9d, 0xa2, 0xa8, 0x2e, 0x10, 0x04, 0xa6, 0xd7, 0x2c, 0xd2, 0x52,
	0xc9, 0x45, 0x01, 0x8d, 0xdd, 0x52, 0xb7, 0x68, 0xd3, 0xc0, 0x6e, 0xda, 0xf3, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x15, 0xbe, 0x84, 0xc3, 0xbf, 0x02, 0x00, 0x00,
}

func (this *OldValidator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OldValidator)
	if !ok {
		that2, ok := that.(OldValidator)
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
	if this.Id != that1.Id {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Pubkey != that1.Pubkey {
		return false
	}
	if this.Certificate != that1.Certificate {
		return false
	}
	if this.Power != that1.Power {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Jailed != that1.Jailed {
		return false
	}
	if this.Operator != that1.Operator {
		return false
	}
	return true
}
func (this *OldNode) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OldNode)
	if !ok {
		that2, ok := that.(OldNode)
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
	if this.Id != that1.Id {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Certificate != that1.Certificate {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.HistoricalEntries != that1.HistoricalEntries {
		return false
	}
	return true
}
func (m *OldValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OldValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OldValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x42
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if m.Power != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Certificate) > 0 {
		i -= len(m.Certificate)
		copy(dAtA[i:], m.Certificate)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Certificate)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OldHistoricalInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OldHistoricalInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OldHistoricalInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Valset) > 0 {
		for iNdEx := len(m.Valset) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Valset[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintNode(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *OldNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OldNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OldNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Certificate) > 0 {
		i -= len(m.Certificate)
		copy(dAtA[i:], m.Certificate)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Certificate)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if m.HistoricalEntries != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.HistoricalEntries))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OldValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovNode(uint64(m.Power))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.Jailed {
		n += 2
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	return n
}

func (m *OldHistoricalInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Header.Size()
	n += 1 + l + sovNode(uint64(l))
	if len(m.Valset) > 0 {
		for _, e := range m.Valset {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	return n
}

func (m *OldNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HistoricalEntries != 0 {
		n += 1 + sovNode(uint64(m.HistoricalEntries))
	}
	return n
}

func sovNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OldValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: OldValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OldValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *OldHistoricalInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: OldHistoricalInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OldHistoricalInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Valset = append(m.Valset, OldValidator{})
			if err := m.Valset[len(m.Valset)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *OldNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: OldNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OldNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalEntries", wireType)
			}
			m.HistoricalEntries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HistoricalEntries |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNode = fmt.Errorf("proto: unexpected end of group")
)
