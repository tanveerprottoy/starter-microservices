// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: src/module/content/proto/content.proto

package proto

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

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	CreateContent(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Content, error)
	ReadContents(ctx context.Context, in *VoidParam, opts ...grpc.CallOption) (*Contents, error)
	ReadContentStream(ctx context.Context, in *VoidParam, opts ...grpc.CallOption) (ContentService_ReadContentStreamClient, error)
	ReadContent(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Content, error)
	UpdateContent(ctx context.Context, in *UpdateContentParam, opts ...grpc.CallOption) (*Content, error)
	DeleteContent(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) CreateContent(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/contentPackage.ContentService/createContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) ReadContents(ctx context.Context, in *VoidParam, opts ...grpc.CallOption) (*Contents, error) {
	out := new(Contents)
	err := c.cc.Invoke(ctx, "/contentPackage.ContentService/readContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) ReadContentStream(ctx context.Context, in *VoidParam, opts ...grpc.CallOption) (ContentService_ReadContentStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ContentService_ServiceDesc.Streams[0], "/contentPackage.ContentService/readContentStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentServiceReadContentStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ContentService_ReadContentStreamClient interface {
	Recv() (*Content, error)
	grpc.ClientStream
}

type contentServiceReadContentStreamClient struct {
	grpc.ClientStream
}

func (x *contentServiceReadContentStreamClient) Recv() (*Content, error) {
	m := new(Content)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentServiceClient) ReadContent(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/contentPackage.ContentService/readContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UpdateContent(ctx context.Context, in *UpdateContentParam, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/contentPackage.ContentService/updateContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteContent(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/contentPackage.ContentService/deleteContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	CreateContent(context.Context, *Content) (*Content, error)
	ReadContents(context.Context, *VoidParam) (*Contents, error)
	ReadContentStream(*VoidParam, ContentService_ReadContentStreamServer) error
	ReadContent(context.Context, *wrapperspb.StringValue) (*Content, error)
	UpdateContent(context.Context, *UpdateContentParam) (*Content, error)
	DeleteContent(context.Context, *wrapperspb.StringValue) (*wrapperspb.BoolValue, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) CreateContent(context.Context, *Content) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContent not implemented")
}
func (UnimplementedContentServiceServer) ReadContents(context.Context, *VoidParam) (*Contents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadContents not implemented")
}
func (UnimplementedContentServiceServer) ReadContentStream(*VoidParam, ContentService_ReadContentStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadContentStream not implemented")
}
func (UnimplementedContentServiceServer) ReadContent(context.Context, *wrapperspb.StringValue) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadContent not implemented")
}
func (UnimplementedContentServiceServer) UpdateContent(context.Context, *UpdateContentParam) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContent not implemented")
}
func (UnimplementedContentServiceServer) DeleteContent(context.Context, *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContent not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_CreateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).CreateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentPackage.ContentService/createContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).CreateContent(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_ReadContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).ReadContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentPackage.ContentService/readContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).ReadContents(ctx, req.(*VoidParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_ReadContentStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VoidParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContentServiceServer).ReadContentStream(m, &contentServiceReadContentStreamServer{stream})
}

type ContentService_ReadContentStreamServer interface {
	Send(*Content) error
	grpc.ServerStream
}

type contentServiceReadContentStreamServer struct {
	grpc.ServerStream
}

func (x *contentServiceReadContentStreamServer) Send(m *Content) error {
	return x.ServerStream.SendMsg(m)
}

func _ContentService_ReadContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).ReadContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentPackage.ContentService/readContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).ReadContent(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_UpdateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContentParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).UpdateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentPackage.ContentService/updateContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).UpdateContent(ctx, req.(*UpdateContentParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentPackage.ContentService/deleteContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteContent(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contentPackage.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createContent",
			Handler:    _ContentService_CreateContent_Handler,
		},
		{
			MethodName: "readContents",
			Handler:    _ContentService_ReadContents_Handler,
		},
		{
			MethodName: "readContent",
			Handler:    _ContentService_ReadContent_Handler,
		},
		{
			MethodName: "updateContent",
			Handler:    _ContentService_UpdateContent_Handler,
		},
		{
			MethodName: "deleteContent",
			Handler:    _ContentService_DeleteContent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "readContentStream",
			Handler:       _ContentService_ReadContentStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/module/content/proto/content.proto",
}
