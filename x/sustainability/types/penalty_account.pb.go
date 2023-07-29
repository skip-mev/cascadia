// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cascadia/sustainability/v1/penalty_account.proto

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

type PenaltyAccount struct {
	MultisigAddress string `protobuf:"bytes,1,opt,name=multisigAddress,proto3" json:"multisigAddress,omitempty"`
	Creator         string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *PenaltyAccount) Reset()         { *m = PenaltyAccount{} }
func (m *PenaltyAccount) String() string { return proto.CompactTextString(m) }
func (*PenaltyAccount) ProtoMessage()    {}
func (*PenaltyAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9d3aacef34c2d97, []int{0}
}
func (m *PenaltyAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PenaltyAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PenaltyAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PenaltyAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PenaltyAccount.Merge(m, src)
}
func (m *PenaltyAccount) XXX_Size() int {
	return m.Size()
}
func (m *PenaltyAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PenaltyAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PenaltyAccount proto.InternalMessageInfo

func (m *PenaltyAccount) GetMultisigAddress() string {
	if m != nil {
		return m.MultisigAddress
	}
	return ""
}

func (m *PenaltyAccount) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*PenaltyAccount)(nil), "cascadia.sustainability.v1.PenaltyAccount")
}

func init() {
	proto.RegisterFile("cascadia/sustainability/v1/penalty_account.proto", fileDescriptor_d9d3aacef34c2d97)
}

var fileDescriptor_d9d3aacef34c2d97 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0x4e, 0x2c, 0x4e,
	0x4e, 0x4c, 0xc9, 0x4c, 0xd4, 0x2f, 0x2e, 0x2d, 0x2e, 0x49, 0xcc, 0xcc, 0x4b, 0x4c, 0xca, 0xcc,
	0xc9, 0x2c, 0xa9, 0xd4, 0x2f, 0x33, 0xd4, 0x2f, 0x48, 0xcd, 0x4b, 0xcc, 0x29, 0xa9, 0x8c, 0x4f,
	0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x82, 0xe9,
	0xd0, 0x43, 0xd5, 0xa1, 0x57, 0x66, 0xa8, 0x14, 0xc2, 0xc5, 0x17, 0x00, 0xd1, 0xe4, 0x08, 0xd1,
	0x23, 0xa4, 0xc1, 0xc5, 0x9f, 0x5b, 0x9a, 0x53, 0x92, 0x59, 0x9c, 0x99, 0xee, 0x98, 0x92, 0x52,
	0x94, 0x5a, 0x5c, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2e, 0x2c, 0x24, 0xc1, 0xc5,
	0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0x04, 0x56, 0x01, 0xe3, 0x3a, 0x85, 0x9f,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6d, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xcc, 0x59, 0x69, 0xf9, 0xa5, 0x79, 0x29, 0x89, 0x25, 0x99,
	0xf9, 0x79, 0x70, 0x21, 0xfd, 0x0a, 0x74, 0xdf, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81,
	0x7d, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x75, 0x45, 0x2a, 0x05, 0x01, 0x00, 0x00,
}

func (m *PenaltyAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PenaltyAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PenaltyAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPenaltyAccount(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MultisigAddress) > 0 {
		i -= len(m.MultisigAddress)
		copy(dAtA[i:], m.MultisigAddress)
		i = encodeVarintPenaltyAccount(dAtA, i, uint64(len(m.MultisigAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPenaltyAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovPenaltyAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PenaltyAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MultisigAddress)
	if l > 0 {
		n += 1 + l + sovPenaltyAccount(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPenaltyAccount(uint64(l))
	}
	return n
}

func sovPenaltyAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPenaltyAccount(x uint64) (n int) {
	return sovPenaltyAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PenaltyAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPenaltyAccount
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
			return fmt.Errorf("proto: PenaltyAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PenaltyAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultisigAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPenaltyAccount
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
				return ErrInvalidLengthPenaltyAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPenaltyAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultisigAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPenaltyAccount
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
				return ErrInvalidLengthPenaltyAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPenaltyAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPenaltyAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPenaltyAccount
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
func skipPenaltyAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPenaltyAccount
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
					return 0, ErrIntOverflowPenaltyAccount
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
					return 0, ErrIntOverflowPenaltyAccount
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
				return 0, ErrInvalidLengthPenaltyAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPenaltyAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPenaltyAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPenaltyAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPenaltyAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPenaltyAccount = fmt.Errorf("proto: unexpected end of group")
)