// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PingRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PingResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "test.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "test.PingResponse")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0xc4, 0xcb, 0x12, 0x73, 0x32,
	0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0xb4, 0x92, 0x06, 0x17, 0x77, 0x40, 0x66, 0x5e,
	0x7a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x24, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0xb3, 0x13, 0xe7, 0x2f, 0x27, 0x36, 0x25, 0x16, 0x81, 0x14, 0x05, 0x86, 0x20,
	0xa6, 0xcc, 0x14, 0x25, 0x4d, 0x2e, 0x1e, 0x88, 0xca, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x3c,
	0x4a, 0x8d, 0xca, 0xb9, 0x38, 0x7d, 0x2b, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xf4,
	0xb9, 0x58, 0x40, 0xfa, 0x84, 0x04, 0xf5, 0xc0, 0xae, 0x42, 0xb2, 0x4d, 0x4a, 0x08, 0x59, 0x08,
	0x62, 0xac, 0x12, 0x83, 0x90, 0x25, 0x17, 0x17, 0x48, 0x24, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0x97,
	0x68, 0x6d, 0x1a, 0x8c, 0x06, 0x8c, 0x4e, 0x6c, 0x51, 0x60, 0xef, 0x26, 0xb1, 0x81, 0x3d, 0x67,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x1a, 0x49, 0xeb, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingStream(ctx context.Context, opts ...grpc.CallOption) (MyService_PingStreamClient, error)
}

type myServiceClient struct {
	cc *grpc.ClientConn
}

func NewMyServiceClient(cc *grpc.ClientConn) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/test.MyService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) PingStream(ctx context.Context, opts ...grpc.CallOption) (MyService_PingStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MyService_serviceDesc.Streams[0], "/test.MyService/PingStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &myServicePingStreamClient{stream}
	return x, nil
}

type MyService_PingStreamClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type myServicePingStreamClient struct {
	grpc.ClientStream
}

func (x *myServicePingStreamClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myServicePingStreamClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyServiceServer is the server API for MyService service.
type MyServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingStream(MyService_PingStreamServer) error
}

// UnimplementedMyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (*UnimplementedMyServiceServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedMyServiceServer) PingStream(srv MyService_PingStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PingStream not implemented")
}

func RegisterMyServiceServer(s *grpc.Server, srv MyServiceServer) {
	s.RegisterService(&_MyService_serviceDesc, srv)
}

func _MyService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.MyService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_PingStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyServiceServer).PingStream(&myServicePingStreamServer{stream})
}

type MyService_PingStreamServer interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type myServicePingStreamServer struct {
	grpc.ServerStream
}

func (x *myServicePingStreamServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myServicePingStreamServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _MyService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingStream",
			Handler:       _MyService_PingStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}