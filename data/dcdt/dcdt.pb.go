// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/dcdt.proto

package dcdt

import (
	bytes "bytes"
	fmt "fmt"
	github_com_kalyan3104_k_chain_core_go_data "github.com/kalyan3104/k-chain-core-go/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// DCDigitalToken holds the data for a kalyan Core Token transaction
type DCDigitalToken struct {
	Type          uint32        `protobuf:"varint,1,opt,name=Type,proto3" json:"Type"`
	Value         *math_big.Int `protobuf:"bytes,2,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/kalyan3104/k-chain-core-go/data.BigIntCaster" json:"Value"`
	Properties    []byte        `protobuf:"bytes,3,opt,name=Properties,proto3" json:"Properties"`
	TokenMetaData *MetaData     `protobuf:"bytes,4,opt,name=TokenMetaData,proto3" json:"MetaData"`
	Reserved      []byte        `protobuf:"bytes,5,opt,name=Reserved,proto3" json:"Reserved"`
}

func (m *DCDigitalToken) Reset()      { *m = DCDigitalToken{} }
func (*DCDigitalToken) ProtoMessage() {}
func (*DCDigitalToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_1350fbac2742f2f2, []int{0}
}
func (m *DCDigitalToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DCDigitalToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DCDigitalToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCDigitalToken.Merge(m, src)
}
func (m *DCDigitalToken) XXX_Size() int {
	return m.Size()
}
func (m *DCDigitalToken) XXX_DiscardUnknown() {
	xxx_messageInfo_DCDigitalToken.DiscardUnknown(m)
}

var xxx_messageInfo_DCDigitalToken proto.InternalMessageInfo

func (m *DCDigitalToken) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *DCDigitalToken) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DCDigitalToken) GetProperties() []byte {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *DCDigitalToken) GetTokenMetaData() *MetaData {
	if m != nil {
		return m.TokenMetaData
	}
	return nil
}

func (m *DCDigitalToken) GetReserved() []byte {
	if m != nil {
		return m.Reserved
	}
	return nil
}

// DCDTRoles holds the roles for a given token and the given address
type DCDTRoles struct {
	Roles [][]byte `protobuf:"bytes,1,rep,name=Roles,proto3" json:"roles"`
}

func (m *DCDTRoles) Reset()      { *m = DCDTRoles{} }
func (*DCDTRoles) ProtoMessage() {}
func (*DCDTRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_1350fbac2742f2f2, []int{1}
}
func (m *DCDTRoles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DCDTRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DCDTRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCDTRoles.Merge(m, src)
}
func (m *DCDTRoles) XXX_Size() int {
	return m.Size()
}
func (m *DCDTRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_DCDTRoles.DiscardUnknown(m)
}

var xxx_messageInfo_DCDTRoles proto.InternalMessageInfo

func (m *DCDTRoles) GetRoles() [][]byte {
	if m != nil {
		return m.Roles
	}
	return nil
}

// MetaData hold the metadata structure for the DCDT token
type MetaData struct {
	Nonce      uint64   `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce"`
	Name       []byte   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name"`
	Creator    []byte   `protobuf:"bytes,3,opt,name=Creator,proto3" json:"Creator"`
	Royalties  uint32   `protobuf:"varint,4,opt,name=Royalties,proto3" json:"Royalties"`
	Hash       []byte   `protobuf:"bytes,5,opt,name=Hash,proto3" json:"Hash"`
	URIs       [][]byte `protobuf:"bytes,6,rep,name=URIs,proto3" json:"URIs"`
	Attributes []byte   `protobuf:"bytes,7,opt,name=Attributes,proto3" json:"Attributes"`
}

func (m *MetaData) Reset()      { *m = MetaData{} }
func (*MetaData) ProtoMessage() {}
func (*MetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1350fbac2742f2f2, []int{2}
}
func (m *MetaData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaData.Merge(m, src)
}
func (m *MetaData) XXX_Size() int {
	return m.Size()
}
func (m *MetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MetaData proto.InternalMessageInfo

func (m *MetaData) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *MetaData) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MetaData) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MetaData) GetRoyalties() uint32 {
	if m != nil {
		return m.Royalties
	}
	return 0
}

func (m *MetaData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *MetaData) GetURIs() [][]byte {
	if m != nil {
		return m.URIs
	}
	return nil
}

func (m *MetaData) GetAttributes() []byte {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*DCDigitalToken)(nil), "protoBuiltInFunctions.DCDigitalToken")
	proto.RegisterType((*DCDTRoles)(nil), "protoBuiltInFunctions.DCDTRoles")
	proto.RegisterType((*MetaData)(nil), "protoBuiltInFunctions.MetaData")
}

func init() { proto.RegisterFile("proto/dcdt.proto", fileDescriptor_1350fbac2742f2f2) }

var fileDescriptor_1350fbac2742f2f2 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xc1, 0x8b, 0xd3, 0x4e,
	0x14, 0xce, 0xec, 0xb6, 0xbb, 0xed, 0x6c, 0xbb, 0xfc, 0x08, 0xfc, 0x20, 0x88, 0x4c, 0x4a, 0x41,
	0x28, 0x68, 0x13, 0xd0, 0xa3, 0x20, 0x98, 0x06, 0xb5, 0x07, 0x57, 0x19, 0xab, 0x07, 0x6f, 0xd3,
	0x74, 0x4c, 0x47, 0xdb, 0x4c, 0x99, 0x4c, 0x84, 0xbd, 0x79, 0xf5, 0xe6, 0x9f, 0x21, 0xfe, 0x25,
	0x7a, 0xeb, 0xb1, 0xa7, 0x68, 0xd3, 0x8b, 0xe4, 0xb4, 0x7f, 0x82, 0xcc, 0xcb, 0x66, 0x5b, 0xc1,
	0xcb, 0xcb, 0xf7, 0x7d, 0xf3, 0x78, 0x6f, 0xe6, 0xfb, 0x08, 0xfe, 0x6f, 0xa5, 0xa4, 0x96, 0xfe,
	0x2c, 0x9a, 0x69, 0x0f, 0xa0, 0xfd, 0x3f, 0x7c, 0x82, 0x4c, 0x2c, 0xf4, 0x38, 0x79, 0x92, 0x25,
	0x91, 0x16, 0x32, 0x49, 0x6f, 0x0d, 0x63, 0xa1, 0xe7, 0xd9, 0xd4, 0x8b, 0xe4, 0xd2, 0x8f, 0x65,
	0x2c, 0x7d, 0x68, 0x9b, 0x66, 0xef, 0x80, 0x55, 0x53, 0x0c, 0xaa, 0xa6, 0xf4, 0x7f, 0x1c, 0xe1,
	0xf3, 0x70, 0x14, 0x8a, 0x58, 0x68, 0xb6, 0x98, 0xc8, 0x0f, 0x3c, 0xb1, 0x6f, 0xe3, 0xc6, 0xe4,
	0x72, 0xc5, 0x1d, 0xd4, 0x43, 0x83, 0x6e, 0xd0, 0x2a, 0x73, 0x17, 0x38, 0x85, 0x6a, 0xbf, 0xc7,
	0xcd, 0x37, 0x6c, 0x91, 0x71, 0xe7, 0xa8, 0x87, 0x06, 0x9d, 0x60, 0x52, 0xe6, 0x6e, 0x25, 0x7c,
	0xfb, 0xe9, 0x3e, 0x5d, 0x32, 0x3d, 0xf7, 0xa7, 0x22, 0xf6, 0xc6, 0x89, 0x7e, 0x78, 0x70, 0x91,
	0x70, 0xce, 0x94, 0xd0, 0x4a, 0xbc, 0x48, 0xb8, 0x3f, 0x53, 0x7a, 0x18, 0xcd, 0x99, 0x48, 0x86,
	0x91, 0x54, 0x7c, 0x18, 0x4b, 0x7f, 0xc6, 0x34, 0xf3, 0x02, 0x11, 0x8f, 0x13, 0x3d, 0x62, 0xa9,
	0xe6, 0x8a, 0x56, 0x13, 0x6d, 0x0f, 0xe3, 0x97, 0x4a, 0xae, 0xb8, 0xd2, 0x82, 0xa7, 0xce, 0x31,
	0x2c, 0x3c, 0x2f, 0x73, 0xf7, 0x40, 0xa5, 0x07, 0xd8, 0x7e, 0x85, 0xbb, 0xf0, 0x84, 0xe7, 0x5c,
	0xb3, 0x90, 0x69, 0xe6, 0x34, 0x7a, 0x68, 0x70, 0x76, 0xdf, 0xf5, 0xfe, 0x69, 0x95, 0x57, 0xb7,
	0x05, 0x9d, 0x32, 0x77, 0x5b, 0x35, 0xa3, 0x7f, 0xcf, 0xb0, 0x07, 0xb8, 0x45, 0x79, 0xca, 0xd5,
	0x47, 0x3e, 0x73, 0x9a, 0x70, 0x05, 0x68, 0xaf, 0x35, 0x7a, 0x83, 0xfa, 0xf7, 0x70, 0x3b, 0x1c,
	0x85, 0x13, 0x2a, 0x17, 0x3c, 0xb5, 0x5d, 0xdc, 0x04, 0xe0, 0xa0, 0xde, 0xf1, 0xa0, 0x13, 0xb4,
	0x8d, 0x4f, 0xca, 0x08, 0xb4, 0xd2, 0xfb, 0x9f, 0x8f, 0xf0, 0xcd, 0x4e, 0xd3, 0x7d, 0x21, 0x93,
	0xa8, 0x32, 0xbd, 0x51, 0x75, 0x83, 0x40, 0xab, 0x8f, 0x09, 0xe5, 0x82, 0x2d, 0x6b, 0xd7, 0x21,
	0x14, 0xc3, 0x29, 0x54, 0xfb, 0x0e, 0x3e, 0x1d, 0x29, 0xce, 0xb4, 0x54, 0xd7, 0x2e, 0x9d, 0x95,
	0xb9, 0x5b, 0x4b, 0xb4, 0x06, 0xf6, 0x5d, 0xdc, 0xa6, 0xf2, 0x92, 0x2d, 0xc0, 0xce, 0x06, 0xc4,
	0xdb, 0x2d, 0x73, 0x77, 0x2f, 0xd2, 0x3d, 0x34, 0x1b, 0x9f, 0xb1, 0x74, 0x7e, 0xfd, 0x66, 0xd8,
	0x68, 0x38, 0x85, 0x6a, 0x4e, 0x5f, 0xd3, 0x71, 0xea, 0x9c, 0xc0, 0xeb, 0xe0, 0xd4, 0x70, 0x0a,
	0xd5, 0x04, 0xf7, 0x58, 0x6b, 0x25, 0xa6, 0x99, 0xe6, 0xa9, 0x73, 0xba, 0x0f, 0x6e, 0xaf, 0xd2,
	0x03, 0x1c, 0x3c, 0x5a, 0x6f, 0x89, 0xb5, 0xd9, 0x12, 0xeb, 0x6a, 0x4b, 0xd0, 0xa7, 0x82, 0xa0,
	0xaf, 0x05, 0x41, 0xdf, 0x0b, 0x82, 0xd6, 0x05, 0x41, 0x9b, 0x82, 0xa0, 0x5f, 0x05, 0x41, 0xbf,
	0x0b, 0x62, 0x5d, 0x15, 0x04, 0x7d, 0xd9, 0x11, 0x6b, 0xbd, 0x23, 0xd6, 0x66, 0x47, 0xac, 0xb7,
	0x0d, 0xf3, 0x47, 0x4c, 0x4f, 0x20, 0xe0, 0x07, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0x4e,
	0x6e, 0xbe, 0x26, 0x03, 0x00, 0x00,
}

func (this *DCDigitalToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DCDigitalToken)
	if !ok {
		that2, ok := that.(DCDigitalToken)
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
	if this.Type != that1.Type {
		return false
	}
	{
		__caster := &github_com_kalyan3104_k_chain_core_go_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.Properties, that1.Properties) {
		return false
	}
	if !this.TokenMetaData.Equal(that1.TokenMetaData) {
		return false
	}
	if !bytes.Equal(this.Reserved, that1.Reserved) {
		return false
	}
	return true
}
func (this *DCDTRoles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DCDTRoles)
	if !ok {
		that2, ok := that.(DCDTRoles)
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
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if !bytes.Equal(this.Roles[i], that1.Roles[i]) {
			return false
		}
	}
	return true
}
func (this *MetaData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetaData)
	if !ok {
		that2, ok := that.(MetaData)
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
	if this.Nonce != that1.Nonce {
		return false
	}
	if !bytes.Equal(this.Name, that1.Name) {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	if this.Royalties != that1.Royalties {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if len(this.URIs) != len(that1.URIs) {
		return false
	}
	for i := range this.URIs {
		if !bytes.Equal(this.URIs[i], that1.URIs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Attributes, that1.Attributes) {
		return false
	}
	return true
}
func (this *DCDigitalToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&dcdt.DCDigitalToken{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "Properties: "+fmt.Sprintf("%#v", this.Properties)+",\n")
	if this.TokenMetaData != nil {
		s = append(s, "TokenMetaData: "+fmt.Sprintf("%#v", this.TokenMetaData)+",\n")
	}
	s = append(s, "Reserved: "+fmt.Sprintf("%#v", this.Reserved)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DCDTRoles) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dcdt.DCDTRoles{")
	s = append(s, "Roles: "+fmt.Sprintf("%#v", this.Roles)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MetaData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&dcdt.MetaData{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Creator: "+fmt.Sprintf("%#v", this.Creator)+",\n")
	s = append(s, "Royalties: "+fmt.Sprintf("%#v", this.Royalties)+",\n")
	s = append(s, "Hash: "+fmt.Sprintf("%#v", this.Hash)+",\n")
	s = append(s, "URIs: "+fmt.Sprintf("%#v", this.URIs)+",\n")
	s = append(s, "Attributes: "+fmt.Sprintf("%#v", this.Attributes)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDcdt(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DCDigitalToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DCDigitalToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DCDigitalToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintDcdt(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TokenMetaData != nil {
		{
			size, err := m.TokenMetaData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDcdt(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Properties) > 0 {
		i -= len(m.Properties)
		copy(dAtA[i:], m.Properties)
		i = encodeVarintDcdt(dAtA, i, uint64(len(m.Properties)))
		i--
		dAtA[i] = 0x1a
	}
	{
		__caster := &github_com_kalyan3104_k_chain_core_go_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDcdt(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Type != 0 {
		i = encodeVarintDcdt(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DCDTRoles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DCDTRoles) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DCDTRoles) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintDcdt(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MetaData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetaData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetaData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		i -= len(m.Attributes)
		copy(dAtA[i:], m.Attributes)
		i = encodeVarintDcdt(dAtA, i, uint64(len(m.Attributes)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.URIs) > 0 {
		for iNdEx := len(m.URIs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.URIs[iNdEx])
			copy(dAtA[i:], m.URIs[iNdEx])
			i = encodeVarintDcdt(dAtA, i, uint64(len(m.URIs[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintDcdt(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Royalties != 0 {
		i = encodeVarintDcdt(dAtA, i, uint64(m.Royalties))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDcdt(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDcdt(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Nonce != 0 {
		i = encodeVarintDcdt(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDcdt(dAtA []byte, offset int, v uint64) int {
	offset -= sovDcdt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DCDigitalToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovDcdt(uint64(m.Type))
	}
	{
		__caster := &github_com_kalyan3104_k_chain_core_go_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovDcdt(uint64(l))
	}
	l = len(m.Properties)
	if l > 0 {
		n += 1 + l + sovDcdt(uint64(l))
	}
	if m.TokenMetaData != nil {
		l = m.TokenMetaData.Size()
		n += 1 + l + sovDcdt(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovDcdt(uint64(l))
	}
	return n
}

func (m *DCDTRoles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Roles) > 0 {
		for _, b := range m.Roles {
			l = len(b)
			n += 1 + l + sovDcdt(uint64(l))
		}
	}
	return n
}

func (m *MetaData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovDcdt(uint64(m.Nonce))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDcdt(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDcdt(uint64(l))
	}
	if m.Royalties != 0 {
		n += 1 + sovDcdt(uint64(m.Royalties))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovDcdt(uint64(l))
	}
	if len(m.URIs) > 0 {
		for _, b := range m.URIs {
			l = len(b)
			n += 1 + l + sovDcdt(uint64(l))
		}
	}
	l = len(m.Attributes)
	if l > 0 {
		n += 1 + l + sovDcdt(uint64(l))
	}
	return n
}

func sovDcdt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDcdt(x uint64) (n int) {
	return sovDcdt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DCDigitalToken) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DCDigitalToken{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Properties:` + fmt.Sprintf("%v", this.Properties) + `,`,
		`TokenMetaData:` + strings.Replace(this.TokenMetaData.String(), "MetaData", "MetaData", 1) + `,`,
		`Reserved:` + fmt.Sprintf("%v", this.Reserved) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DCDTRoles) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DCDTRoles{`,
		`Roles:` + fmt.Sprintf("%v", this.Roles) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MetaData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MetaData{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Creator:` + fmt.Sprintf("%v", this.Creator) + `,`,
		`Royalties:` + fmt.Sprintf("%v", this.Royalties) + `,`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`URIs:` + fmt.Sprintf("%v", this.URIs) + `,`,
		`Attributes:` + fmt.Sprintf("%v", this.Attributes) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDcdt(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DCDigitalToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDcdt
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
			return fmt.Errorf("proto: DCDigitalToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DCDigitalToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_kalyan3104_k_chain_core_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = append(m.Properties[:0], dAtA[iNdEx:postIndex]...)
			if m.Properties == nil {
				m.Properties = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMetaData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
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
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TokenMetaData == nil {
				m.TokenMetaData = &MetaData{}
			}
			if err := m.TokenMetaData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = append(m.Reserved[:0], dAtA[iNdEx:postIndex]...)
			if m.Reserved == nil {
				m.Reserved = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDcdt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDcdt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDcdt
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
func (m *DCDTRoles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDcdt
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
			return fmt.Errorf("proto: DCDTRoles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DCDTRoles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, make([]byte, postIndex-iNdEx))
			copy(m.Roles[len(m.Roles)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDcdt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDcdt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDcdt
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
func (m *MetaData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDcdt
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
			return fmt.Errorf("proto: MetaData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetaData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = append(m.Name[:0], dAtA[iNdEx:postIndex]...)
			if m.Name == nil {
				m.Name = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Royalties", wireType)
			}
			m.Royalties = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Royalties |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URIs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URIs = append(m.URIs, make([]byte, postIndex-iNdEx))
			copy(m.URIs[len(m.URIs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDcdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDcdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDcdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes[:0], dAtA[iNdEx:postIndex]...)
			if m.Attributes == nil {
				m.Attributes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDcdt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDcdt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDcdt
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
func skipDcdt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDcdt
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
					return 0, ErrIntOverflowDcdt
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
					return 0, ErrIntOverflowDcdt
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
				return 0, ErrInvalidLengthDcdt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDcdt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDcdt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDcdt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDcdt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDcdt = fmt.Errorf("proto: unexpected end of group")
)
