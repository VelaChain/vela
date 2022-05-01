// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/liq_prov.proto

package types

import (
	fmt "fmt"
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

type LiqProv struct {
	PoolName    string `protobuf:"bytes,1,opt,name=poolName,proto3" json:"poolName,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ShareAmount string `protobuf:"bytes,3,opt,name=shareAmount,proto3" json:"shareAmount,omitempty"`
}

func (m *LiqProv) Reset()         { *m = LiqProv{} }
func (m *LiqProv) String() string { return proto.CompactTextString(m) }
func (*LiqProv) ProtoMessage()    {}
func (*LiqProv) Descriptor() ([]byte, []int) {
	return fileDescriptor_75b8024f7bfc81a0, []int{0}
}
func (m *LiqProv) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiqProv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiqProv.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiqProv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiqProv.Merge(m, src)
}
func (m *LiqProv) XXX_Size() int {
	return m.Size()
}
func (m *LiqProv) XXX_DiscardUnknown() {
	xxx_messageInfo_LiqProv.DiscardUnknown(m)
}

var xxx_messageInfo_LiqProv proto.InternalMessageInfo

func (m *LiqProv) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *LiqProv) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *LiqProv) GetShareAmount() string {
	if m != nil {
		return m.ShareAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*LiqProv)(nil), "VelaChain.vela.market.LiqProv")
}

func init() { proto.RegisterFile("market/liq_prov.proto", fileDescriptor_75b8024f7bfc81a0) }

var fileDescriptor_75b8024f7bfc81a0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0xd1, 0xcf, 0xc9, 0x2c, 0x8c, 0x2f, 0x28, 0xca, 0x2f, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x0d, 0x4b, 0xcd, 0x49, 0x74, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0x4b, 0xcd,
	0x49, 0xd4, 0x83, 0xa8, 0x52, 0x4a, 0xe4, 0x62, 0xf7, 0xc9, 0x2c, 0x0c, 0x28, 0xca, 0x2f, 0x13,
	0x92, 0xe2, 0xe2, 0x28, 0xc8, 0xcf, 0xcf, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x24, 0xb8, 0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25,
	0x98, 0xc0, 0x52, 0x30, 0xae, 0x90, 0x02, 0x17, 0x77, 0x71, 0x46, 0x62, 0x51, 0xaa, 0x63, 0x6e,
	0x7e, 0x69, 0x5e, 0x89, 0x04, 0x33, 0x58, 0x16, 0x59, 0xc8, 0xc9, 0xe9, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x34, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0xe1, 0xce, 0xd3, 0x07, 0x39, 0x4f, 0xbf, 0x42, 0x1f, 0xea, 0x8d, 0x92, 0xca, 0x82,
	0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x27, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x39, 0x61,
	0xfa, 0xdd, 0x00, 0x00, 0x00,
}

func (m *LiqProv) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiqProv) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiqProv) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShareAmount) > 0 {
		i -= len(m.ShareAmount)
		copy(dAtA[i:], m.ShareAmount)
		i = encodeVarintLiqProv(dAtA, i, uint64(len(m.ShareAmount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintLiqProv(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolName) > 0 {
		i -= len(m.PoolName)
		copy(dAtA[i:], m.PoolName)
		i = encodeVarintLiqProv(dAtA, i, uint64(len(m.PoolName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiqProv(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiqProv(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiqProv) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolName)
	if l > 0 {
		n += 1 + l + sovLiqProv(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovLiqProv(uint64(l))
	}
	l = len(m.ShareAmount)
	if l > 0 {
		n += 1 + l + sovLiqProv(uint64(l))
	}
	return n
}

func sovLiqProv(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiqProv(x uint64) (n int) {
	return sovLiqProv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiqProv) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiqProv
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
			return fmt.Errorf("proto: LiqProv: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiqProv: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiqProv
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
				return ErrInvalidLengthLiqProv
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiqProv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiqProv
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
				return ErrInvalidLengthLiqProv
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiqProv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiqProv
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
				return ErrInvalidLengthLiqProv
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiqProv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShareAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiqProv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiqProv
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
func skipLiqProv(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiqProv
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
					return 0, ErrIntOverflowLiqProv
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
					return 0, ErrIntOverflowLiqProv
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
				return 0, ErrInvalidLengthLiqProv
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiqProv
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiqProv
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiqProv        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiqProv          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiqProv = fmt.Errorf("proto: unexpected end of group")
)