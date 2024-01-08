// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filespacechain/filespacechain/hosting_offer.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type HostingOffer struct {
	Id            uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Region        string     `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	PricePerBlock types.Coin `protobuf:"bytes,3,opt,name=pricePerBlock,proto3" json:"pricePerBlock"`
	Creator       string     `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *HostingOffer) Reset()         { *m = HostingOffer{} }
func (m *HostingOffer) String() string { return proto.CompactTextString(m) }
func (*HostingOffer) ProtoMessage()    {}
func (*HostingOffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b1ddde0ee789015, []int{0}
}
func (m *HostingOffer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostingOffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostingOffer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostingOffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostingOffer.Merge(m, src)
}
func (m *HostingOffer) XXX_Size() int {
	return m.Size()
}
func (m *HostingOffer) XXX_DiscardUnknown() {
	xxx_messageInfo_HostingOffer.DiscardUnknown(m)
}

var xxx_messageInfo_HostingOffer proto.InternalMessageInfo

func (m *HostingOffer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HostingOffer) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *HostingOffer) GetPricePerBlock() types.Coin {
	if m != nil {
		return m.PricePerBlock
	}
	return types.Coin{}
}

func (m *HostingOffer) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*HostingOffer)(nil), "filespacechain.filespacechain.HostingOffer")
}

func init() {
	proto.RegisterFile("filespacechain/filespacechain/hosting_offer.proto", fileDescriptor_7b1ddde0ee789015)
}

var fileDescriptor_7b1ddde0ee789015 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3f, 0x4b, 0xc4, 0x30,
	0x1c, 0x86, 0x9b, 0xf3, 0x38, 0x31, 0xfe, 0x19, 0x8a, 0x48, 0x3d, 0x30, 0x16, 0xa7, 0x2e, 0x26,
	0x54, 0x17, 0xe7, 0x8a, 0xe0, 0xa6, 0x74, 0x74, 0x91, 0x34, 0x97, 0xb6, 0xc1, 0xbb, 0xfc, 0x6a,
	0x12, 0x45, 0xbf, 0x85, 0x9b, 0x5f, 0xe9, 0xc6, 0x1b, 0x9d, 0x44, 0xda, 0x2f, 0x22, 0x6d, 0x4f,
	0xa4, 0xb7, 0xe5, 0x21, 0xef, 0x9b, 0x87, 0xbc, 0x38, 0xce, 0xd5, 0x5c, 0xda, 0x8a, 0x0b, 0x29,
	0x4a, 0xae, 0x34, 0xdb, 0xc0, 0x12, 0xac, 0x53, 0xba, 0x78, 0x84, 0x3c, 0x97, 0x86, 0x56, 0x06,
	0x1c, 0xf8, 0x27, 0xc3, 0x0c, 0x1d, 0xe2, 0xf4, 0xb0, 0x80, 0x02, 0xba, 0x24, 0x6b, 0x4f, 0x7d,
	0x69, 0x4a, 0x04, 0xd8, 0x05, 0x58, 0x96, 0x71, 0x2b, 0xd9, 0x6b, 0x9c, 0x49, 0xc7, 0x63, 0x26,
	0x40, 0xe9, 0xfe, 0xfe, 0xec, 0x13, 0xe1, 0xbd, 0xdb, 0x5e, 0x76, 0xd7, 0xba, 0xfc, 0x03, 0x3c,
	0x52, 0xb3, 0x00, 0x85, 0x28, 0x1a, 0xa7, 0x23, 0x35, 0xf3, 0x8f, 0xf0, 0xc4, 0xc8, 0x42, 0x81,
	0x0e, 0x46, 0x21, 0x8a, 0x76, 0xd2, 0x35, 0xf9, 0x37, 0x78, 0xbf, 0x32, 0x4a, 0xc8, 0x7b, 0x69,
	0x92, 0x39, 0x88, 0xa7, 0x60, 0x2b, 0x44, 0xd1, 0xee, 0xc5, 0x31, 0xed, 0x85, 0xb4, 0x15, 0xd2,
	0xb5, 0x90, 0x5e, 0x83, 0xd2, 0xc9, 0x78, 0xf9, 0x7d, 0xea, 0xa5, 0xc3, 0x96, 0x1f, 0xe0, 0x6d,
	0x61, 0x24, 0x77, 0x60, 0x82, 0x71, 0xf7, 0xfe, 0x1f, 0x26, 0xe9, 0xb2, 0x26, 0x68, 0x55, 0x13,
	0xf4, 0x53, 0x13, 0xf4, 0xd1, 0x10, 0x6f, 0xd5, 0x10, 0xef, 0xab, 0x21, 0xde, 0xc3, 0x55, 0xa1,
	0x5c, 0xf9, 0x92, 0x51, 0x01, 0x0b, 0x56, 0x72, 0x6d, 0xcb, 0xe7, 0xff, 0xf9, 0xce, 0xfb, 0xfd,
	0xde, 0x36, 0x07, 0x75, 0xef, 0x95, 0xb4, 0xd9, 0xa4, 0xfb, 0xf4, 0xe5, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xef, 0x9b, 0x31, 0x6a, 0x7e, 0x01, 0x00, 0x00,
}

func (m *HostingOffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostingOffer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostingOffer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintHostingOffer(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.PricePerBlock.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHostingOffer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Region) > 0 {
		i -= len(m.Region)
		copy(dAtA[i:], m.Region)
		i = encodeVarintHostingOffer(dAtA, i, uint64(len(m.Region)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintHostingOffer(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHostingOffer(dAtA []byte, offset int, v uint64) int {
	offset -= sovHostingOffer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostingOffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovHostingOffer(uint64(m.Id))
	}
	l = len(m.Region)
	if l > 0 {
		n += 1 + l + sovHostingOffer(uint64(l))
	}
	l = m.PricePerBlock.Size()
	n += 1 + l + sovHostingOffer(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovHostingOffer(uint64(l))
	}
	return n
}

func sovHostingOffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHostingOffer(x uint64) (n int) {
	return sovHostingOffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostingOffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHostingOffer
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
			return fmt.Errorf("proto: HostingOffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostingOffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Region", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingOffer
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
				return ErrInvalidLengthHostingOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostingOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Region = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PricePerBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingOffer
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
				return ErrInvalidLengthHostingOffer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostingOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PricePerBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingOffer
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
				return ErrInvalidLengthHostingOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostingOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHostingOffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHostingOffer
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
func skipHostingOffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHostingOffer
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
					return 0, ErrIntOverflowHostingOffer
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
					return 0, ErrIntOverflowHostingOffer
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
				return 0, ErrInvalidLengthHostingOffer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHostingOffer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHostingOffer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHostingOffer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHostingOffer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHostingOffer = fmt.Errorf("proto: unexpected end of group")
)