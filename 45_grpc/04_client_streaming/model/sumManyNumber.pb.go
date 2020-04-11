// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/sumManyNumber.proto

package model

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SumReq struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumReq) Reset()         { *m = SumReq{} }
func (m *SumReq) String() string { return proto.CompactTextString(m) }
func (*SumReq) ProtoMessage()    {}
func (*SumReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a031677d4850f4, []int{0}
}

func (m *SumReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumReq.Unmarshal(m, b)
}
func (m *SumReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumReq.Marshal(b, m, deterministic)
}
func (m *SumReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumReq.Merge(m, src)
}
func (m *SumReq) XXX_Size() int {
	return xxx_messageInfo_SumReq.Size(m)
}
func (m *SumReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SumReq.DiscardUnknown(m)
}

var xxx_messageInfo_SumReq proto.InternalMessageInfo

func (m *SumReq) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SumResp struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResp) Reset()         { *m = SumResp{} }
func (m *SumResp) String() string { return proto.CompactTextString(m) }
func (*SumResp) ProtoMessage()    {}
func (*SumResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a031677d4850f4, []int{1}
}

func (m *SumResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResp.Unmarshal(m, b)
}
func (m *SumResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResp.Marshal(b, m, deterministic)
}
func (m *SumResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResp.Merge(m, src)
}
func (m *SumResp) XXX_Size() int {
	return xxx_messageInfo_SumResp.Size(m)
}
func (m *SumResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResp.DiscardUnknown(m)
}

var xxx_messageInfo_SumResp proto.InternalMessageInfo

func (m *SumResp) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SumReq)(nil), "model.sumReq")
	proto.RegisterType((*SumResp)(nil), "model.sumResp")
}

func init() { proto.RegisterFile("model/sumManyNumber.proto", fileDescriptor_30a031677d4850f4) }

var fileDescriptor_30a031677d4850f4 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2f, 0x2e, 0xcd, 0xf5, 0x4d, 0xcc, 0xab, 0xf4, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x4b, 0x29, 0x29, 0x70, 0xb1, 0x15, 0x97, 0xe6,
	0x06, 0xa5, 0x16, 0x0a, 0x89, 0x71, 0xb1, 0xe5, 0x81, 0x15, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30,
	0x07, 0x41, 0x79, 0x4a, 0x8a, 0x5c, 0xec, 0x60, 0x15, 0xc5, 0x05, 0x20, 0x25, 0x45, 0xa9, 0xc5,
	0xa5, 0x39, 0x25, 0x30, 0x25, 0x10, 0x9e, 0x91, 0x05, 0x17, 0x67, 0x71, 0x69, 0x2e, 0xc4, 0x78,
	0x21, 0x6d, 0xb0, 0x89, 0x8e, 0x39, 0x39, 0x42, 0xbc, 0x7a, 0x60, 0x3b, 0xf4, 0x20, 0x16, 0x48,
	0xf1, 0x21, 0x73, 0x8b, 0x0b, 0x94, 0x18, 0x34, 0x18, 0x9d, 0xd8, 0xa3, 0x20, 0xee, 0x48, 0x62,
	0x03, 0xbb, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xef, 0x44, 0xa6, 0xac, 0xb2, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumNumberClient is the client API for SumNumber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumNumberClient interface {
	SumAll(ctx context.Context, opts ...grpc.CallOption) (SumNumber_SumAllClient, error)
}

type sumNumberClient struct {
	cc *grpc.ClientConn
}

func NewSumNumberClient(cc *grpc.ClientConn) SumNumberClient {
	return &sumNumberClient{cc}
}

func (c *sumNumberClient) SumAll(ctx context.Context, opts ...grpc.CallOption) (SumNumber_SumAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumNumber_serviceDesc.Streams[0], "/model.sumNumber/sumAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumNumberSumAllClient{stream}
	return x, nil
}

type SumNumber_SumAllClient interface {
	Send(*SumReq) error
	CloseAndRecv() (*SumResp, error)
	grpc.ClientStream
}

type sumNumberSumAllClient struct {
	grpc.ClientStream
}

func (x *sumNumberSumAllClient) Send(m *SumReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumNumberSumAllClient) CloseAndRecv() (*SumResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SumResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumNumberServer is the server API for SumNumber service.
type SumNumberServer interface {
	SumAll(SumNumber_SumAllServer) error
}

func RegisterSumNumberServer(s *grpc.Server, srv SumNumberServer) {
	s.RegisterService(&_SumNumber_serviceDesc, srv)
}

func _SumNumber_SumAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumNumberServer).SumAll(&sumNumberSumAllServer{stream})
}

type SumNumber_SumAllServer interface {
	SendAndClose(*SumResp) error
	Recv() (*SumReq, error)
	grpc.ServerStream
}

type sumNumberSumAllServer struct {
	grpc.ServerStream
}

func (x *sumNumberSumAllServer) SendAndClose(m *SumResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumNumberSumAllServer) Recv() (*SumReq, error) {
	m := new(SumReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SumNumber_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.sumNumber",
	HandlerType: (*SumNumberServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sumAll",
			Handler:       _SumNumber_SumAll_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "model/sumManyNumber.proto",
}