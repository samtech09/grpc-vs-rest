// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/output.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Filter struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7e928544d0b2f74, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Filter) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type StudentDetails struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RollNo               string   `protobuf:"bytes,3,opt,name=rollNo,proto3" json:"rollNo,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	ExamCleared          bool     `protobuf:"varint,5,opt,name=examCleared,proto3" json:"examCleared,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StudentDetails) Reset()         { *m = StudentDetails{} }
func (m *StudentDetails) String() string { return proto.CompactTextString(m) }
func (*StudentDetails) ProtoMessage()    {}
func (*StudentDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7e928544d0b2f74, []int{1}
}

func (m *StudentDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StudentDetails.Unmarshal(m, b)
}
func (m *StudentDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StudentDetails.Marshal(b, m, deterministic)
}
func (m *StudentDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentDetails.Merge(m, src)
}
func (m *StudentDetails) XXX_Size() int {
	return xxx_messageInfo_StudentDetails.Size(m)
}
func (m *StudentDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentDetails.DiscardUnknown(m)
}

var xxx_messageInfo_StudentDetails proto.InternalMessageInfo

func (m *StudentDetails) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StudentDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StudentDetails) GetRollNo() string {
	if m != nil {
		return m.RollNo
	}
	return ""
}

func (m *StudentDetails) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *StudentDetails) GetExamCleared() bool {
	if m != nil {
		return m.ExamCleared
	}
	return false
}

func init() {
	proto.RegisterType((*Filter)(nil), "model.Filter")
	proto.RegisterType((*StudentDetails)(nil), "model.studentDetails")
}

func init() { proto.RegisterFile("model/output.proto", fileDescriptor_e7e928544d0b2f74) }

var fileDescriptor_e7e928544d0b2f74 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xbd, 0x6a, 0xc4, 0x30,
	0x0c, 0x80, 0xb1, 0xf3, 0x43, 0xab, 0x42, 0x28, 0x1a, 0x8a, 0xc7, 0x90, 0x29, 0x43, 0x69, 0x87,
	0x3e, 0x42, 0x4b, 0xc7, 0x1b, 0xf2, 0x06, 0x3e, 0xac, 0x3b, 0x0c, 0x76, 0x14, 0x1c, 0x05, 0x6e,
	0xbc, 0x47, 0x3f, 0x22, 0x32, 0xdc, 0xf6, 0xe9, 0x43, 0x9f, 0x10, 0x60, 0xe6, 0x40, 0xe9, 0x9b,
	0x37, 0x59, 0x36, 0xf9, 0x5a, 0x0a, 0x0b, 0x63, 0xa3, 0x6e, 0xf8, 0x84, 0xf6, 0x3f, 0x26, 0xa1,
	0x82, 0x08, 0xf5, 0xa5, 0x70, 0x76, 0xa6, 0x37, 0xe3, 0xeb, 0xa4, 0x8c, 0x1d, 0x58, 0x61, 0x67,
	0xd5, 0x58, 0xe1, 0xe1, 0x6e, 0xa0, 0x5b, 0x65, 0x0b, 0x34, 0xcb, 0x1f, 0x89, 0x8f, 0x69, 0xdd,
	0x57, 0x62, 0xd0, 0xa8, 0x99, 0x6c, 0x0c, 0xfb, 0x99, 0xd9, 0x67, 0x3a, 0x22, 0x65, 0xfc, 0x80,
	0xb6, 0x70, 0x4a, 0x27, 0x76, 0x95, 0xda, 0x63, 0xc2, 0x77, 0xa8, 0xfc, 0x95, 0x5c, 0xad, 0xf1,
	0x8e, 0xd8, 0xc3, 0x1b, 0xdd, 0x7c, 0xfe, 0x4d, 0xe4, 0x0b, 0x05, 0xd7, 0xf4, 0x66, 0x7c, 0x99,
	0x9e, 0xd5, 0xb9, 0xd5, 0xf7, 0x7f, 0x1e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x51, 0x90, 0x2b, 0xb9,
	0xd4, 0x00, 0x00, 0x00,
}
