// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pgork.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	pgork.proto

It has these top-level messages:
	WriteReq
	WriteResp
	ReadReq
	ReadResp
*/
package service

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

type WriteReq struct {
	PipeName string `protobuf:"bytes,1,opt,name=pipeName" json:"pipeName,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WriteReq) Reset()                    { *m = WriteReq{} }
func (m *WriteReq) String() string            { return proto.CompactTextString(m) }
func (*WriteReq) ProtoMessage()               {}
func (*WriteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WriteReq) GetPipeName() string {
	if m != nil {
		return m.PipeName
	}
	return ""
}

func (m *WriteReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WriteResp struct {
}

func (m *WriteResp) Reset()                    { *m = WriteResp{} }
func (m *WriteResp) String() string            { return proto.CompactTextString(m) }
func (*WriteResp) ProtoMessage()               {}
func (*WriteResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ReadReq struct {
	PipeName string `protobuf:"bytes,1,opt,name=pipeName" json:"pipeName,omitempty"`
}

func (m *ReadReq) Reset()                    { *m = ReadReq{} }
func (m *ReadReq) String() string            { return proto.CompactTextString(m) }
func (*ReadReq) ProtoMessage()               {}
func (*ReadReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadReq) GetPipeName() string {
	if m != nil {
		return m.PipeName
	}
	return ""
}

type ReadResp struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ReadResp) Reset()                    { *m = ReadResp{} }
func (m *ReadResp) String() string            { return proto.CompactTextString(m) }
func (*ReadResp) ProtoMessage()               {}
func (*ReadResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadResp) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteReq)(nil), "service.WriteReq")
	proto.RegisterType((*WriteResp)(nil), "service.WriteResp")
	proto.RegisterType((*ReadReq)(nil), "service.ReadReq")
	proto.RegisterType((*ReadResp)(nil), "service.ReadResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pipe service

type PipeClient interface {
	Write(ctx context.Context, in *WriteReq, opts ...grpc.CallOption) (*WriteResp, error)
	Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (*ReadResp, error)
}

type pipeClient struct {
	cc *grpc.ClientConn
}

func NewPipeClient(cc *grpc.ClientConn) PipeClient {
	return &pipeClient{cc}
}

func (c *pipeClient) Write(ctx context.Context, in *WriteReq, opts ...grpc.CallOption) (*WriteResp, error) {
	out := new(WriteResp)
	err := grpc.Invoke(ctx, "/service.Pipe/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeClient) Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (*ReadResp, error) {
	out := new(ReadResp)
	err := grpc.Invoke(ctx, "/service.Pipe/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pipe service

type PipeServer interface {
	Write(context.Context, *WriteReq) (*WriteResp, error)
	Read(context.Context, *ReadReq) (*ReadResp, error)
}

func RegisterPipeServer(s *grpc.Server, srv PipeServer) {
	s.RegisterService(&_Pipe_serviceDesc, srv)
}

func _Pipe_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Pipe/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).Write(ctx, req.(*WriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipe_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Pipe/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).Read(ctx, req.(*ReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pipe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Pipe",
	HandlerType: (*PipeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Pipe_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Pipe_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pgork.proto",
}

func init() { proto.RegisterFile("pgork.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x48, 0xcf, 0x2f,
	0xca, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x55, 0xb2, 0xe2, 0xe2, 0x08, 0x2f, 0xca, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x14, 0x92, 0xe2, 0xe2,
	0x28, 0xc8, 0x2c, 0x48, 0xf5, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0xf3, 0x85, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82,
	0xc0, 0x6c, 0x25, 0x6e, 0x2e, 0x4e, 0xa8, 0xde, 0xe2, 0x02, 0x25, 0x55, 0x2e, 0xf6, 0xa0, 0xd4,
	0xc4, 0x14, 0x02, 0xe6, 0x28, 0xc9, 0x71, 0x71, 0x40, 0x94, 0x15, 0x17, 0xc0, 0xcd, 0x64, 0x44,
	0x98, 0x69, 0x94, 0xce, 0xc5, 0x12, 0x90, 0x59, 0x90, 0x2a, 0x64, 0xc0, 0xc5, 0x0a, 0x36, 0x5b,
	0x48, 0x50, 0x0f, 0xea, 0x54, 0x3d, 0x98, 0x3b, 0xa5, 0x84, 0xd0, 0x85, 0x8a, 0x0b, 0x94, 0x18,
	0x84, 0x74, 0xb9, 0x58, 0x40, 0x26, 0x0b, 0x09, 0xc0, 0x65, 0xa1, 0xee, 0x91, 0x12, 0x44, 0x13,
	0x01, 0x29, 0x4f, 0x62, 0x03, 0x07, 0x84, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x81, 0x9f, 0x7b,
	0xaf, 0x17, 0x01, 0x00, 0x00,
}
