// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
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

type Request struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Request) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type ResponseInt struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseInt) Reset()         { *m = ResponseInt{} }
func (m *ResponseInt) String() string { return proto.CompactTextString(m) }
func (*ResponseInt) ProtoMessage()    {}
func (*ResponseInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ResponseInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseInt.Unmarshal(m, b)
}
func (m *ResponseInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseInt.Marshal(b, m, deterministic)
}
func (m *ResponseInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseInt.Merge(m, src)
}
func (m *ResponseInt) XXX_Size() int {
	return xxx_messageInfo_ResponseInt.Size(m)
}
func (m *ResponseInt) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseInt.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseInt proto.InternalMessageInfo

func (m *ResponseInt) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ResponseFloat struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseFloat) Reset()         { *m = ResponseFloat{} }
func (m *ResponseFloat) String() string { return proto.CompactTextString(m) }
func (*ResponseFloat) ProtoMessage()    {}
func (*ResponseFloat) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *ResponseFloat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseFloat.Unmarshal(m, b)
}
func (m *ResponseFloat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseFloat.Marshal(b, m, deterministic)
}
func (m *ResponseFloat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseFloat.Merge(m, src)
}
func (m *ResponseFloat) XXX_Size() int {
	return xxx_messageInfo_ResponseFloat.Size(m)
}
func (m *ResponseFloat) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseFloat.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseFloat proto.InternalMessageInfo

func (m *ResponseFloat) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*ResponseInt)(nil), "proto.ResponseInt")
	proto.RegisterType((*ResponseFloat)(nil), "proto.ResponseFloat")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xaa, 0x5c,
	0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x3c, 0x5c, 0x8c, 0x89, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xcc, 0x41, 0x8c, 0x89, 0x20, 0x5e, 0x92, 0x04, 0x13, 0x84, 0x97, 0xa4, 0xa4, 0xca,
	0xc5, 0x1d, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xea, 0x99, 0x57, 0x22, 0x24, 0xc6, 0xc5,
	0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0x55, 0x0f, 0xe5, 0x29, 0xa9, 0x73, 0xf1, 0xc2, 0x94,
	0xb9, 0xe5, 0xe4, 0x27, 0xa2, 0x2b, 0x64, 0x82, 0x29, 0x34, 0x9a, 0xc9, 0xc8, 0xc5, 0xe5, 0x98,
	0x92, 0x12, 0x0c, 0x71, 0x92, 0x90, 0x26, 0x17, 0xb3, 0x63, 0x4a, 0x8a, 0x10, 0x1f, 0xc4, 0x6d,
	0x7a, 0x50, 0x17, 0x49, 0x09, 0xc1, 0xf9, 0x08, 0xab, 0xf5, 0xb8, 0x38, 0x7c, 0x4b, 0x73, 0x4a,
	0x32, 0x0b, 0x72, 0x2a, 0x89, 0x52, 0x6f, 0xc0, 0xc5, 0xe1, 0x92, 0x59, 0x96, 0x59, 0x9c, 0x99,
	0x9f, 0x87, 0xa1, 0x5e, 0x04, 0x4d, 0x3d, 0xd8, 0xcd, 0x49, 0x6c, 0x60, 0x41, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0xf5, 0xa5, 0x9b, 0x31, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseInt, error)
	Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseInt, error)
	Division(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseFloat, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseInt, error) {
	out := new(ResponseInt)
	err := c.cc.Invoke(ctx, "/proto.AddService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseInt, error) {
	out := new(ResponseInt)
	err := c.cc.Invoke(ctx, "/proto.AddService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Division(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseFloat, error) {
	out := new(ResponseFloat)
	err := c.cc.Invoke(ctx, "/proto.AddService/Division", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	Add(context.Context, *Request) (*ResponseInt, error)
	Multiply(context.Context, *Request) (*ResponseInt, error)
	Division(context.Context, *Request) (*ResponseFloat, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) Add(ctx context.Context, req *Request) (*ResponseInt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedAddServiceServer) Multiply(ctx context.Context, req *Request) (*ResponseInt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (*UnimplementedAddServiceServer) Division(ctx context.Context, req *Request) (*ResponseFloat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Division not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Multiply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Division_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Division(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Division",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Division(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddService_Add_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _AddService_Multiply_Handler,
		},
		{
			MethodName: "Division",
			Handler:    _AddService_Division_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}