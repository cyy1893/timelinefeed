// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: follow.proto

package follow

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

// FollowServiceClient is the client API for FollowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowServiceClient interface {
	// Get followers by following ID
	GetFollowersByFollowingID(ctx context.Context, in *FollowingIDRequest, opts ...grpc.CallOption) (*FollowersResponse, error)
}

type followServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowServiceClient(cc grpc.ClientConnInterface) FollowServiceClient {
	return &followServiceClient{cc}
}

func (c *followServiceClient) GetFollowersByFollowingID(ctx context.Context, in *FollowingIDRequest, opts ...grpc.CallOption) (*FollowersResponse, error) {
	out := new(FollowersResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/GetFollowersByFollowingID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServiceServer is the server API for FollowService service.
// All implementations must embed UnimplementedFollowServiceServer
// for forward compatibility
type FollowServiceServer interface {
	// Get followers by following ID
	GetFollowersByFollowingID(context.Context, *FollowingIDRequest) (*FollowersResponse, error)
	mustEmbedUnimplementedFollowServiceServer()
}

// UnimplementedFollowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowServiceServer struct {
}

func (UnimplementedFollowServiceServer) GetFollowersByFollowingID(context.Context, *FollowingIDRequest) (*FollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowersByFollowingID not implemented")
}
func (UnimplementedFollowServiceServer) mustEmbedUnimplementedFollowServiceServer() {}

// UnsafeFollowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServiceServer will
// result in compilation errors.
type UnsafeFollowServiceServer interface {
	mustEmbedUnimplementedFollowServiceServer()
}

func RegisterFollowServiceServer(s grpc.ServiceRegistrar, srv FollowServiceServer) {
	s.RegisterService(&FollowService_ServiceDesc, srv)
}

func _FollowService_GetFollowersByFollowingID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowingIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).GetFollowersByFollowingID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/GetFollowersByFollowingID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).GetFollowersByFollowingID(ctx, req.(*FollowingIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowService_ServiceDesc is the grpc.ServiceDesc for FollowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "follow.FollowService",
	HandlerType: (*FollowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFollowersByFollowingID",
			Handler:    _FollowService_GetFollowersByFollowingID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "follow.proto",
}
