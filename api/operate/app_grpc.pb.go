// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: operate/app.proto

package operate

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

const (
	App_CreateContent_FullMethodName = "/api.operate.App/CreateContent"
	App_UpdateContent_FullMethodName = "/api.operate.App/UpdateContent"
)

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	// 创建内容
	CreateContent(ctx context.Context, in *CreateContentReq, opts ...grpc.CallOption) (*CreateContentRsp, error)
	// 内容更新
	UpdateContent(ctx context.Context, in *UpdateContentReq, opts ...grpc.CallOption) (*UpdateContentRsp, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) CreateContent(ctx context.Context, in *CreateContentReq, opts ...grpc.CallOption) (*CreateContentRsp, error) {
	out := new(CreateContentRsp)
	err := c.cc.Invoke(ctx, App_CreateContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) UpdateContent(ctx context.Context, in *UpdateContentReq, opts ...grpc.CallOption) (*UpdateContentRsp, error) {
	out := new(UpdateContentRsp)
	err := c.cc.Invoke(ctx, App_UpdateContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	// 创建内容
	CreateContent(context.Context, *CreateContentReq) (*CreateContentRsp, error)
	// 内容更新
	UpdateContent(context.Context, *UpdateContentReq) (*UpdateContentRsp, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) CreateContent(context.Context, *CreateContentReq) (*CreateContentRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContent not implemented")
}
func (UnimplementedAppServer) UpdateContent(context.Context, *UpdateContentReq) (*UpdateContentRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContent not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_CreateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).CreateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_CreateContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).CreateContent(ctx, req.(*CreateContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_UpdateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).UpdateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_UpdateContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).UpdateContent(ctx, req.(*UpdateContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.operate.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContent",
			Handler:    _App_CreateContent_Handler,
		},
		{
			MethodName: "UpdateContent",
			Handler:    _App_UpdateContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operate/app.proto",
}
