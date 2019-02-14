// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package notification

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_notification_95bceaf56281d715, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_notification_95bceaf56281d715, []int{1}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (dst *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(dst, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "notification.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "notification.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	ManyHellos(ctx context.Context, opts ...grpc.CallOption) (Greeter_ManyHellosClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/notification.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ManyHellos(ctx context.Context, opts ...grpc.CallOption) (Greeter_ManyHellosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/notification.Greeter/ManyHellos", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterManyHellosClient{stream}
	return x, nil
}

type Greeter_ManyHellosClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterManyHellosClient struct {
	grpc.ClientStream
}

func (x *greeterManyHellosClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterManyHellosClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	ManyHellos(Greeter_ManyHellosServer) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ManyHellos_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).ManyHellos(&greeterManyHellosServer{stream})
}

type Greeter_ManyHellosServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterManyHellosServer struct {
	grpc.ServerStream
}

func (x *greeterManyHellosServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterManyHellosServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notification.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ManyHellos",
			Handler:       _Greeter_ManyHellos_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "notification.proto",
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_notification_95bceaf56281d715) }

var fileDescriptor_notification_95bceaf56281d715 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x53, 0x52, 0xe2, 0xe2, 0xf1, 0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0xd4, 0xb8, 0xb8, 0xa0, 0x6a, 0x0a, 0x72, 0x2a, 0x85, 0x24, 0xb8, 0xd8,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0x8a, 0x60, 0x5c, 0xa3, 0xa9, 0x8c, 0x5c, 0xec, 0xee,
	0x45, 0xa9, 0xa9, 0x25, 0xa9, 0x45, 0x42, 0x4e, 0x5c, 0x1c, 0xc1, 0x89, 0x95, 0x60, 0x6d, 0x42,
	0x52, 0x7a, 0x28, 0xce, 0x40, 0xb6, 0x4f, 0x4a, 0x02, 0xab, 0x5c, 0x41, 0x4e, 0xa5, 0x12, 0x83,
	0x90, 0x1b, 0x17, 0x97, 0x6f, 0x62, 0x1e, 0xc4, 0x90, 0x62, 0x72, 0x4d, 0xd1, 0x60, 0x74, 0x32,
	0xe3, 0x92, 0xce, 0xcc, 0xd7, 0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0x2d, 0xd6, 0xcb, 0x00, 0x29, 0x29, 0xcf, 0x2f, 0xca, 0x49, 0x71, 0xe2, 0x07, 0x2b, 0x0f,
	0x07, 0xb1, 0x03, 0x40, 0x21, 0x14, 0xc0, 0xb8, 0x88, 0x89, 0xd9, 0xc3, 0x27, 0x3c, 0x89, 0x0d,
	0x1c, 0x60, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x95, 0xad, 0xa5, 0x46, 0x01, 0x00,
	0x00,
}