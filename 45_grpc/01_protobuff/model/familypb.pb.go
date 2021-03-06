// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/familypb.proto

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

type JobDesc struct {
	Company              string   `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Possition            string   `protobuf:"bytes,2,opt,name=possition,proto3" json:"possition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobDesc) Reset()         { *m = JobDesc{} }
func (m *JobDesc) String() string { return proto.CompactTextString(m) }
func (*JobDesc) ProtoMessage()    {}
func (*JobDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ebeb307ae0b942a, []int{0}
}

func (m *JobDesc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobDesc.Unmarshal(m, b)
}
func (m *JobDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobDesc.Marshal(b, m, deterministic)
}
func (m *JobDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobDesc.Merge(m, src)
}
func (m *JobDesc) XXX_Size() int {
	return xxx_messageInfo_JobDesc.Size(m)
}
func (m *JobDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_JobDesc.DiscardUnknown(m)
}

var xxx_messageInfo_JobDesc proto.InternalMessageInfo

func (m *JobDesc) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *JobDesc) GetPossition() string {
	if m != nil {
		return m.Possition
	}
	return ""
}

type Family struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Job                  *JobDesc `protobuf:"bytes,6,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Family) Reset()         { *m = Family{} }
func (m *Family) String() string { return proto.CompactTextString(m) }
func (*Family) ProtoMessage()    {}
func (*Family) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ebeb307ae0b942a, []int{1}
}

func (m *Family) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Family.Unmarshal(m, b)
}
func (m *Family) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Family.Marshal(b, m, deterministic)
}
func (m *Family) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Family.Merge(m, src)
}
func (m *Family) XXX_Size() int {
	return xxx_messageInfo_Family.Size(m)
}
func (m *Family) XXX_DiscardUnknown() {
	xxx_messageInfo_Family.DiscardUnknown(m)
}

var xxx_messageInfo_Family proto.InternalMessageInfo

func (m *Family) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Family) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Family) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Family) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Family) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Family) GetJob() *JobDesc {
	if m != nil {
		return m.Job
	}
	return nil
}

func init() {
	proto.RegisterType((*JobDesc)(nil), "model.JobDesc")
	proto.RegisterType((*Family)(nil), "model.Family")
}

func init() { proto.RegisterFile("model/familypb.proto", fileDescriptor_8ebeb307ae0b942a) }

var fileDescriptor_8ebeb307ae0b942a = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x49, 0xbb, 0xcd, 0xb2, 0x23, 0x2c, 0x32, 0x88, 0xe4, 0xe0, 0xa1, 0xec, 0x69, 0x4f,
	0x15, 0xf4, 0x09, 0x04, 0xf1, 0xe0, 0xb1, 0x6f, 0x90, 0x34, 0xb1, 0x44, 0x9a, 0x4c, 0x68, 0xe2,
	0xa1, 0x6f, 0xe1, 0x23, 0x4b, 0xa7, 0x95, 0xbd, 0xcd, 0xff, 0x6b, 0xe0, 0x07, 0x0f, 0x81, 0xac,
	0x9b, 0x9e, 0xbf, 0x74, 0xf0, 0xd3, 0x92, 0x4c, 0x97, 0x66, 0x2a, 0x84, 0x0d, 0xbb, 0x97, 0x37,
	0x38, 0x7e, 0x92, 0x79, 0x77, 0x79, 0x40, 0x05, 0xc7, 0x81, 0x42, 0xd2, 0x71, 0x51, 0xa2, 0x15,
	0xd7, 0x53, 0xff, 0x2f, 0xf1, 0x09, 0x4e, 0x89, 0x72, 0xf6, 0xc5, 0x53, 0x54, 0x15, 0x67, 0x37,
	0xe3, 0xf2, 0x2b, 0x40, 0x7e, 0xf0, 0x73, 0x3c, 0x43, 0xe5, 0x2d, 0xaf, 0x9b, 0xbe, 0xf2, 0x16,
	0x11, 0x0e, 0x51, 0x07, 0xb7, 0x6f, 0xf8, 0xc6, 0x47, 0x90, 0xa3, 0x8b, 0xd6, 0xcd, 0xaa, 0x66,
	0x77, 0x57, 0x78, 0x0f, 0xb5, 0x1e, 0x9d, 0x3a, 0xf0, 0x78, 0x3d, 0xd7, 0x66, 0x2e, 0xba, 0xfc,
	0x64, 0xd5, 0x6c, 0xcd, 0x4d, 0x61, 0x0b, 0xf5, 0x37, 0x19, 0x25, 0x5b, 0x71, 0xbd, 0x7b, 0x39,
	0x77, 0x0c, 0xd2, 0xed, 0x14, 0xfd, 0x1a, 0x19, 0xc9, 0x8c, 0xaf, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x76, 0x89, 0x29, 0x9f, 0xfb, 0x00, 0x00, 0x00,
}
