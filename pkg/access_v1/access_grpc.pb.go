// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: access.proto

package access_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Access_Check_FullMethodName = "/access_v1.Access/Check"
)

// AccessClient is the client API for Access service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type accessClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessClient(cc grpc.ClientConnInterface) AccessClient {
	return &accessClient{cc}
}

func (c *accessClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Access_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessServer is the server API for Access service.
// All implementations must embed UnimplementedAccessServer
// for forward compatibility.
type AccessServer interface {
	Check(context.Context, *CheckRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAccessServer()
}

// UnimplementedAccessServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccessServer struct{}

func (UnimplementedAccessServer) Check(context.Context, *CheckRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedAccessServer) mustEmbedUnimplementedAccessServer() {}
func (UnimplementedAccessServer) testEmbeddedByValue()                {}

// UnsafeAccessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessServer will
// result in compilation errors.
type UnsafeAccessServer interface {
	mustEmbedUnimplementedAccessServer()
}

func RegisterAccessServer(s grpc.ServiceRegistrar, srv AccessServer) {
	// If the following call pancis, it indicates UnimplementedAccessServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Access_ServiceDesc, srv)
}

func _Access_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Access_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Access_ServiceDesc is the grpc.ServiceDesc for Access service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Access_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "access_v1.Access",
	HandlerType: (*AccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Access_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access.proto",
}
