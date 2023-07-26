// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: iritamod/identity/tx.proto

package identity

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
	Msg_CreateIdentity_FullMethodName = "/iritamod.identity.Msg/CreateIdentity"
	Msg_UpdateIdentity_FullMethodName = "/iritamod.identity.Msg/UpdateIdentity"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateIdentity defines a method for creating a new identity.
	CreateIdentity(ctx context.Context, in *MsgCreateIdentity, opts ...grpc.CallOption) (*MsgCreateIdentityResponse, error)
	// UpdateIdentity defines a method for Updating a identity.
	UpdateIdentity(ctx context.Context, in *MsgUpdateIdentity, opts ...grpc.CallOption) (*MsgUpdateIdentityResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateIdentity(ctx context.Context, in *MsgCreateIdentity, opts ...grpc.CallOption) (*MsgCreateIdentityResponse, error) {
	out := new(MsgCreateIdentityResponse)
	err := c.cc.Invoke(ctx, Msg_CreateIdentity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateIdentity(ctx context.Context, in *MsgUpdateIdentity, opts ...grpc.CallOption) (*MsgUpdateIdentityResponse, error) {
	out := new(MsgUpdateIdentityResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateIdentity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateIdentity defines a method for creating a new identity.
	CreateIdentity(context.Context, *MsgCreateIdentity) (*MsgCreateIdentityResponse, error)
	// UpdateIdentity defines a method for Updating a identity.
	UpdateIdentity(context.Context, *MsgUpdateIdentity) (*MsgUpdateIdentityResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateIdentity(context.Context, *MsgCreateIdentity) (*MsgCreateIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIdentity not implemented")
}
func (UnimplementedMsgServer) UpdateIdentity(context.Context, *MsgUpdateIdentity) (*MsgUpdateIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIdentity not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateIdentity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateIdentity(ctx, req.(*MsgCreateIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateIdentity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateIdentity(ctx, req.(*MsgUpdateIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iritamod.identity.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIdentity",
			Handler:    _Msg_CreateIdentity_Handler,
		},
		{
			MethodName: "UpdateIdentity",
			Handler:    _Msg_UpdateIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iritamod/identity/tx.proto",
}