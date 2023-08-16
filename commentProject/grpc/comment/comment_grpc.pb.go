// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: comment.proto

package comment

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

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	GetCommentCountByFeedID(ctx context.Context, in *CommentCountRequest, opts ...grpc.CallOption) (*CommentCountResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) GetCommentCountByFeedID(ctx context.Context, in *CommentCountRequest, opts ...grpc.CallOption) (*CommentCountResponse, error) {
	out := new(CommentCountResponse)
	err := c.cc.Invoke(ctx, "/comment.CommentService/GetCommentCountByFeedID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	GetCommentCountByFeedID(context.Context, *CommentCountRequest) (*CommentCountResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) GetCommentCountByFeedID(context.Context, *CommentCountRequest) (*CommentCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentCountByFeedID not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_GetCommentCountByFeedID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentCountByFeedID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.CommentService/GetCommentCountByFeedID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentCountByFeedID(ctx, req.(*CommentCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommentCountByFeedID",
			Handler:    _CommentService_GetCommentCountByFeedID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
