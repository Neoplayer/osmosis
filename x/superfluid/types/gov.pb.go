// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/v1beta1/gov.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// SetSuperfluidAssetsProposal is a gov Content type to update the superfluid
// assets
type SetSuperfluidAssetsProposal struct {
	Title       string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Assets      []SuperfluidAsset `protobuf:"bytes,3,rep,name=assets,proto3" json:"assets"`
}

func (m *SetSuperfluidAssetsProposal) Reset()      { *m = SetSuperfluidAssetsProposal{} }
func (*SetSuperfluidAssetsProposal) ProtoMessage() {}
func (*SetSuperfluidAssetsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e37d6a8d0e42294, []int{0}
}
func (m *SetSuperfluidAssetsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetSuperfluidAssetsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetSuperfluidAssetsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetSuperfluidAssetsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSuperfluidAssetsProposal.Merge(m, src)
}
func (m *SetSuperfluidAssetsProposal) XXX_Size() int {
	return m.Size()
}
func (m *SetSuperfluidAssetsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSuperfluidAssetsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetSuperfluidAssetsProposal proto.InternalMessageInfo

// RemoveSuperfluidAssetsProposal is a gov Content type to remove the superfluid
// assets by denom
type RemoveSuperfluidAssetsProposal struct {
	Title                 string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description           string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SuperfluidAssetDenoms []string `protobuf:"bytes,3,rep,name=superfluid_asset_denoms,json=superfluidAssetDenoms,proto3" json:"superfluid_asset_denoms,omitempty"`
}

func (m *RemoveSuperfluidAssetsProposal) Reset()      { *m = RemoveSuperfluidAssetsProposal{} }
func (*RemoveSuperfluidAssetsProposal) ProtoMessage() {}
func (*RemoveSuperfluidAssetsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e37d6a8d0e42294, []int{1}
}
func (m *RemoveSuperfluidAssetsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveSuperfluidAssetsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveSuperfluidAssetsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveSuperfluidAssetsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveSuperfluidAssetsProposal.Merge(m, src)
}
func (m *RemoveSuperfluidAssetsProposal) XXX_Size() int {
	return m.Size()
}
func (m *RemoveSuperfluidAssetsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveSuperfluidAssetsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveSuperfluidAssetsProposal proto.InternalMessageInfo

// UpdateUnpoolWhiteListProposal is a gov Content type to update the
// allowed list of pool ids.
type UpdateUnpoolWhiteListProposal struct {
	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Ids         []uint64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	IsOverwrite bool     `protobuf:"varint,4,opt,name=is_overwrite,json=isOverwrite,proto3" json:"is_overwrite,omitempty"`
}

func (m *UpdateUnpoolWhiteListProposal) Reset()      { *m = UpdateUnpoolWhiteListProposal{} }
func (*UpdateUnpoolWhiteListProposal) ProtoMessage() {}
func (*UpdateUnpoolWhiteListProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e37d6a8d0e42294, []int{2}
}
func (m *UpdateUnpoolWhiteListProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateUnpoolWhiteListProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateUnpoolWhiteListProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateUnpoolWhiteListProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUnpoolWhiteListProposal.Merge(m, src)
}
func (m *UpdateUnpoolWhiteListProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateUnpoolWhiteListProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUnpoolWhiteListProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUnpoolWhiteListProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SetSuperfluidAssetsProposal)(nil), "osmosis.superfluid.v1beta1.SetSuperfluidAssetsProposal")
	proto.RegisterType((*RemoveSuperfluidAssetsProposal)(nil), "osmosis.superfluid.v1beta1.RemoveSuperfluidAssetsProposal")
	proto.RegisterType((*UpdateUnpoolWhiteListProposal)(nil), "osmosis.superfluid.v1beta1.UpdateUnpoolWhiteListProposal")
}

func init() {
	proto.RegisterFile("osmosis/superfluid/v1beta1/gov.proto", fileDescriptor_2e37d6a8d0e42294)
}

var fileDescriptor_2e37d6a8d0e42294 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0x4e, 0xc2, 0x50,
	0x14, 0xc6, 0x7b, 0x05, 0x89, 0x5c, 0x18, 0x4c, 0x83, 0xb1, 0xc1, 0xd8, 0x56, 0x70, 0x60, 0xb1,
	0x0d, 0x1a, 0x19, 0xdc, 0x20, 0x8e, 0x26, 0x92, 0x12, 0x62, 0xe2, 0x42, 0x5a, 0x7a, 0x2d, 0x37,
	0x69, 0x39, 0x4d, 0xef, 0xa5, 0xea, 0x1b, 0x38, 0x3a, 0x1a, 0x07, 0xc3, 0xe4, 0xb3, 0x30, 0x32,
	0x3a, 0x19, 0x03, 0x8b, 0x8f, 0x61, 0x68, 0x8b, 0xfc, 0x09, 0x9b, 0x6e, 0xf7, 0x9c, 0xf3, 0x7d,
	0xe7, 0xfb, 0x25, 0xf7, 0xe0, 0x63, 0x60, 0x1e, 0x30, 0xca, 0x74, 0x36, 0xf0, 0x49, 0x70, 0xe7,
	0x0e, 0xa8, 0xad, 0x87, 0x55, 0x8b, 0x70, 0xb3, 0xaa, 0x3b, 0x10, 0x6a, 0x7e, 0x00, 0x1c, 0xc4,
	0x62, 0xa2, 0xd2, 0x16, 0x2a, 0x2d, 0x51, 0x15, 0x0b, 0x0e, 0x38, 0x10, 0xc9, 0xf4, 0xd9, 0x2b,
	0x76, 0x14, 0xcb, 0x1b, 0xf6, 0x2e, 0x99, 0x23, 0x51, 0xe9, 0x1d, 0xe1, 0x83, 0x16, 0xe1, 0xad,
	0xdf, 0x7e, 0x9d, 0x31, 0xc2, 0x59, 0x33, 0x00, 0x1f, 0x98, 0xe9, 0x8a, 0x05, 0xbc, 0xcd, 0x29,
	0x77, 0x89, 0x84, 0x54, 0x54, 0xc9, 0x1a, 0x71, 0x21, 0xaa, 0x38, 0x67, 0x13, 0xd6, 0x0d, 0xa8,
	0xcf, 0x29, 0xf4, 0xa5, 0xad, 0x68, 0xb6, 0xdc, 0x12, 0xeb, 0x38, 0x63, 0x46, 0x9b, 0xa4, 0x94,
	0x9a, 0xaa, 0xe4, 0x4e, 0xcb, 0xda, 0x06, 0xfe, 0xb5, 0xd4, 0x46, 0x7a, 0xf4, 0xa9, 0x08, 0x46,
	0x62, 0xbc, 0xc8, 0x3f, 0x0d, 0x15, 0xe1, 0x65, 0xa8, 0x08, 0xdf, 0x43, 0x05, 0x95, 0xde, 0x10,
	0x96, 0x0d, 0xe2, 0x41, 0x48, 0xfe, 0x9d, 0xb5, 0x86, 0xf7, 0x17, 0x50, 0x9d, 0x28, 0xbd, 0x63,
	0x93, 0x3e, 0x78, 0x31, 0x7c, 0xd6, 0xd8, 0x63, 0xab, 0x91, 0x97, 0xd1, 0x70, 0x0d, 0xf0, 0x15,
	0xe1, 0xc3, 0xb6, 0x6f, 0x9b, 0x9c, 0xb4, 0xfb, 0x3e, 0x80, 0x7b, 0xd3, 0xa3, 0x9c, 0x5c, 0x51,
	0xc6, 0xff, 0xcc, 0xb7, 0x8b, 0x53, 0xd4, 0x8e, 0x59, 0xd2, 0xc6, 0xec, 0x29, 0x1e, 0xe1, 0x3c,
	0x65, 0x1d, 0x08, 0x49, 0x70, 0x1f, 0x50, 0x4e, 0xa4, 0xb4, 0x8a, 0x2a, 0x3b, 0x46, 0x8e, 0xb2,
	0xeb, 0x79, 0x6b, 0x15, 0xae, 0xd1, 0x1c, 0x4d, 0x64, 0x34, 0x9e, 0xc8, 0xe8, 0x6b, 0x22, 0xa3,
	0xe7, 0xa9, 0x2c, 0x8c, 0xa7, 0xb2, 0xf0, 0x31, 0x95, 0x85, 0xdb, 0x9a, 0x43, 0x79, 0x6f, 0x60,
	0x69, 0x5d, 0xf0, 0xf4, 0xe4, 0x8b, 0x4e, 0x5c, 0xd3, 0x62, 0xf3, 0x42, 0x0f, 0xab, 0xe7, 0xfa,
	0xc3, 0xf2, 0x0d, 0xf1, 0x47, 0x9f, 0x30, 0x2b, 0x13, 0xdd, 0xcf, 0xd9, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf1, 0xe3, 0x6a, 0xf5, 0xbe, 0x02, 0x00, 0x00,
}

func (this *SetSuperfluidAssetsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetSuperfluidAssetsProposal)
	if !ok {
		that2, ok := that.(SetSuperfluidAssetsProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Assets) != len(that1.Assets) {
		return false
	}
	for i := range this.Assets {
		if !this.Assets[i].Equal(&that1.Assets[i]) {
			return false
		}
	}
	return true
}
func (this *RemoveSuperfluidAssetsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RemoveSuperfluidAssetsProposal)
	if !ok {
		that2, ok := that.(RemoveSuperfluidAssetsProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.SuperfluidAssetDenoms) != len(that1.SuperfluidAssetDenoms) {
		return false
	}
	for i := range this.SuperfluidAssetDenoms {
		if this.SuperfluidAssetDenoms[i] != that1.SuperfluidAssetDenoms[i] {
			return false
		}
	}
	return true
}
func (this *UpdateUnpoolWhiteListProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateUnpoolWhiteListProposal)
	if !ok {
		that2, ok := that.(UpdateUnpoolWhiteListProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Ids) != len(that1.Ids) {
		return false
	}
	for i := range this.Ids {
		if this.Ids[i] != that1.Ids[i] {
			return false
		}
	}
	if this.IsOverwrite != that1.IsOverwrite {
		return false
	}
	return true
}
func (m *SetSuperfluidAssetsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetSuperfluidAssetsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetSuperfluidAssetsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoveSuperfluidAssetsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveSuperfluidAssetsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveSuperfluidAssetsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SuperfluidAssetDenoms) > 0 {
		for iNdEx := len(m.SuperfluidAssetDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SuperfluidAssetDenoms[iNdEx])
			copy(dAtA[i:], m.SuperfluidAssetDenoms[iNdEx])
			i = encodeVarintGov(dAtA, i, uint64(len(m.SuperfluidAssetDenoms[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateUnpoolWhiteListProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateUnpoolWhiteListProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateUnpoolWhiteListProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsOverwrite {
		i--
		if m.IsOverwrite {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Ids) > 0 {
		dAtA2 := make([]byte, len(m.Ids)*10)
		var j1 int
		for _, num := range m.Ids {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGov(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SetSuperfluidAssetsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *RemoveSuperfluidAssetsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.SuperfluidAssetDenoms) > 0 {
		for _, s := range m.SuperfluidAssetDenoms {
			l = len(s)
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *UpdateUnpoolWhiteListProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovGov(uint64(e))
		}
		n += 1 + sovGov(uint64(l)) + l
	}
	if m.IsOverwrite {
		n += 2
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SetSuperfluidAssetsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: SetSuperfluidAssetsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetSuperfluidAssetsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, SuperfluidAsset{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *RemoveSuperfluidAssetsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: RemoveSuperfluidAssetsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveSuperfluidAssetsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuperfluidAssetDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SuperfluidAssetDenoms = append(m.SuperfluidAssetDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *UpdateUnpoolWhiteListProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateUnpoolWhiteListProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateUnpoolWhiteListProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGov
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGov
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGov
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsOverwrite", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
			m.IsOverwrite = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
