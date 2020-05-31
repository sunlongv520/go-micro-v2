// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Course.proto

package Course

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Timestamp struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Timestr              string               `protobuf:"bytes,2,opt,name=timestr,proto3" json:"timestr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_06dcbc50c02cef66, []int{0}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Timestamp) GetTimestr() string {
	if m != nil {
		return m.Timestr
	}
	return ""
}

type CourseModel struct {
	CourseId       int32   `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CourseName     string  `protobuf:"bytes,2,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	CourseDispName string  `protobuf:"bytes,3,opt,name=course_disp_name,json=courseDispName,proto3" json:"course_disp_name,omitempty"`
	CoursePrice    float32 `protobuf:"fixed32,4,opt,name=course_price,json=coursePrice,proto3" json:"course_price,omitempty"`
	CoursePrice2   float32 `protobuf:"fixed32,5,opt,name=course_price2,json=coursePrice2,proto3" json:"course_price2,omitempty"`
	// @inject_tag: gorm:"type:timestamp"
	Addtime              *Timestamp `protobuf:"bytes,6,opt,name=addtime,proto3" json:"addtime,omitempty" gorm:"type:timestamp"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CourseModel) Reset()         { *m = CourseModel{} }
func (m *CourseModel) String() string { return proto.CompactTextString(m) }
func (*CourseModel) ProtoMessage()    {}
func (*CourseModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_06dcbc50c02cef66, []int{1}
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

func (m *CourseModel) GetCourseDispName() string {
	if m != nil {
		return m.CourseDispName
	}
	return ""
}

func (m *CourseModel) GetCoursePrice() float32 {
	if m != nil {
		return m.CoursePrice
	}
	return 0
}

func (m *CourseModel) GetCoursePrice2() float32 {
	if m != nil {
		return m.CoursePrice2
	}
	return 0
}

func (m *CourseModel) GetAddtime() *Timestamp {
	if m != nil {
		return m.Addtime
	}
	return nil
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
	return fileDescriptor_06dcbc50c02cef66, []int{2}
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

type DetailRequest struct {
	// @inject_tag: uri:"course_id"
	CourseId             int32    `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty" uri:"course_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06dcbc50c02cef66, []int{3}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

func (m *DetailRequest) GetCourseId() int32 {
	if m != nil {
		return m.CourseId
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
	return fileDescriptor_06dcbc50c02cef66, []int{4}
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

type DetailResponse struct {
	Result               *CourseModel `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DetailResponse) Reset()         { *m = DetailResponse{} }
func (m *DetailResponse) String() string { return proto.CompactTextString(m) }
func (*DetailResponse) ProtoMessage()    {}
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06dcbc50c02cef66, []int{5}
}

func (m *DetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResponse.Unmarshal(m, b)
}
func (m *DetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResponse.Marshal(b, m, deterministic)
}
func (m *DetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResponse.Merge(m, src)
}
func (m *DetailResponse) XXX_Size() int {
	return xxx_messageInfo_DetailResponse.Size(m)
}
func (m *DetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResponse proto.InternalMessageInfo

func (m *DetailResponse) GetResult() *CourseModel {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "Course.Timestamp")
	proto.RegisterType((*CourseModel)(nil), "Course.CourseModel")
	proto.RegisterType((*ListRequest)(nil), "Course.ListRequest")
	proto.RegisterType((*DetailRequest)(nil), "Course.DetailRequest")
	proto.RegisterType((*ListResponse)(nil), "Course.ListResponse")
	proto.RegisterType((*DetailResponse)(nil), "Course.DetailResponse")
}

func init() {
	proto.RegisterFile("Course.proto", fileDescriptor_06dcbc50c02cef66)
}

var fileDescriptor_06dcbc50c02cef66 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcb, 0x6b, 0xdb, 0x40,
	0x10, 0xc6, 0x59, 0x3f, 0xe4, 0x6a, 0x24, 0x9b, 0x76, 0xfb, 0x40, 0xa8, 0x07, 0xcb, 0xea, 0x45,
	0xe0, 0x22, 0x83, 0x7a, 0x68, 0x69, 0xe9, 0xa9, 0xa6, 0xa5, 0x90, 0x84, 0xa0, 0xf8, 0x6e, 0x64,
	0x6b, 0x63, 0x16, 0x24, 0xaf, 0xa2, 0x5d, 0xe5, 0x90, 0x73, 0xfe, 0xda, 0xfc, 0x15, 0x41, 0xfb,
	0xb0, 0xe5, 0x10, 0x92, 0x9b, 0xe6, 0xdb, 0xdf, 0x7c, 0x33, 0xf3, 0x21, 0x70, 0xff, 0xb0, 0xa6,
	0xe6, 0x24, 0xae, 0x6a, 0x26, 0x18, 0xb6, 0x54, 0xe5, 0x4f, 0x77, 0x8c, 0xed, 0x0a, 0xb2, 0x90,
	0xea, 0xa6, 0xb9, 0x5e, 0x08, 0x5a, 0x12, 0x2e, 0xb2, 0xb2, 0x52, 0x60, 0xb8, 0x06, 0x7b, 0x65,
	0x24, 0xfc, 0x03, 0xec, 0xc3, 0xbb, 0x87, 0x02, 0x14, 0x39, 0x89, 0x1f, 0x2b, 0x87, 0xd8, 0x38,
	0xc4, 0x07, 0x3c, 0x3d, 0xc2, 0xd8, 0x83, 0x91, 0x2a, 0x6a, 0xaf, 0x17, 0xa0, 0xc8, 0x4e, 0x4d,
	0x19, 0x3e, 0x20, 0x70, 0xd4, 0x32, 0xe7, 0x2c, 0x27, 0x05, 0xfe, 0x0c, 0xf6, 0x56, 0x96, 0x6b,
	0x9a, 0xcb, 0x19, 0xc3, 0xf4, 0x8d, 0x12, 0xfe, 0xe7, 0x78, 0x0a, 0x8e, 0x7e, 0xdc, 0x67, 0x25,
	0xd1, 0x56, 0xa0, 0xa4, 0x8b, 0xac, 0x24, 0x38, 0x82, 0xb7, 0x1a, 0xc8, 0x29, 0xaf, 0x14, 0xd5,
	0x97, 0xd4, 0x44, 0xe9, 0x4b, 0xca, 0x2b, 0x49, 0xce, 0xc0, 0xd5, 0x64, 0x55, 0xd3, 0x2d, 0xf1,
	0x06, 0x01, 0x8a, 0x7a, 0xa9, 0xb6, 0xbf, 0x6c, 0x25, 0xfc, 0x05, 0xc6, 0x5d, 0x24, 0xf1, 0x86,
	0x92, 0x71, 0x3b, 0x4c, 0x82, 0xe7, 0x30, 0xca, 0xf2, 0xbc, 0xbd, 0xc6, 0xb3, 0x64, 0x22, 0xef,
	0x62, 0x9d, 0xf4, 0x31, 0x08, 0x43, 0x84, 0x33, 0x70, 0xce, 0x28, 0x17, 0x29, 0xb9, 0x69, 0x08,
	0x17, 0x18, 0xc3, 0x80, 0xd3, 0x3b, 0xa2, 0xcf, 0x94, 0xdf, 0xe1, 0x57, 0x18, 0x2f, 0x89, 0xc8,
	0x68, 0x61, 0xa0, 0x97, 0x02, 0x09, 0x7f, 0x81, 0xab, 0x0c, 0x79, 0xc5, 0xf6, 0x9c, 0xe0, 0x39,
	0x58, 0x35, 0xe1, 0x4d, 0x21, 0x3c, 0x14, 0xf4, 0x23, 0x27, 0x79, 0x6f, 0x96, 0xe9, 0x44, 0x9c,
	0x6a, 0x24, 0xfc, 0x0d, 0x13, 0x33, 0xea, 0x99, 0x76, 0xf4, 0x4a, 0x7b, 0x72, 0x8f, 0x60, 0xac,
	0xf4, 0x2b, 0x52, 0xdf, 0xb6, 0x81, 0x7d, 0x07, 0x68, 0xb7, 0xf9, 0xcb, 0xea, 0x15, 0xab, 0xf0,
	0xa1, 0xb9, 0x73, 0xb2, 0xff, 0xe1, 0x54, 0xd4, 0x73, 0x7f, 0x82, 0xfd, 0x8f, 0x08, 0xb5, 0x0c,
	0xfe, 0x68, 0x90, 0x93, 0x1c, 0xfc, 0x4f, 0x4f, 0x65, 0xd5, 0xbb, 0xb1, 0xe4, 0x9f, 0xf7, 0xed,
	0x31, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xcf, 0xc7, 0x39, 0xe1, 0x02, 0x00, 0x00,
}
