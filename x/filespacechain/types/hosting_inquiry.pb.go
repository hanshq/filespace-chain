// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filespacechain/filespacechain/hosting_inquiry.proto

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

type HostingInquiry struct {
	Id               uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FileEntryCid     string     `protobuf:"bytes,2,opt,name=fileEntryCid,proto3" json:"fileEntryCid,omitempty"`
	ReplicationRate  uint64     `protobuf:"varint,3,opt,name=replicationRate,proto3" json:"replicationRate,omitempty"`
	EscrowAmount     types.Coin `protobuf:"bytes,4,opt,name=escrowAmount,proto3" json:"escrowAmount"`
	EndTime          uint64     `protobuf:"varint,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Creator          string     `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	MaxPricePerBlock uint64     `protobuf:"varint,7,opt,name=maxPricePerBlock,proto3" json:"maxPricePerBlock,omitempty"`
}

func (m *HostingInquiry) Reset()         { *m = HostingInquiry{} }
func (m *HostingInquiry) String() string { return proto.CompactTextString(m) }
func (*HostingInquiry) ProtoMessage()    {}
func (*HostingInquiry) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a298bb519cdbacc, []int{0}
}
func (m *HostingInquiry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostingInquiry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostingInquiry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostingInquiry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostingInquiry.Merge(m, src)
}
func (m *HostingInquiry) XXX_Size() int {
	return m.Size()
}
func (m *HostingInquiry) XXX_DiscardUnknown() {
	xxx_messageInfo_HostingInquiry.DiscardUnknown(m)
}

var xxx_messageInfo_HostingInquiry proto.InternalMessageInfo

func (m *HostingInquiry) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HostingInquiry) GetFileEntryCid() string {
	if m != nil {
		return m.FileEntryCid
	}
	return ""
}

func (m *HostingInquiry) GetReplicationRate() uint64 {
	if m != nil {
		return m.ReplicationRate
	}
	return 0
}

func (m *HostingInquiry) GetEscrowAmount() types.Coin {
	if m != nil {
		return m.EscrowAmount
	}
	return types.Coin{}
}

func (m *HostingInquiry) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *HostingInquiry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *HostingInquiry) GetMaxPricePerBlock() uint64 {
	if m != nil {
		return m.MaxPricePerBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*HostingInquiry)(nil), "filespacechain.filespacechain.HostingInquiry")
}

func init() {
	proto.RegisterFile("filespacechain/filespacechain/hosting_inquiry.proto", fileDescriptor_6a298bb519cdbacc)
}

var fileDescriptor_6a298bb519cdbacc = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xc1, 0x6a, 0xdb, 0x40,
	0x14, 0x94, 0x54, 0xd7, 0xa6, 0x5b, 0xe3, 0x16, 0xd1, 0xc3, 0xd6, 0x50, 0xd5, 0xf8, 0x24, 0x0a,
	0x95, 0x70, 0x7c, 0xc9, 0x35, 0x36, 0x81, 0xe4, 0x66, 0x44, 0x4e, 0xb9, 0x84, 0xd5, 0x6a, 0x23,
	0x3d, 0x62, 0xed, 0x93, 0x77, 0xd7, 0x89, 0xfd, 0x17, 0x81, 0xfc, 0x94, 0x8f, 0x3e, 0xe6, 0x14,
	0x82, 0xfd, 0x23, 0x41, 0x52, 0x4c, 0x90, 0x73, 0xdb, 0x99, 0xf7, 0x66, 0x67, 0x78, 0x43, 0xc6,
	0xb7, 0x30, 0x17, 0xba, 0x60, 0x5c, 0xf0, 0x8c, 0x81, 0x0c, 0x8f, 0x60, 0x86, 0xda, 0x80, 0x4c,
	0x6f, 0x40, 0x2e, 0x96, 0xa0, 0xd6, 0x41, 0xa1, 0xd0, 0xa0, 0xfb, 0xa7, 0xb9, 0x15, 0x34, 0x61,
	0xff, 0x57, 0x8a, 0x29, 0x56, 0x9b, 0x61, 0xf9, 0xaa, 0x45, 0x7d, 0x8f, 0xa3, 0xce, 0x51, 0x87,
	0x31, 0xd3, 0x22, 0xbc, 0x1f, 0xc5, 0xc2, 0xb0, 0x51, 0xc8, 0x11, 0x64, 0x3d, 0x1f, 0x3e, 0x39,
	0xa4, 0x77, 0x51, 0xdb, 0x5d, 0xd6, 0x6e, 0x6e, 0x8f, 0x38, 0x90, 0x50, 0x7b, 0x60, 0xfb, 0xad,
	0xc8, 0x81, 0xc4, 0x1d, 0x92, 0x6e, 0x69, 0x75, 0x2e, 0x8d, 0x5a, 0x4f, 0x21, 0xa1, 0xce, 0xc0,
	0xf6, 0xbf, 0x45, 0x0d, 0xce, 0xf5, 0xc9, 0x0f, 0x25, 0x8a, 0x39, 0x70, 0x66, 0x00, 0x65, 0xc4,
	0x8c, 0xa0, 0x5f, 0xaa, 0x0f, 0x8e, 0x69, 0x77, 0x4a, 0xba, 0x42, 0x73, 0x85, 0x0f, 0x67, 0x39,
	0x2e, 0xa5, 0xa1, 0xad, 0x81, 0xed, 0x7f, 0x3f, 0xf9, 0x1d, 0xd4, 0x39, 0x83, 0x32, 0x67, 0xf0,
	0x9e, 0x33, 0x98, 0x22, 0xc8, 0x49, 0x6b, 0xf3, 0xf2, 0xd7, 0x8a, 0x1a, 0x22, 0x97, 0x92, 0x8e,
	0x90, 0xc9, 0x15, 0xe4, 0x82, 0x7e, 0xad, 0x6c, 0x0e, 0xb0, 0x9c, 0x70, 0x25, 0x98, 0x41, 0x45,
	0xdb, 0x55, 0xce, 0x03, 0x74, 0xff, 0x91, 0x9f, 0x39, 0x5b, 0xcd, 0x14, 0x70, 0x31, 0x13, 0x6a,
	0x32, 0x47, 0x7e, 0x47, 0x3b, 0x95, 0xf8, 0x13, 0x3f, 0x89, 0x36, 0x3b, 0xcf, 0xde, 0xee, 0x3c,
	0xfb, 0x75, 0xe7, 0xd9, 0x8f, 0x7b, 0xcf, 0xda, 0xee, 0x3d, 0xeb, 0x79, 0xef, 0x59, 0xd7, 0xa7,
	0x29, 0x98, 0x6c, 0x19, 0x07, 0x1c, 0xf3, 0x30, 0x63, 0x52, 0x67, 0x8b, 0x8f, 0xf2, 0xfe, 0xd7,
	0xed, 0xad, 0x8e, 0xeb, 0x34, 0xeb, 0x42, 0xe8, 0xb8, 0x5d, 0x1d, 0x7c, 0xfc, 0x16, 0x00, 0x00,
	0xff, 0xff, 0x3b, 0x5c, 0x34, 0xaa, 0xfc, 0x01, 0x00, 0x00,
}

func (m *HostingInquiry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostingInquiry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostingInquiry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxPricePerBlock != 0 {
		i = encodeVarintHostingInquiry(dAtA, i, uint64(m.MaxPricePerBlock))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintHostingInquiry(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.EndTime != 0 {
		i = encodeVarintHostingInquiry(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.EscrowAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHostingInquiry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.ReplicationRate != 0 {
		i = encodeVarintHostingInquiry(dAtA, i, uint64(m.ReplicationRate))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FileEntryCid) > 0 {
		i -= len(m.FileEntryCid)
		copy(dAtA[i:], m.FileEntryCid)
		i = encodeVarintHostingInquiry(dAtA, i, uint64(len(m.FileEntryCid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintHostingInquiry(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHostingInquiry(dAtA []byte, offset int, v uint64) int {
	offset -= sovHostingInquiry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostingInquiry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovHostingInquiry(uint64(m.Id))
	}
	l = len(m.FileEntryCid)
	if l > 0 {
		n += 1 + l + sovHostingInquiry(uint64(l))
	}
	if m.ReplicationRate != 0 {
		n += 1 + sovHostingInquiry(uint64(m.ReplicationRate))
	}
	l = m.EscrowAmount.Size()
	n += 1 + l + sovHostingInquiry(uint64(l))
	if m.EndTime != 0 {
		n += 1 + sovHostingInquiry(uint64(m.EndTime))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovHostingInquiry(uint64(l))
	}
	if m.MaxPricePerBlock != 0 {
		n += 1 + sovHostingInquiry(uint64(m.MaxPricePerBlock))
	}
	return n
}

func sovHostingInquiry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHostingInquiry(x uint64) (n int) {
	return sovHostingInquiry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostingInquiry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHostingInquiry
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
			return fmt.Errorf("proto: HostingInquiry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostingInquiry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingInquiry
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
				return fmt.Errorf("proto: wrong wireType = %d for field FileEntryCid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingInquiry
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
				return ErrInvalidLengthHostingInquiry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostingInquiry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileEntryCid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicationRate", wireType)
			}
			m.ReplicationRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingInquiry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicationRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingInquiry
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
				return ErrInvalidLengthHostingInquiry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostingInquiry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EscrowAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingInquiry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingInquiry
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
				return ErrInvalidLengthHostingInquiry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostingInquiry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPricePerBlock", wireType)
			}
			m.MaxPricePerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostingInquiry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPricePerBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHostingInquiry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHostingInquiry
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
func skipHostingInquiry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHostingInquiry
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
					return 0, ErrIntOverflowHostingInquiry
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
					return 0, ErrIntOverflowHostingInquiry
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
				return 0, ErrInvalidLengthHostingInquiry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHostingInquiry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHostingInquiry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHostingInquiry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHostingInquiry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHostingInquiry = fmt.Errorf("proto: unexpected end of group")
)
