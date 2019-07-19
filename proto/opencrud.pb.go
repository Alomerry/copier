// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencrud.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	// id: ID!
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name: String!
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// age: Int!
	Age int64 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	// sex: Sex!
	Sex string `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	// weight: Float!
	Weight float64 `protobuf:"fixed64,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// alive: Boolean!
	Alive bool `protobuf:"varint,6,opt,name=alive,proto3" json:"alive,omitempty"`
	// num64: Int64!
	Num64 int64 `protobuf:"varint,7,opt,name=num64,proto3" json:"num64,omitempty"`
	// optional_num: Int
	OptionalNum *wrappers.Int64Value `protobuf:"bytes,8,opt,name=optional_num,json=optionalNum,proto3" json:"optional_num,omitempty"`
	// optional_num64: Int64
	OptionalNum64 *wrappers.Int64Value `protobuf:"bytes,9,opt,name=optional_num64,json=optionalNum64,proto3" json:"optional_num64,omitempty"`
	// birth_date: Date
	BirthDate string `protobuf:"bytes,10,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	// created_at: Time!
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// modified_at: Time!
	ModifiedAt           *timestamp.Timestamp `protobuf:"bytes,12,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0434410d2705632f, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *User) GetWeight() float64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *User) GetAlive() bool {
	if m != nil {
		return m.Alive
	}
	return false
}

func (m *User) GetNum64() int64 {
	if m != nil {
		return m.Num64
	}
	return 0
}

func (m *User) GetOptionalNum() *wrappers.Int64Value {
	if m != nil {
		return m.OptionalNum
	}
	return nil
}

func (m *User) GetOptionalNum64() *wrappers.Int64Value {
	if m != nil {
		return m.OptionalNum64
	}
	return nil
}

func (m *User) GetBirthDate() string {
	if m != nil {
		return m.BirthDate
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetModifiedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ModifiedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
}

func init() { proto.RegisterFile("opencrud.proto", fileDescriptor_0434410d2705632f) }

var fileDescriptor_0434410d2705632f = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x64, 0x9b, 0xb6, 0x36, 0x2f, 0xb5, 0xc8, 0x22, 0xb2, 0x54, 0xd4, 0xe0, 0x29, 0xa7, 0x14,
	0x34, 0x04, 0x44, 0x10, 0x2a, 0x5e, 0xbc, 0x78, 0x08, 0xea, 0xb5, 0x6c, 0x9a, 0xd7, 0x74, 0x21,
	0x9b, 0x0d, 0x9b, 0x8d, 0xf5, 0xa7, 0xfc, 0x47, 0xc9, 0xa6, 0x41, 0xd1, 0x43, 0x4f, 0x99, 0x99,
	0x37, 0x33, 0x30, 0x59, 0x98, 0xa9, 0x0a, 0xcb, 0xb5, 0x6e, 0xb2, 0xb0, 0xd2, 0xca, 0x28, 0x3a,
	0xb2, 0x9f, 0xf9, 0x55, 0xae, 0x54, 0x5e, 0xe0, 0xc2, 0xb2, 0xb4, 0xd9, 0x2c, 0x8c, 0x90, 0x58,
	0x1b, 0x2e, 0xab, 0xce, 0x37, 0xbf, 0xfc, 0x6b, 0xd8, 0x69, 0x5e, 0x55, 0xa8, 0xeb, 0xee, 0x7e,
	0xfd, 0xe5, 0xc0, 0xf0, 0xad, 0x46, 0x4d, 0x67, 0x30, 0x10, 0x19, 0x23, 0x3e, 0x09, 0xdc, 0x64,
	0x20, 0x32, 0x4a, 0x61, 0x58, 0x72, 0x89, 0x6c, 0x60, 0x15, 0x8b, 0xe9, 0x09, 0x38, 0x3c, 0x47,
	0xe6, 0xf8, 0x24, 0x70, 0x92, 0x16, 0xb6, 0x4a, 0x8d, 0x9f, 0x6c, 0x68, 0x4d, 0x2d, 0xa4, 0x67,
	0x30, 0xde, 0xa1, 0xc8, 0xb7, 0x86, 0x8d, 0x7c, 0x12, 0x90, 0x64, 0xcf, 0xe8, 0x29, 0x8c, 0x78,
	0x21, 0x3e, 0x90, 0x8d, 0x7d, 0x12, 0x4c, 0x92, 0x8e, 0xb4, 0x6a, 0xd9, 0xc8, 0x38, 0x62, 0x47,
	0xb6, 0xb3, 0x23, 0xf4, 0x01, 0xa6, 0xaa, 0x32, 0x42, 0x95, 0xbc, 0x58, 0x95, 0x8d, 0x64, 0x13,
	0x9f, 0x04, 0xde, 0xcd, 0x79, 0xd8, 0x6d, 0x09, 0xfb, 0x2d, 0xe1, 0x73, 0x69, 0xe2, 0xe8, 0x9d,
	0x17, 0x0d, 0x26, 0x5e, 0x1f, 0x78, 0x69, 0x24, 0x7d, 0x6c, 0x7f, 0xd7, 0x4f, 0x3e, 0x8e, 0x98,
	0x7b, 0xb8, 0xe1, 0xf8, 0x57, 0x43, 0x1c, 0xd1, 0x0b, 0x80, 0x54, 0x68, 0xb3, 0x5d, 0x65, 0xdc,
	0x20, 0x03, 0x3b, 0xd0, 0xb5, 0xca, 0x13, 0x37, 0x48, 0xef, 0x00, 0xd6, 0x1a, 0xb9, 0xc1, 0x6c,
	0xc5, 0x0d, 0xf3, 0x6c, 0xfd, 0xfc, 0x5f, 0xfd, 0x6b, 0xff, 0x1a, 0x89, 0xbb, 0x77, 0x2f, 0x0d,
	0xbd, 0x07, 0x4f, 0xaa, 0x4c, 0x6c, 0x44, 0x97, 0x9d, 0x1e, 0xcc, 0x42, 0x6f, 0x5f, 0x9a, 0x74,
	0x6c, 0xef, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0xe7, 0x34, 0x1c, 0x10, 0x02, 0x00,
	0x00,
}