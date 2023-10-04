// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/auth.proto

package proto

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	CreateAuth(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption) (*CreateAuthResponse, error)
	UpdateAuth(ctx context.Context, in *UpdateAuthRequest, opts ...grpc.CallOption) (*UpdateAuthResponse, error)
	DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error)
	CompareAuth(ctx context.Context, in *CompareAuthRequest, opts ...grpc.CallOption) (*CompareAuthResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateAuth(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption) (*CreateAuthResponse, error) {
	out := new(CreateAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CreateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateAuth(ctx context.Context, in *UpdateAuthRequest, opts ...grpc.CallOption) (*UpdateAuthResponse, error) {
	out := new(UpdateAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error) {
	out := new(DeleteAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DeleteAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CompareAuth(ctx context.Context, in *CompareAuthRequest, opts ...grpc.CallOption) (*CompareAuthResponse, error) {
	out := new(CompareAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CompareAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	CreateAuth(context.Context, *CreateAuthRequest) (*CreateAuthResponse, error)
	UpdateAuth(context.Context, *UpdateAuthRequest) (*UpdateAuthResponse, error)
	DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error)
	CompareAuth(context.Context, *CompareAuthRequest) (*CompareAuthResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) CreateAuth(context.Context, *CreateAuthRequest) (*CreateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuth not implemented")
}
func (UnimplementedAuthServiceServer) UpdateAuth(context.Context, *UpdateAuthRequest) (*UpdateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuth not implemented")
}
func (UnimplementedAuthServiceServer) DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuth not implemented")
}
func (UnimplementedAuthServiceServer) CompareAuth(context.Context, *CompareAuthRequest) (*CompareAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompareAuth not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_CreateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CreateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateAuth(ctx, req.(*CreateAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateAuth(ctx, req.(*UpdateAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DeleteAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteAuth(ctx, req.(*DeleteAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CompareAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompareAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CompareAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CompareAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CompareAuth(ctx, req.(*CompareAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuth",
			Handler:    _AuthService_CreateAuth_Handler,
		},
		{
			MethodName: "UpdateAuth",
			Handler:    _AuthService_UpdateAuth_Handler,
		},
		{
			MethodName: "DeleteAuth",
			Handler:    _AuthService_DeleteAuth_Handler,
		},
		{
			MethodName: "CompareAuth",
			Handler:    _AuthService_CompareAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}
