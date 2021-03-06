// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeatConnect.proto

package proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the request
type OutputRequest struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputRequest) Reset()         { *m = OutputRequest{} }
func (m *OutputRequest) String() string { return proto.CompactTextString(m) }
func (*OutputRequest) ProtoMessage()    {}
func (*OutputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac6fcabb4a7d182, []int{0}
}

func (m *OutputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputRequest.Unmarshal(m, b)
}
func (m *OutputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputRequest.Marshal(b, m, deterministic)
}
func (m *OutputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputRequest.Merge(m, src)
}
func (m *OutputRequest) XXX_Size() int {
	return xxx_messageInfo_OutputRequest.Size(m)
}
func (m *OutputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OutputRequest proto.InternalMessageInfo

func (m *OutputRequest) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

// The response message containing the greetings
type OutputReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputReply) Reset()         { *m = OutputReply{} }
func (m *OutputReply) String() string { return proto.CompactTextString(m) }
func (*OutputReply) ProtoMessage()    {}
func (*OutputReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac6fcabb4a7d182, []int{1}
}

func (m *OutputReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputReply.Unmarshal(m, b)
}
func (m *OutputReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputReply.Marshal(b, m, deterministic)
}
func (m *OutputReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputReply.Merge(m, src)
}
func (m *OutputReply) XXX_Size() int {
	return xxx_messageInfo_OutputReply.Size(m)
}
func (m *OutputReply) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputReply.DiscardUnknown(m)
}

var xxx_messageInfo_OutputReply proto.InternalMessageInfo

func (m *OutputReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*OutputRequest)(nil), "proto.OutputRequest")
	proto.RegisterType((*OutputReply)(nil), "proto.OutputReply")
}

func init() { proto.RegisterFile("heartbeatConnect.proto", fileDescriptor_8ac6fcabb4a7d182) }

var fileDescriptor_8ac6fcabb4a7d182 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0x71, 0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xea, 0x5c, 0xbc, 0xfe, 0xa5, 0x25, 0x05, 0xa5, 0x25,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0xf9, 0x60, 0x01, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x49, 0x9d, 0x8b, 0x1b, 0xa6, 0xb0, 0x20, 0xa7, 0x52,
	0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x15, 0xaa, 0x0e, 0xc6, 0x35, 0x6a,
	0x60, 0xe4, 0xe2, 0xf1, 0x48, 0xcc, 0x4b, 0xc9, 0x49, 0x85, 0xa8, 0x17, 0xb2, 0xe0, 0xe2, 0x0a,
	0x4e, 0xcd, 0x4b, 0x81, 0xf2, 0x44, 0x20, 0xf6, 0xeb, 0xa1, 0xd8, 0x2a, 0x25, 0x84, 0x26, 0x5a,
	0x90, 0x53, 0xa9, 0xc4, 0x20, 0x64, 0xc6, 0xc5, 0x01, 0xd2, 0xe9, 0x93, 0x9f, 0x5e, 0x4c, 0x8a,
	0xbe, 0x24, 0x36, 0xb0, 0xa0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd7, 0xfc, 0x53, 0xfc,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HandleOutputClient is the client API for HandleOutput service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandleOutputClient interface {
	// Sends a greeting
	SendOutput(ctx context.Context, in *OutputRequest, opts ...grpc.CallOption) (*OutputReply, error)
	SendLogs(ctx context.Context, in *OutputRequest, opts ...grpc.CallOption) (*OutputReply, error)
}

type handleOutputClient struct {
	cc *grpc.ClientConn
}

func NewHandleOutputClient(cc *grpc.ClientConn) HandleOutputClient {
	return &handleOutputClient{cc}
}

func (c *handleOutputClient) SendOutput(ctx context.Context, in *OutputRequest, opts ...grpc.CallOption) (*OutputReply, error) {
	out := new(OutputReply)
	err := c.cc.Invoke(ctx, "/proto.HandleOutput/SendOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handleOutputClient) SendLogs(ctx context.Context, in *OutputRequest, opts ...grpc.CallOption) (*OutputReply, error) {
	out := new(OutputReply)
	err := c.cc.Invoke(ctx, "/proto.HandleOutput/SendLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandleOutputServer is the server API for HandleOutput service.
type HandleOutputServer interface {
	// Sends a greeting
	SendOutput(context.Context, *OutputRequest) (*OutputReply, error)
	SendLogs(context.Context, *OutputRequest) (*OutputReply, error)
}

func RegisterHandleOutputServer(s *grpc.Server, srv HandleOutputServer) {
	s.RegisterService(&_HandleOutput_serviceDesc, srv)
}

func _HandleOutput_SendOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleOutputServer).SendOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HandleOutput/SendOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleOutputServer).SendOutput(ctx, req.(*OutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandleOutput_SendLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleOutputServer).SendLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HandleOutput/SendLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleOutputServer).SendLogs(ctx, req.(*OutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleOutput_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HandleOutput",
	HandlerType: (*HandleOutputServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOutput",
			Handler:    _HandleOutput_SendOutput_Handler,
		},
		{
			MethodName: "SendLogs",
			Handler:    _HandleOutput_SendLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeatConnect.proto",
}
