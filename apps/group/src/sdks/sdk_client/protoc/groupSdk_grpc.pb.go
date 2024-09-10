// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: groupSdk.proto

package protoc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GroupSdkService_GetMember_FullMethodName = "/groupSdk.GroupSdkService/GetMember"
)

// GroupSdkServiceClient is the client API for GroupSdkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupSdkServiceClient interface {
	GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*GetMemberResponse, error)
}

type groupSdkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupSdkServiceClient(cc grpc.ClientConnInterface) GroupSdkServiceClient {
	return &groupSdkServiceClient{cc}
}

func (c *groupSdkServiceClient) GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*GetMemberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMemberResponse)
	err := c.cc.Invoke(ctx, GroupSdkService_GetMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupSdkServiceServer is the server API for GroupSdkService service.
// All implementations should embed UnimplementedGroupSdkServiceServer
// for forward compatibility.
type GroupSdkServiceServer interface {
	GetMember(context.Context, *GetMemberRequest) (*GetMemberResponse, error)
}

// UnimplementedGroupSdkServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGroupSdkServiceServer struct{}

func (UnimplementedGroupSdkServiceServer) GetMember(context.Context, *GetMemberRequest) (*GetMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (UnimplementedGroupSdkServiceServer) testEmbeddedByValue() {}

// UnsafeGroupSdkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupSdkServiceServer will
// result in compilation errors.
type UnsafeGroupSdkServiceServer interface {
	mustEmbedUnimplementedGroupSdkServiceServer()
}

func RegisterGroupSdkServiceServer(s grpc.ServiceRegistrar, srv GroupSdkServiceServer) {
	// If the following call pancis, it indicates UnimplementedGroupSdkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GroupSdkService_ServiceDesc, srv)
}

func _GroupSdkService_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSdkServiceServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupSdkService_GetMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSdkServiceServer).GetMember(ctx, req.(*GetMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupSdkService_ServiceDesc is the grpc.ServiceDesc for GroupSdkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupSdkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "groupSdk.GroupSdkService",
	HandlerType: (*GroupSdkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMember",
			Handler:    _GroupSdkService_GetMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "groupSdk.proto",
}
