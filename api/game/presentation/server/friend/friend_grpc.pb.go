// フレンド

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: friend.proto

package friend

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
	Friend_Get_FullMethodName        = "/proto.Friend/Get"
	Friend_Send_FullMethodName       = "/proto.Friend/Send"
	Friend_Approve_FullMethodName    = "/proto.Friend/Approve"
	Friend_Disapprove_FullMethodName = "/proto.Friend/Disapprove"
	Friend_Delete_FullMethodName     = "/proto.Friend/Delete"
)

// FriendClient is the client API for Friend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendClient interface {
	Get(ctx context.Context, in *FriendGetRequest, opts ...grpc.CallOption) (*FriendGetResponse, error)
	Send(ctx context.Context, in *FriendSendRequest, opts ...grpc.CallOption) (*FriendSendResponse, error)
	Approve(ctx context.Context, in *FriendApproveRequest, opts ...grpc.CallOption) (*FriendApproveResponse, error)
	Disapprove(ctx context.Context, in *FriendDisapproveRequest, opts ...grpc.CallOption) (*FriendDisapproveResponse, error)
	Delete(ctx context.Context, in *FriendDeleteRequest, opts ...grpc.CallOption) (*FriendDeleteResponse, error)
}

type friendClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendClient(cc grpc.ClientConnInterface) FriendClient {
	return &friendClient{cc}
}

func (c *friendClient) Get(ctx context.Context, in *FriendGetRequest, opts ...grpc.CallOption) (*FriendGetResponse, error) {
	out := new(FriendGetResponse)
	err := c.cc.Invoke(ctx, Friend_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) Send(ctx context.Context, in *FriendSendRequest, opts ...grpc.CallOption) (*FriendSendResponse, error) {
	out := new(FriendSendResponse)
	err := c.cc.Invoke(ctx, Friend_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) Approve(ctx context.Context, in *FriendApproveRequest, opts ...grpc.CallOption) (*FriendApproveResponse, error) {
	out := new(FriendApproveResponse)
	err := c.cc.Invoke(ctx, Friend_Approve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) Disapprove(ctx context.Context, in *FriendDisapproveRequest, opts ...grpc.CallOption) (*FriendDisapproveResponse, error) {
	out := new(FriendDisapproveResponse)
	err := c.cc.Invoke(ctx, Friend_Disapprove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) Delete(ctx context.Context, in *FriendDeleteRequest, opts ...grpc.CallOption) (*FriendDeleteResponse, error) {
	out := new(FriendDeleteResponse)
	err := c.cc.Invoke(ctx, Friend_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServer is the server API for Friend service.
// All implementations must embed UnimplementedFriendServer
// for forward compatibility
type FriendServer interface {
	Get(context.Context, *FriendGetRequest) (*FriendGetResponse, error)
	Send(context.Context, *FriendSendRequest) (*FriendSendResponse, error)
	Approve(context.Context, *FriendApproveRequest) (*FriendApproveResponse, error)
	Disapprove(context.Context, *FriendDisapproveRequest) (*FriendDisapproveResponse, error)
	Delete(context.Context, *FriendDeleteRequest) (*FriendDeleteResponse, error)
	mustEmbedUnimplementedFriendServer()
}

// UnimplementedFriendServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServer struct {
}

func (UnimplementedFriendServer) Get(context.Context, *FriendGetRequest) (*FriendGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFriendServer) Send(context.Context, *FriendSendRequest) (*FriendSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedFriendServer) Approve(context.Context, *FriendApproveRequest) (*FriendApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}
func (UnimplementedFriendServer) Disapprove(context.Context, *FriendDisapproveRequest) (*FriendDisapproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disapprove not implemented")
}
func (UnimplementedFriendServer) Delete(context.Context, *FriendDeleteRequest) (*FriendDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFriendServer) mustEmbedUnimplementedFriendServer() {}

// UnsafeFriendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServer will
// result in compilation errors.
type UnsafeFriendServer interface {
	mustEmbedUnimplementedFriendServer()
}

func RegisterFriendServer(s grpc.ServiceRegistrar, srv FriendServer) {
	s.RegisterService(&Friend_ServiceDesc, srv)
}

func _Friend_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).Get(ctx, req.(*FriendGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).Send(ctx, req.(*FriendSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_Approve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).Approve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_Approve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).Approve(ctx, req.(*FriendApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_Disapprove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendDisapproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).Disapprove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_Disapprove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).Disapprove(ctx, req.(*FriendDisapproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).Delete(ctx, req.(*FriendDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Friend_ServiceDesc is the grpc.ServiceDesc for Friend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Friend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Friend",
	HandlerType: (*FriendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Friend_Get_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Friend_Send_Handler,
		},
		{
			MethodName: "Approve",
			Handler:    _Friend_Approve_Handler,
		},
		{
			MethodName: "Disapprove",
			Handler:    _Friend_Disapprove_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Friend_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend.proto",
}
