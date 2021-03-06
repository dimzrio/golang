// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type RequestSum struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestSum) Reset()         { *m = RequestSum{} }
func (m *RequestSum) String() string { return proto.CompactTextString(m) }
func (*RequestSum) ProtoMessage()    {}
func (*RequestSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{0}
}

func (m *RequestSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestSum.Unmarshal(m, b)
}
func (m *RequestSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestSum.Marshal(b, m, deterministic)
}
func (m *RequestSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSum.Merge(m, src)
}
func (m *RequestSum) XXX_Size() int {
	return xxx_messageInfo_RequestSum.Size(m)
}
func (m *RequestSum) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSum.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSum proto.InternalMessageInfo

func (m *RequestSum) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *RequestSum) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type ResponseSum struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseSum) Reset()         { *m = ResponseSum{} }
func (m *ResponseSum) String() string { return proto.CompactTextString(m) }
func (*ResponseSum) ProtoMessage()    {}
func (*ResponseSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{1}
}

func (m *ResponseSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseSum.Unmarshal(m, b)
}
func (m *ResponseSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseSum.Marshal(b, m, deterministic)
}
func (m *ResponseSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseSum.Merge(m, src)
}
func (m *ResponseSum) XXX_Size() int {
	return xxx_messageInfo_ResponseSum.Size(m)
}
func (m *ResponseSum) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseSum.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseSum proto.InternalMessageInfo

func (m *ResponseSum) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestSum)(nil), "model.RequestSum")
	proto.RegisterType((*ResponseSum)(nil), "model.ResponseSum")
}

func init() { proto.RegisterFile("calc.proto", fileDescriptor_a2b9900dc883ea68) }

var fileDescriptor_a2b9900dc883ea68 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0xcc, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x51, 0xd2, 0xe0,
	0xe2, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x09, 0x2e, 0xcd, 0x15, 0xe2, 0xe1, 0x62, 0xac,
	0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xac, 0x00, 0xf1, 0x2a, 0x25, 0x98, 0x20, 0xbc,
	0x4a, 0x25, 0x55, 0x2e, 0xee, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0x52, 0x31,
	0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x7a, 0x28, 0xcf, 0xc8, 0x9c, 0x8b, 0xdd,
	0x39, 0x31, 0x27, 0x19, 0xa4, 0x44, 0x87, 0x8b, 0x19, 0x44, 0x09, 0xea, 0x81, 0xad, 0xd2, 0x43,
	0xd8, 0x23, 0x25, 0x04, 0x17, 0x82, 0x1b, 0xe8, 0xc4, 0x1e, 0x05, 0x71, 0x52, 0x12, 0x1b, 0xd8,
	0x81, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x2b, 0xc2, 0x1d, 0xae, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalcSumClient is the client API for CalcSum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcSumClient interface {
	Sum(ctx context.Context, in *RequestSum, opts ...grpc.CallOption) (*ResponseSum, error)
}

type calcSumClient struct {
	cc *grpc.ClientConn
}

func NewCalcSumClient(cc *grpc.ClientConn) CalcSumClient {
	return &calcSumClient{cc}
}

func (c *calcSumClient) Sum(ctx context.Context, in *RequestSum, opts ...grpc.CallOption) (*ResponseSum, error) {
	out := new(ResponseSum)
	err := c.cc.Invoke(ctx, "/model.CalcSum/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcSumServer is the server API for CalcSum service.
type CalcSumServer interface {
	Sum(context.Context, *RequestSum) (*ResponseSum, error)
}

func RegisterCalcSumServer(s *grpc.Server, srv CalcSumServer) {
	s.RegisterService(&_CalcSum_serviceDesc, srv)
}

func _CalcSum_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcSumServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.CalcSum/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcSumServer).Sum(ctx, req.(*RequestSum))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalcSum_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.CalcSum",
	HandlerType: (*CalcSumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalcSum_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}
