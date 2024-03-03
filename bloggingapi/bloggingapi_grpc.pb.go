// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: bloggingapi/bloggingapi.proto

package bloggingapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Bloggingapi_CreateBlog_FullMethodName = "/bloggingapi.bloggingapi/CreateBlog"
	Bloggingapi_ReadBlog_FullMethodName   = "/bloggingapi.bloggingapi/ReadBlog"
	Bloggingapi_UpdateBlog_FullMethodName = "/bloggingapi.bloggingapi/UpdateBlog"
	Bloggingapi_DeleteBlog_FullMethodName = "/bloggingapi.bloggingapi/DeleteBlog"
)

// BloggingapiClient is the client API for Bloggingapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BloggingapiClient interface {
	CreateBlog(ctx context.Context, in *BlogPost, opts ...grpc.CallOption) (*BlogPostWithUid, error)
	ReadBlog(ctx context.Context, in *BlogRequest, opts ...grpc.CallOption) (*BlogPostWithUid, error)
	UpdateBlog(ctx context.Context, in *BlogPostWithUid, opts ...grpc.CallOption) (*BlogPost, error)
	DeleteBlog(ctx context.Context, in *BlogRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type bloggingapiClient struct {
	cc grpc.ClientConnInterface
}

func NewBloggingapiClient(cc grpc.ClientConnInterface) BloggingapiClient {
	return &bloggingapiClient{cc}
}

func (c *bloggingapiClient) CreateBlog(ctx context.Context, in *BlogPost, opts ...grpc.CallOption) (*BlogPostWithUid, error) {
	out := new(BlogPostWithUid)
	err := c.cc.Invoke(ctx, Bloggingapi_CreateBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloggingapiClient) ReadBlog(ctx context.Context, in *BlogRequest, opts ...grpc.CallOption) (*BlogPostWithUid, error) {
	out := new(BlogPostWithUid)
	err := c.cc.Invoke(ctx, Bloggingapi_ReadBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloggingapiClient) UpdateBlog(ctx context.Context, in *BlogPostWithUid, opts ...grpc.CallOption) (*BlogPost, error) {
	out := new(BlogPost)
	err := c.cc.Invoke(ctx, Bloggingapi_UpdateBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloggingapiClient) DeleteBlog(ctx context.Context, in *BlogRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, Bloggingapi_DeleteBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BloggingapiServer is the server API for Bloggingapi service.
// All implementations must embed UnimplementedBloggingapiServer
// for forward compatibility
type BloggingapiServer interface {
	CreateBlog(context.Context, *BlogPost) (*BlogPostWithUid, error)
	ReadBlog(context.Context, *BlogRequest) (*BlogPostWithUid, error)
	UpdateBlog(context.Context, *BlogPostWithUid) (*BlogPost, error)
	DeleteBlog(context.Context, *BlogRequest) (*wrapperspb.BoolValue, error)
	mustEmbedUnimplementedBloggingapiServer()
}

// UnimplementedBloggingapiServer must be embedded to have forward compatible implementations.
type UnimplementedBloggingapiServer struct {
}

func (UnimplementedBloggingapiServer) CreateBlog(context.Context, *BlogPost) (*BlogPostWithUid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBloggingapiServer) ReadBlog(context.Context, *BlogRequest) (*BlogPostWithUid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (UnimplementedBloggingapiServer) UpdateBlog(context.Context, *BlogPostWithUid) (*BlogPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (UnimplementedBloggingapiServer) DeleteBlog(context.Context, *BlogRequest) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (UnimplementedBloggingapiServer) mustEmbedUnimplementedBloggingapiServer() {}

// UnsafeBloggingapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BloggingapiServer will
// result in compilation errors.
type UnsafeBloggingapiServer interface {
	mustEmbedUnimplementedBloggingapiServer()
}

func RegisterBloggingapiServer(s grpc.ServiceRegistrar, srv BloggingapiServer) {
	s.RegisterService(&Bloggingapi_ServiceDesc, srv)
}

func _Bloggingapi_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloggingapiServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bloggingapi_CreateBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloggingapiServer).CreateBlog(ctx, req.(*BlogPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bloggingapi_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloggingapiServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bloggingapi_ReadBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloggingapiServer).ReadBlog(ctx, req.(*BlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bloggingapi_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogPostWithUid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloggingapiServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bloggingapi_UpdateBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloggingapiServer).UpdateBlog(ctx, req.(*BlogPostWithUid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bloggingapi_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloggingapiServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bloggingapi_DeleteBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloggingapiServer).DeleteBlog(ctx, req.(*BlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bloggingapi_ServiceDesc is the grpc.ServiceDesc for Bloggingapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bloggingapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bloggingapi.bloggingapi",
	HandlerType: (*BloggingapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _Bloggingapi_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _Bloggingapi_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _Bloggingapi_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _Bloggingapi_DeleteBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bloggingapi/bloggingapi.proto",
}
