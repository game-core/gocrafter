// ログインボーナス

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: login_bonus.proto

package loginBonus

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
	LoginBonus_GetUser_FullMethodName   = "/proto.LoginBonus/GetUser"
	LoginBonus_GetMaster_FullMethodName = "/proto.LoginBonus/GetMaster"
	LoginBonus_Receive_FullMethodName   = "/proto.LoginBonus/Receive"
)

// LoginBonusClient is the client API for LoginBonus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientMysqlConn.NewStream.
type LoginBonusClient interface {
	GetUser(ctx context.Context, in *LoginBonusGetUserRequest, opts ...grpc.CallOption) (*LoginBonusGetUserResponse, error)
	GetMaster(ctx context.Context, in *LoginBonusGetMasterRequest, opts ...grpc.CallOption) (*LoginBonusGetMasterResponse, error)
	Receive(ctx context.Context, in *LoginBonusReceiveRequest, opts ...grpc.CallOption) (*LoginBonusReceiveResponse, error)
}

type loginBonusClient struct {
	cc grpc.ClientMysqlConnInterface
}

func NewLoginBonusClient(cc grpc.ClientMysqlConnInterface) LoginBonusClient {
	return &loginBonusClient{cc}
}

func (c *loginBonusClient) GetUser(ctx context.Context, in *LoginBonusGetUserRequest, opts ...grpc.CallOption) (*LoginBonusGetUserResponse, error) {
	out := new(LoginBonusGetUserResponse)
	err := c.cc.Invoke(ctx, LoginBonus_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginBonusClient) GetMaster(ctx context.Context, in *LoginBonusGetMasterRequest, opts ...grpc.CallOption) (*LoginBonusGetMasterResponse, error) {
	out := new(LoginBonusGetMasterResponse)
	err := c.cc.Invoke(ctx, LoginBonus_GetMaster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginBonusClient) Receive(ctx context.Context, in *LoginBonusReceiveRequest, opts ...grpc.CallOption) (*LoginBonusReceiveResponse, error) {
	out := new(LoginBonusReceiveResponse)
	err := c.cc.Invoke(ctx, LoginBonus_Receive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginBonusServer is the server API for LoginBonus service.
// All implementations must embed UnimplementedLoginBonusServer
// for forward compatibility
type LoginBonusServer interface {
	GetUser(context.Context, *LoginBonusGetUserRequest) (*LoginBonusGetUserResponse, error)
	GetMaster(context.Context, *LoginBonusGetMasterRequest) (*LoginBonusGetMasterResponse, error)
	Receive(context.Context, *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error)
	mustEmbedUnimplementedLoginBonusServer()
}

// UnimplementedLoginBonusServer must be embedded to have forward compatible implementations.
type UnimplementedLoginBonusServer struct {
}

func (UnimplementedLoginBonusServer) GetUser(context.Context, *LoginBonusGetUserRequest) (*LoginBonusGetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedLoginBonusServer) GetMaster(context.Context, *LoginBonusGetMasterRequest) (*LoginBonusGetMasterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaster not implemented")
}
func (UnimplementedLoginBonusServer) Receive(context.Context, *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedLoginBonusServer) mustEmbedUnimplementedLoginBonusServer() {}

// UnsafeLoginBonusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginBonusServer will
// result in compilation errors.
type UnsafeLoginBonusServer interface {
	mustEmbedUnimplementedLoginBonusServer()
}

func RegisterLoginBonusServer(s grpc.ServiceRegistrar, srv LoginBonusServer) {
	s.RegisterService(&LoginBonus_ServiceDesc, srv)
}

func _LoginBonus_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginBonusGetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginBonusServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginBonus_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginBonusServer).GetUser(ctx, req.(*LoginBonusGetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginBonus_GetMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginBonusGetMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginBonusServer).GetMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginBonus_GetMaster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginBonusServer).GetMaster(ctx, req.(*LoginBonusGetMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginBonus_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginBonusReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginBonusServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginBonus_Receive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginBonusServer).Receive(ctx, req.(*LoginBonusReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginBonus_ServiceDesc is the grpc.ServiceDesc for LoginBonus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginBonus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LoginBonus",
	HandlerType: (*LoginBonusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _LoginBonus_GetUser_Handler,
		},
		{
			MethodName: "GetMaster",
			Handler:    _LoginBonus_GetMaster_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _LoginBonus_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login_bonus.proto",
}
