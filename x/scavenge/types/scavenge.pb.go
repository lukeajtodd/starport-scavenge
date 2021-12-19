// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scavenge/scavenge.proto

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

type Scavenge struct {
	Index        string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	SolutionHash string `protobuf:"bytes,2,opt,name=solutionHash,proto3" json:"solutionHash,omitempty"`
	Solution     string `protobuf:"bytes,3,opt,name=solution,proto3" json:"solution,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Reward       string `protobuf:"bytes,5,opt,name=reward,proto3" json:"reward,omitempty"`
	Scavenger    string `protobuf:"bytes,6,opt,name=scavenger,proto3" json:"scavenger,omitempty"`
}

func (m *Scavenge) Reset()         { *m = Scavenge{} }
func (m *Scavenge) String() string { return proto.CompactTextString(m) }
func (*Scavenge) ProtoMessage()    {}
func (*Scavenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ca21a8cb5d41f9, []int{0}
}
func (m *Scavenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scavenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scavenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scavenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scavenge.Merge(m, src)
}
func (m *Scavenge) XXX_Size() int {
	return m.Size()
}
func (m *Scavenge) XXX_DiscardUnknown() {
	xxx_messageInfo_Scavenge.DiscardUnknown(m)
}

var xxx_messageInfo_Scavenge proto.InternalMessageInfo

func (m *Scavenge) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Scavenge) GetSolutionHash() string {
	if m != nil {
		return m.SolutionHash
	}
	return ""
}

func (m *Scavenge) GetSolution() string {
	if m != nil {
		return m.Solution
	}
	return ""
}

func (m *Scavenge) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Scavenge) GetReward() string {
	if m != nil {
		return m.Reward
	}
	return ""
}

func (m *Scavenge) GetScavenger() string {
	if m != nil {
		return m.Scavenger
	}
	return ""
}

func init() {
	proto.RegisterType((*Scavenge)(nil), "lukeajtodd.scavenge.scavenge.Scavenge")
}

func init() { proto.RegisterFile("scavenge/scavenge.proto", fileDescriptor_80ca21a8cb5d41f9) }

var fileDescriptor_80ca21a8cb5d41f9 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x4e, 0x2c,
	0x4b, 0xcd, 0x4b, 0x4f, 0xd5, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x64, 0x72,
	0x4a, 0xb3, 0x53, 0x13, 0xb3, 0x4a, 0xf2, 0x53, 0x52, 0xf4, 0xe0, 0x52, 0x30, 0x86, 0xd2, 0x0e,
	0x46, 0x2e, 0x8e, 0x60, 0x28, 0x47, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x11, 0x52, 0xe2, 0xe2, 0x29, 0xce, 0xcf, 0x29, 0x2d,
	0xc9, 0xcc, 0xcf, 0xf3, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x02, 0x4b, 0xa2, 0x88, 0x09, 0x49, 0x71,
	0x71, 0xc0, 0xf8, 0x12, 0xcc, 0x60, 0x79, 0x38, 0x5f, 0x48, 0x81, 0x8b, 0x3b, 0x25, 0xb5, 0x38,
	0xb9, 0x28, 0xb3, 0x00, 0x2c, 0xcd, 0x02, 0x96, 0x46, 0x16, 0x12, 0x12, 0xe3, 0x62, 0x2b, 0x4a,
	0x2d, 0x4f, 0x2c, 0x4a, 0x91, 0x60, 0x05, 0x4b, 0x42, 0x79, 0x42, 0x32, 0x5c, 0x9c, 0x30, 0x87,
	0x16, 0x49, 0xb0, 0x81, 0xa5, 0x10, 0x02, 0x4e, 0x9e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xa5, 0x9f, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f,
	0xf0, 0x3d, 0x3c, 0x60, 0xf4, 0x2b, 0x10, 0xcc, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70,
	0x50, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x00, 0xbd, 0x9d, 0x45, 0x01, 0x00, 0x00,
}

func (m *Scavenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scavenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scavenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Scavenger) > 0 {
		i -= len(m.Scavenger)
		copy(dAtA[i:], m.Scavenger)
		i = encodeVarintScavenge(dAtA, i, uint64(len(m.Scavenger)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Reward) > 0 {
		i -= len(m.Reward)
		copy(dAtA[i:], m.Reward)
		i = encodeVarintScavenge(dAtA, i, uint64(len(m.Reward)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintScavenge(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Solution) > 0 {
		i -= len(m.Solution)
		copy(dAtA[i:], m.Solution)
		i = encodeVarintScavenge(dAtA, i, uint64(len(m.Solution)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SolutionHash) > 0 {
		i -= len(m.SolutionHash)
		copy(dAtA[i:], m.SolutionHash)
		i = encodeVarintScavenge(dAtA, i, uint64(len(m.SolutionHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintScavenge(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScavenge(dAtA []byte, offset int, v uint64) int {
	offset -= sovScavenge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Scavenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovScavenge(uint64(l))
	}
	l = len(m.SolutionHash)
	if l > 0 {
		n += 1 + l + sovScavenge(uint64(l))
	}
	l = len(m.Solution)
	if l > 0 {
		n += 1 + l + sovScavenge(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovScavenge(uint64(l))
	}
	l = len(m.Reward)
	if l > 0 {
		n += 1 + l + sovScavenge(uint64(l))
	}
	l = len(m.Scavenger)
	if l > 0 {
		n += 1 + l + sovScavenge(uint64(l))
	}
	return n
}

func sovScavenge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScavenge(x uint64) (n int) {
	return sovScavenge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Scavenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScavenge
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
			return fmt.Errorf("proto: Scavenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scavenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScavenge
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
				return ErrInvalidLengthScavenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScavenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SolutionHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScavenge
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
				return ErrInvalidLengthScavenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScavenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SolutionHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Solution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScavenge
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
				return ErrInvalidLengthScavenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScavenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Solution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScavenge
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
				return ErrInvalidLengthScavenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScavenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScavenge
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
				return ErrInvalidLengthScavenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScavenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reward = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scavenger", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScavenge
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
				return ErrInvalidLengthScavenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScavenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scavenger = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScavenge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScavenge
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
func skipScavenge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScavenge
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
					return 0, ErrIntOverflowScavenge
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
					return 0, ErrIntOverflowScavenge
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
				return 0, ErrInvalidLengthScavenge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScavenge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScavenge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScavenge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScavenge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScavenge = fmt.Errorf("proto: unexpected end of group")
)
