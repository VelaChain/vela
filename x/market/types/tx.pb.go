// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgCreatePool struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	AmountA   string `protobuf:"bytes,2,opt,name=amountA,proto3" json:"amountA,omitempty"`
	DenomA    string `protobuf:"bytes,3,opt,name=denomA,proto3" json:"denomA,omitempty"`
	AmountB   string `protobuf:"bytes,4,opt,name=amountB,proto3" json:"amountB,omitempty"`
	DenomB    string `protobuf:"bytes,5,opt,name=denomB,proto3" json:"denomB,omitempty"`
	MinShares string `protobuf:"bytes,6,opt,name=minShares,proto3" json:"minShares,omitempty"`
}

func (m *MsgCreatePool) Reset()         { *m = MsgCreatePool{} }
func (m *MsgCreatePool) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePool) ProtoMessage()    {}
func (*MsgCreatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_2966ca2342567dca, []int{0}
}
func (m *MsgCreatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePool.Merge(m, src)
}
func (m *MsgCreatePool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePool proto.InternalMessageInfo

func (m *MsgCreatePool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePool) GetAmountA() string {
	if m != nil {
		return m.AmountA
	}
	return ""
}

func (m *MsgCreatePool) GetDenomA() string {
	if m != nil {
		return m.DenomA
	}
	return ""
}

func (m *MsgCreatePool) GetAmountB() string {
	if m != nil {
		return m.AmountB
	}
	return ""
}

func (m *MsgCreatePool) GetDenomB() string {
	if m != nil {
		return m.DenomB
	}
	return ""
}

func (m *MsgCreatePool) GetMinShares() string {
	if m != nil {
		return m.MinShares
	}
	return ""
}

type MsgCreatePoolResponse struct {
}

func (m *MsgCreatePoolResponse) Reset()         { *m = MsgCreatePoolResponse{} }
func (m *MsgCreatePoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePoolResponse) ProtoMessage()    {}
func (*MsgCreatePoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2966ca2342567dca, []int{1}
}
func (m *MsgCreatePoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePoolResponse.Merge(m, src)
}
func (m *MsgCreatePoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePoolResponse proto.InternalMessageInfo

type MsgJoinPool struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	AmountA   string `protobuf:"bytes,2,opt,name=amountA,proto3" json:"amountA,omitempty"`
	DenomA    string `protobuf:"bytes,3,opt,name=denomA,proto3" json:"denomA,omitempty"`
	AmountB   string `protobuf:"bytes,4,opt,name=amountB,proto3" json:"amountB,omitempty"`
	DenomB    string `protobuf:"bytes,5,opt,name=denomB,proto3" json:"denomB,omitempty"`
	MinShares string `protobuf:"bytes,6,opt,name=minShares,proto3" json:"minShares,omitempty"`
}

func (m *MsgJoinPool) Reset()         { *m = MsgJoinPool{} }
func (m *MsgJoinPool) String() string { return proto.CompactTextString(m) }
func (*MsgJoinPool) ProtoMessage()    {}
func (*MsgJoinPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_2966ca2342567dca, []int{2}
}
func (m *MsgJoinPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgJoinPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgJoinPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgJoinPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgJoinPool.Merge(m, src)
}
func (m *MsgJoinPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgJoinPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgJoinPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgJoinPool proto.InternalMessageInfo

func (m *MsgJoinPool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgJoinPool) GetAmountA() string {
	if m != nil {
		return m.AmountA
	}
	return ""
}

func (m *MsgJoinPool) GetDenomA() string {
	if m != nil {
		return m.DenomA
	}
	return ""
}

func (m *MsgJoinPool) GetAmountB() string {
	if m != nil {
		return m.AmountB
	}
	return ""
}

func (m *MsgJoinPool) GetDenomB() string {
	if m != nil {
		return m.DenomB
	}
	return ""
}

func (m *MsgJoinPool) GetMinShares() string {
	if m != nil {
		return m.MinShares
	}
	return ""
}

type MsgJoinPoolResponse struct {
}

func (m *MsgJoinPoolResponse) Reset()         { *m = MsgJoinPoolResponse{} }
func (m *MsgJoinPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgJoinPoolResponse) ProtoMessage()    {}
func (*MsgJoinPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2966ca2342567dca, []int{3}
}
func (m *MsgJoinPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgJoinPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgJoinPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgJoinPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgJoinPoolResponse.Merge(m, src)
}
func (m *MsgJoinPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgJoinPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgJoinPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgJoinPoolResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePool)(nil), "VelaChain.vela.market.MsgCreatePool")
	proto.RegisterType((*MsgCreatePoolResponse)(nil), "VelaChain.vela.market.MsgCreatePoolResponse")
	proto.RegisterType((*MsgJoinPool)(nil), "VelaChain.vela.market.MsgJoinPool")
	proto.RegisterType((*MsgJoinPoolResponse)(nil), "VelaChain.vela.market.MsgJoinPoolResponse")
}

func init() { proto.RegisterFile("market/tx.proto", fileDescriptor_2966ca2342567dca) }

var fileDescriptor_2966ca2342567dca = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0xd1, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x0d, 0x4b, 0xcd,
	0x49, 0x74, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0x4b, 0xcd, 0x49, 0xd4, 0x83, 0xc8, 0x2b, 0xad,
	0x66, 0xe4, 0xe2, 0xf5, 0x2d, 0x4e, 0x77, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x0d, 0xc8, 0xcf, 0xcf,
	0x11, 0x92, 0xe0, 0x62, 0x4f, 0x06, 0xf1, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x60, 0x5c, 0x90, 0x4c, 0x62, 0x6e, 0x7e, 0x69, 0x5e, 0x89, 0xa3, 0x04, 0x13, 0x44, 0x06, 0xca,
	0x15, 0x12, 0xe3, 0x62, 0x4b, 0x49, 0xcd, 0xcb, 0xcf, 0x75, 0x94, 0x60, 0x06, 0x4b, 0x40, 0x79,
	0x08, 0x1d, 0x4e, 0x12, 0x2c, 0xc8, 0x3a, 0x9c, 0xe0, 0x3a, 0x9c, 0x24, 0x58, 0x91, 0x74, 0x38,
	0x09, 0xc9, 0x70, 0x71, 0xe6, 0x66, 0xe6, 0x05, 0x67, 0x24, 0x16, 0xa5, 0x16, 0x4b, 0xb0, 0x81,
	0xa5, 0x10, 0x02, 0x4a, 0xe2, 0x5c, 0xa2, 0x28, 0x8e, 0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0x55, 0x5a, 0xc9, 0xc8, 0xc5, 0xed, 0x5b, 0x9c, 0xee, 0x95, 0x9f, 0x99, 0x37, 0xe8, 0x3d,
	0x21, 0xca, 0x25, 0x8c, 0xe4, 0x54, 0x98, 0x17, 0x8c, 0x0e, 0x33, 0x72, 0x31, 0xfb, 0x16, 0xa7,
	0x0b, 0x25, 0x70, 0x71, 0x21, 0xc5, 0x86, 0x8a, 0x1e, 0xd6, 0x78, 0xd3, 0x43, 0x09, 0x06, 0x29,
	0x1d, 0x62, 0x54, 0xc1, 0x6c, 0x12, 0x8a, 0xe2, 0xe2, 0x80, 0x07, 0x94, 0x12, 0x6e, 0x9d, 0x30,
	0x35, 0x52, 0x5a, 0x84, 0xd5, 0xc0, 0xcc, 0x76, 0x72, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6,
	0x63, 0x39, 0x86, 0x28, 0x8d, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d,
	0xb8, 0x79, 0xfa, 0x20, 0xf3, 0xf4, 0x2b, 0xf4, 0x61, 0xa9, 0xb5, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0x9c, 0x62, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xeb, 0xcc, 0x8d, 0xc4, 0x02,
	0x00, 0x00,
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
	CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
	JoinPool(ctx context.Context, in *MsgJoinPool, opts ...grpc.CallOption) (*MsgJoinPoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, "/VelaChain.vela.market.Msg/CreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinPool(ctx context.Context, in *MsgJoinPool, opts ...grpc.CallOption) (*MsgJoinPoolResponse, error) {
	out := new(MsgJoinPoolResponse)
	err := c.cc.Invoke(ctx, "/VelaChain.vela.market.Msg/JoinPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error)
	JoinPool(context.Context, *MsgJoinPool) (*MsgJoinPoolResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePool(ctx context.Context, req *MsgCreatePool) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (*UnimplementedMsgServer) JoinPool(ctx context.Context, req *MsgJoinPool) (*MsgJoinPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VelaChain.vela.market.Msg/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*MsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VelaChain.vela.market.Msg/JoinPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinPool(ctx, req.(*MsgJoinPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "VelaChain.vela.market.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
		{
			MethodName: "JoinPool",
			Handler:    _Msg_JoinPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market/tx.proto",
}

func (m *MsgCreatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinShares) > 0 {
		i -= len(m.MinShares)
		copy(dAtA[i:], m.MinShares)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MinShares)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DenomB) > 0 {
		i -= len(m.DenomB)
		copy(dAtA[i:], m.DenomB)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DenomB)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AmountB) > 0 {
		i -= len(m.AmountB)
		copy(dAtA[i:], m.AmountB)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AmountB)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DenomA) > 0 {
		i -= len(m.DenomA)
		copy(dAtA[i:], m.DenomA)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DenomA)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AmountA) > 0 {
		i -= len(m.AmountA)
		copy(dAtA[i:], m.AmountA)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AmountA)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgJoinPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgJoinPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgJoinPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinShares) > 0 {
		i -= len(m.MinShares)
		copy(dAtA[i:], m.MinShares)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MinShares)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DenomB) > 0 {
		i -= len(m.DenomB)
		copy(dAtA[i:], m.DenomB)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DenomB)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AmountB) > 0 {
		i -= len(m.AmountB)
		copy(dAtA[i:], m.AmountB)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AmountB)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DenomA) > 0 {
		i -= len(m.DenomA)
		copy(dAtA[i:], m.DenomA)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DenomA)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AmountA) > 0 {
		i -= len(m.AmountA)
		copy(dAtA[i:], m.AmountA)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AmountA)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgJoinPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgJoinPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgJoinPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AmountA)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DenomA)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AmountB)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DenomB)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MinShares)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreatePoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgJoinPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AmountA)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DenomA)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AmountB)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DenomB)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MinShares)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgJoinPoolResponse) Size() (n int) {
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
func (m *MsgCreatePool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountA", wireType)
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
			m.AmountA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomA", wireType)
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
			m.DenomA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountB", wireType)
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
			m.AmountB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomB", wireType)
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
			m.DenomB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinShares", wireType)
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
			m.MinShares = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreatePoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgJoinPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgJoinPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgJoinPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountA", wireType)
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
			m.AmountA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomA", wireType)
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
			m.DenomA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountB", wireType)
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
			m.AmountB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomB", wireType)
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
			m.DenomB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinShares", wireType)
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
			m.MinShares = string(dAtA[iNdEx:postIndex])
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
func (m *MsgJoinPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgJoinPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgJoinPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
