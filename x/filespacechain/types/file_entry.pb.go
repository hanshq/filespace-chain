// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filespacechain/filespacechain/file_entry.proto

package types

import (
	fmt "fmt"
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

type FileEntry struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cid       string `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	RootCid   string `protobuf:"bytes,3,opt,name=rootCid,proto3" json:"rootCid,omitempty"`
	ParentCid string `protobuf:"bytes,4,opt,name=parentCid,proto3" json:"parentCid,omitempty"`
	MetaData  string `protobuf:"bytes,5,opt,name=metaData,proto3" json:"metaData,omitempty"`
	FileSize  uint64 `protobuf:"varint,6,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	Creator   string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *FileEntry) Reset()         { *m = FileEntry{} }
func (m *FileEntry) String() string { return proto.CompactTextString(m) }
func (*FileEntry) ProtoMessage()    {}
func (*FileEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5676b2ee239130f, []int{0}
}
func (m *FileEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileEntry.Merge(m, src)
}
func (m *FileEntry) XXX_Size() int {
	return m.Size()
}
func (m *FileEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FileEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FileEntry proto.InternalMessageInfo

func (m *FileEntry) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FileEntry) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *FileEntry) GetRootCid() string {
	if m != nil {
		return m.RootCid
	}
	return ""
}

func (m *FileEntry) GetParentCid() string {
	if m != nil {
		return m.ParentCid
	}
	return ""
}

func (m *FileEntry) GetMetaData() string {
	if m != nil {
		return m.MetaData
	}
	return ""
}

func (m *FileEntry) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *FileEntry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*FileEntry)(nil), "filespacechain.filespacechain.FileEntry")
}

func init() {
	proto.RegisterFile("filespacechain/filespacechain/file_entry.proto", fileDescriptor_f5676b2ee239130f)
}

var fileDescriptor_f5676b2ee239130f = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcb, 0xcc, 0x49,
	0x2d, 0x2e, 0x48, 0x4c, 0x4e, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xc7, 0xc2, 0x8d, 0x4f, 0xcd,
	0x2b, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x45, 0x55, 0x80, 0xa6, 0x5d,
	0x69, 0x3b, 0x23, 0x17, 0xa7, 0x5b, 0x66, 0x4e, 0xaa, 0x2b, 0x48, 0x8b, 0x10, 0x1f, 0x17, 0x53,
	0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x00, 0x17, 0x73,
	0x72, 0x66, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc1, 0xc5, 0x5e,
	0x94, 0x9f, 0x5f, 0xe2, 0x9c, 0x99, 0x22, 0xc1, 0x0c, 0x16, 0x85, 0x71, 0x85, 0x64, 0xb8, 0x38,
	0x0b, 0x12, 0x8b, 0x52, 0xf3, 0xc0, 0x72, 0x2c, 0x60, 0x39, 0x84, 0x80, 0x90, 0x14, 0x17, 0x47,
	0x6e, 0x6a, 0x49, 0xa2, 0x4b, 0x62, 0x49, 0xa2, 0x04, 0x2b, 0x58, 0x12, 0xce, 0x07, 0xc9, 0x81,
	0x5c, 0x15, 0x9c, 0x59, 0x95, 0x2a, 0xc1, 0x06, 0xb6, 0x1b, 0xce, 0x07, 0xd9, 0x97, 0x5c, 0x94,
	0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0x0e, 0xb1, 0x0f, 0xca, 0x75, 0x0a, 0x3a, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58,
	0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x8b, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0xfd, 0x8c, 0xc4, 0xbc, 0xe2, 0x8c, 0x42, 0x44, 0x28, 0xe9, 0x42, 0x82, 0xa9, 0x02,
	0x3d, 0xdc, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x61, 0x66, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0xf9, 0x0c, 0xd6, 0x65, 0x01, 0x00, 0x00,
}

func (m *FileEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintFileEntry(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if m.FileSize != 0 {
		i = encodeVarintFileEntry(dAtA, i, uint64(m.FileSize))
		i--
		dAtA[i] = 0x30
	}
	if len(m.MetaData) > 0 {
		i -= len(m.MetaData)
		copy(dAtA[i:], m.MetaData)
		i = encodeVarintFileEntry(dAtA, i, uint64(len(m.MetaData)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ParentCid) > 0 {
		i -= len(m.ParentCid)
		copy(dAtA[i:], m.ParentCid)
		i = encodeVarintFileEntry(dAtA, i, uint64(len(m.ParentCid)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RootCid) > 0 {
		i -= len(m.RootCid)
		copy(dAtA[i:], m.RootCid)
		i = encodeVarintFileEntry(dAtA, i, uint64(len(m.RootCid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintFileEntry(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintFileEntry(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFileEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovFileEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FileEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovFileEntry(uint64(m.Id))
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovFileEntry(uint64(l))
	}
	l = len(m.RootCid)
	if l > 0 {
		n += 1 + l + sovFileEntry(uint64(l))
	}
	l = len(m.ParentCid)
	if l > 0 {
		n += 1 + l + sovFileEntry(uint64(l))
	}
	l = len(m.MetaData)
	if l > 0 {
		n += 1 + l + sovFileEntry(uint64(l))
	}
	if m.FileSize != 0 {
		n += 1 + sovFileEntry(uint64(m.FileSize))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovFileEntry(uint64(l))
	}
	return n
}

func sovFileEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFileEntry(x uint64) (n int) {
	return sovFileEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileEntry
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
			return fmt.Errorf("proto: FileEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileEntry
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
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileEntry
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
				return ErrInvalidLengthFileEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootCid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileEntry
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
				return ErrInvalidLengthFileEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootCid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentCid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileEntry
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
				return ErrInvalidLengthFileEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentCid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileEntry
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
				return ErrInvalidLengthFileEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetaData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileSize", wireType)
			}
			m.FileSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileEntry
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
				return ErrInvalidLengthFileEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFileEntry
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
func skipFileEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileEntry
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
					return 0, ErrIntOverflowFileEntry
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
					return 0, ErrIntOverflowFileEntry
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
				return 0, ErrInvalidLengthFileEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFileEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFileEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFileEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFileEntry = fmt.Errorf("proto: unexpected end of group")
)
