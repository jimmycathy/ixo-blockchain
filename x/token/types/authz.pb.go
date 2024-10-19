// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/token/v1beta1/authz.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
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

type MintAuthorization struct {
	// address of minter
	Minter      string             `protobuf:"bytes,1,opt,name=minter,proto3" json:"minter,omitempty"`
	Constraints []*MintConstraints `protobuf:"bytes,2,rep,name=constraints,proto3" json:"constraints,omitempty"`
}

func (m *MintAuthorization) Reset()         { *m = MintAuthorization{} }
func (m *MintAuthorization) String() string { return proto.CompactTextString(m) }
func (*MintAuthorization) ProtoMessage()    {}
func (*MintAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca6ff8f08fc36db, []int{0}
}
func (m *MintAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintAuthorization.Merge(m, src)
}
func (m *MintAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *MintAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_MintAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_MintAuthorization proto.InternalMessageInfo

func (m *MintAuthorization) GetMinter() string {
	if m != nil {
		return m.Minter
	}
	return ""
}

func (m *MintAuthorization) GetConstraints() []*MintConstraints {
	if m != nil {
		return m.Constraints
	}
	return nil
}

type MintConstraints struct {
	ContractAddress string                 `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Amount          cosmossdk_io_math.Uint `protobuf:"bytes,2,opt,name=amount,proto3,customtype=cosmossdk.io/math.Uint" json:"amount"`
	// name is the token name, which must be unique (namespace), will be verified
	// against Token name provided on msgCreateToken
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// index is the unique identifier hexstring that identifies the token
	Index string `protobuf:"bytes,4,opt,name=index,proto3" json:"index,omitempty"`
	// did of collection (eg Supamoto Malawi)
	Collection string `protobuf:"bytes,5,opt,name=collection,proto3" json:"collection,omitempty"`
	// tokenData is the linkedResources added to tokenMetadata when queried eg
	// (credential link ***.ipfs)
	TokenData []*TokenData `protobuf:"bytes,6,rep,name=tokenData,proto3" json:"tokenData,omitempty"`
}

func (m *MintConstraints) Reset()         { *m = MintConstraints{} }
func (m *MintConstraints) String() string { return proto.CompactTextString(m) }
func (*MintConstraints) ProtoMessage()    {}
func (*MintConstraints) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca6ff8f08fc36db, []int{1}
}
func (m *MintConstraints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintConstraints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintConstraints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintConstraints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintConstraints.Merge(m, src)
}
func (m *MintConstraints) XXX_Size() int {
	return m.Size()
}
func (m *MintConstraints) XXX_DiscardUnknown() {
	xxx_messageInfo_MintConstraints.DiscardUnknown(m)
}

var xxx_messageInfo_MintConstraints proto.InternalMessageInfo

func (m *MintConstraints) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MintConstraints) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MintConstraints) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *MintConstraints) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *MintConstraints) GetTokenData() []*TokenData {
	if m != nil {
		return m.TokenData
	}
	return nil
}

func init() {
	proto.RegisterType((*MintAuthorization)(nil), "ixo.token.v1beta1.MintAuthorization")
	proto.RegisterType((*MintConstraints)(nil), "ixo.token.v1beta1.MintConstraints")
}

func init() { proto.RegisterFile("ixo/token/v1beta1/authz.proto", fileDescriptor_cca6ff8f08fc36db) }

var fileDescriptor_cca6ff8f08fc36db = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xb1, 0x4e, 0xdc, 0x40,
	0x10, 0x3d, 0x1f, 0x70, 0x12, 0x7b, 0x05, 0x61, 0x85, 0x90, 0x83, 0x12, 0x83, 0xae, 0x88, 0x48,
	0x71, 0x5e, 0x11, 0x9a, 0x88, 0x0e, 0x42, 0xca, 0x34, 0x56, 0xd2, 0xa4, 0x41, 0xeb, 0xf5, 0xe6,
	0x3c, 0xba, 0xf3, 0x0e, 0xf2, 0x8e, 0x91, 0xc3, 0x1f, 0xa4, 0x4b, 0x9b, 0xff, 0xe0, 0x23, 0x50,
	0x2a, 0x94, 0x2a, 0x4a, 0x81, 0xa2, 0xbb, 0x1f, 0x89, 0xbc, 0xbb, 0x17, 0x48, 0xa0, 0xf3, 0x9b,
	0xf7, 0xe6, 0xcd, 0xf3, 0xd3, 0xb2, 0xe7, 0xd0, 0xa2, 0x20, 0x9c, 0x6a, 0x23, 0x2e, 0x0e, 0x72,
	0x4d, 0xf2, 0x40, 0xc8, 0x86, 0xca, 0xcb, 0xf4, 0xbc, 0x46, 0x42, 0xbe, 0x09, 0x2d, 0xa6, 0x8e,
	0x4e, 0x03, 0xbd, 0xb3, 0x35, 0xc1, 0x09, 0x3a, 0x56, 0x74, 0x5f, 0x5e, 0xb8, 0xf3, 0xb4, 0xf3,
	0x01, 0x28, 0xfe, 0xba, 0x00, 0x14, 0x81, 0x7a, 0xe4, 0x84, 0x77, 0x0c, 0x9b, 0x0a, 0x6d, 0x85,
	0xf6, 0xcc, 0x5b, 0x7a, 0xe0, 0xa9, 0xd1, 0xb7, 0x88, 0x6d, 0xbe, 0x03, 0x43, 0xc7, 0x0d, 0x95,
	0x58, 0xc3, 0xa5, 0x24, 0x40, 0xc3, 0xb7, 0xd9, 0xa0, 0x02, 0x43, 0xba, 0x8e, 0xa3, 0xbd, 0x68,
	0x7f, 0x3d, 0x0b, 0x88, 0x9f, 0xb2, 0xa1, 0x42, 0x63, 0xa9, 0x96, 0x60, 0xc8, 0xc6, 0xfd, 0xbd,
	0x95, 0xfd, 0xe1, 0xab, 0x51, 0xfa, 0xe0, 0x0f, 0xd2, 0xce, 0xf2, 0xcd, 0x9d, 0x32, 0xbb, 0xbf,
	0x76, 0xf4, 0xe2, 0xfb, 0xd5, 0x78, 0x14, 0x52, 0xf8, 0x26, 0x96, 0x6b, 0xff, 0xa4, 0x18, 0x7d,
	0xe9, 0xb3, 0x8d, 0xff, 0x8c, 0xf8, 0x4b, 0xf6, 0x44, 0xa1, 0xa1, 0x5a, 0x2a, 0x3a, 0x93, 0x45,
	0x51, 0x6b, 0x6b, 0x43, 0xc6, 0x8d, 0xe5, 0xfc, 0xd8, 0x8f, 0xf9, 0x5b, 0x36, 0x90, 0x15, 0x36,
	0x86, 0xe2, 0x7e, 0x27, 0x38, 0x19, 0x5f, 0xdf, 0xee, 0xf6, 0x7e, 0xdd, 0xee, 0x6e, 0xfb, 0xd3,
	0xb6, 0x98, 0xa6, 0x80, 0xa2, 0x92, 0x54, 0xa6, 0x1f, 0xc0, 0xd0, 0x8f, 0xab, 0xf1, 0x30, 0x84,
	0xea, 0x60, 0x16, 0x96, 0x39, 0x67, 0xab, 0x46, 0x56, 0x3a, 0x5e, 0x71, 0x57, 0xdc, 0x37, 0xdf,
	0x62, 0x6b, 0x60, 0x0a, 0xdd, 0xc6, 0xab, 0x6e, 0xe8, 0x01, 0x4f, 0x18, 0x53, 0x38, 0x9b, 0x69,
	0xd5, 0xa5, 0x8f, 0xd7, 0x1c, 0x75, 0x6f, 0xc2, 0x8f, 0xd8, 0xba, 0x6b, 0xe9, 0x54, 0x92, 0x8c,
	0x07, 0xae, 0xbb, 0x67, 0x8f, 0x74, 0xf7, 0x7e, 0xa9, 0xc9, 0xee, 0xe4, 0x27, 0xd9, 0xf5, 0x3c,
	0x89, 0x6e, 0xe6, 0x49, 0xf4, 0x7b, 0x9e, 0x44, 0x5f, 0x17, 0x49, 0xef, 0x66, 0x91, 0xf4, 0x7e,
	0x2e, 0x92, 0xde, 0xc7, 0xd7, 0x13, 0xa0, 0xb2, 0xc9, 0x53, 0x85, 0x95, 0x80, 0x16, 0x3f, 0x61,
	0x63, 0x0a, 0xd7, 0x5f, 0x87, 0xc6, 0xf9, 0x0c, 0xd5, 0x54, 0x95, 0x12, 0x8c, 0xb8, 0x38, 0x14,
	0x6d, 0x78, 0x24, 0xf4, 0xf9, 0x5c, 0xdb, 0x7c, 0xe0, 0x9e, 0xc0, 0xe1, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x34, 0xae, 0xad, 0x33, 0xa1, 0x02, 0x00, 0x00,
}

func (m *MintAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Constraints) > 0 {
		for iNdEx := len(m.Constraints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Constraints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Minter) > 0 {
		i -= len(m.Minter)
		copy(dAtA[i:], m.Minter)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.Minter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MintConstraints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintConstraints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintConstraints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenData) > 0 {
		for iNdEx := len(m.TokenData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Collection) > 0 {
		i -= len(m.Collection)
		copy(dAtA[i:], m.Collection)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.Collection)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuthz(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Minter)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	if len(m.Constraints) > 0 {
		for _, e := range m.Constraints {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func (m *MintConstraints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAuthz(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	l = len(m.Collection)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	if len(m.TokenData) > 0 {
		for _, e := range m.TokenData {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: MintAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constraints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Constraints = append(m.Constraints, &MintConstraints{})
			if err := m.Constraints[len(m.Constraints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *MintConstraints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: MintConstraints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintConstraints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collection", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collection = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenData = append(m.TokenData, &TokenData{})
			if err := m.TokenData[len(m.TokenData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
