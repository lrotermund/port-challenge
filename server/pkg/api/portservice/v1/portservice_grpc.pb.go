// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortDomainClient is the client API for PortDomain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortDomainClient interface {
	CreatePort(ctx context.Context, in *CreatePortRequest, opts ...grpc.CallOption) (*CreatePortResponse, error)
	ReadPort(ctx context.Context, in *ReadPortRequest, opts ...grpc.CallOption) (*ReadPortResponse, error)
}

type portDomainClient struct {
	cc grpc.ClientConnInterface
}

func NewPortDomainClient(cc grpc.ClientConnInterface) PortDomainClient {
	return &portDomainClient{cc}
}

func (c *portDomainClient) CreatePort(ctx context.Context, in *CreatePortRequest, opts ...grpc.CallOption) (*CreatePortResponse, error) {
	out := new(CreatePortResponse)
	err := c.cc.Invoke(ctx, "/PortDomain/CreatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainClient) ReadPort(ctx context.Context, in *ReadPortRequest, opts ...grpc.CallOption) (*ReadPortResponse, error) {
	out := new(ReadPortResponse)
	err := c.cc.Invoke(ctx, "/PortDomain/ReadPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortDomainServer is the server API for PortDomain service.
// All implementations must embed UnimplementedPortDomainServer
// for forward compatibility
type PortDomainServer interface {
	CreatePort(context.Context, *CreatePortRequest) (*CreatePortResponse, error)
	ReadPort(context.Context, *ReadPortRequest) (*ReadPortResponse, error)
	mustEmbedUnimplementedPortDomainServer()
}

// UnimplementedPortDomainServer must be embedded to have forward compatible implementations.
type UnimplementedPortDomainServer struct {
}

func (UnimplementedPortDomainServer) CreatePort(context.Context, *CreatePortRequest) (*CreatePortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePort not implemented")
}
func (UnimplementedPortDomainServer) ReadPort(context.Context, *ReadPortRequest) (*ReadPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPort not implemented")
}
func (UnimplementedPortDomainServer) mustEmbedUnimplementedPortDomainServer() {}

// UnsafePortDomainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortDomainServer will
// result in compilation errors.
type UnsafePortDomainServer interface {
	mustEmbedUnimplementedPortDomainServer()
}

func RegisterPortDomainServer(s grpc.ServiceRegistrar, srv PortDomainServer) {
	s.RegisterService(&PortDomain_ServiceDesc, srv)
}

func _PortDomain_CreatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServer).CreatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortDomain/CreatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServer).CreatePort(ctx, req.(*CreatePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomain_ReadPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServer).ReadPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortDomain/ReadPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServer).ReadPort(ctx, req.(*ReadPortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortDomain_ServiceDesc is the grpc.ServiceDesc for PortDomain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortDomain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PortDomain",
	HandlerType: (*PortDomainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePort",
			Handler:    _PortDomain_CreatePort_Handler,
		},
		{
			MethodName: "ReadPort",
			Handler:    _PortDomain_ReadPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portservice.proto",
}
