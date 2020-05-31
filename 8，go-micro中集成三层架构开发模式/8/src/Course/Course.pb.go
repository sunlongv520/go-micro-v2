// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Course.proto

package Course

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

type CourseModel struct {
	CourseId             int32    `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CourseName           string   `protobuf:"bytes,2,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CourseModel) Reset()         { *m = CourseModel{} }
func (m *CourseModel) String() string { return proto.CompactTextString(m) }
func (*CourseModel) ProtoMessage()    {}
func (*CourseModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_06dcbc50c02cef66, []int{0}
}

func (m *CourseModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CourseModel.Unmarshal(m, b)
}
func (m *CourseModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CourseModel.Marshal(b, m, deterministic)
}
func (m *CourseModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CourseModel.Merge(m, src)
}
func (m *CourseModel) XXX_Size() int {
	return xxx_messageInfo_CourseModel.Size(m)
}
func (m *CourseModel) XXX_DiscardUnknown() {
	xxx_messageInfo_CourseModel.DiscardUnknown(m)
}

var xxx_messageInfo_CourseModel proto.InternalMessageInfo

func (m *CourseModel) GetCourseId() int32 {
	if m != nil {
		return m.CourseId
	}
	return 0
}

func (m *CourseModel) GetCourseName() string {
	if m != nil {
		return m.CourseName
	}
	return ""
}

type ListRequest struct {
	Size                 int32    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06dcbc50c02cef66, []int{1}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ListResponse struct {
	Result               []*CourseModel `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06dcbc50c02cef66, []int{2}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetResult() []*CourseModel {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CourseModel)(nil), "Course.CourseModel")
	proto.RegisterType((*ListRequest)(nil), "Course.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "Course.ListResponse")
}

func init() {
	proto.RegisterFile("Course.proto", fileDescriptor_06dcbc50c02cef66)
}

var fileDescriptor_06dcbc50c02cef66 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x71, 0xce, 0x2f, 0x2d,
	0x2a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xbc, 0xb9, 0xb8,
	0x21, 0x2c, 0xdf, 0xfc, 0x94, 0xd4, 0x1c, 0x21, 0x69, 0x2e, 0xce, 0x64, 0x30, 0x37, 0x3e, 0x33,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0x03, 0x22, 0xe0, 0x99, 0x22, 0x24, 0xcf, 0xc5,
	0x0d, 0x95, 0xcc, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x82, 0x08,
	0xf9, 0x25, 0xe6, 0xa6, 0x2a, 0x29, 0x72, 0x71, 0xfb, 0x64, 0x16, 0x97, 0x04, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0x14, 0x67, 0x56, 0xa5, 0x42, 0xcd, 0x01, 0xb3, 0x95,
	0xac, 0xb9, 0x78, 0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xb4, 0xb9, 0xd8, 0x8a,
	0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x24, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0xa0, 0xce,
	0x44, 0x72, 0x55, 0x10, 0x54, 0x89, 0x91, 0x17, 0x17, 0x2f, 0x44, 0x38, 0x38, 0xb5, 0xa8, 0x2c,
	0x33, 0x39, 0x55, 0xc8, 0x92, 0x8b, 0x0b, 0x64, 0x9a, 0x5b, 0x7e, 0x51, 0x48, 0x7e, 0x81, 0x10,
	0x5c, 0x2f, 0x92, 0x23, 0xa4, 0x44, 0x50, 0x05, 0x21, 0xd6, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0xc3,
	0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x62, 0x25, 0x46, 0x88, 0x17, 0x01, 0x00, 0x00,
}
