// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: iritamod/side-chain/v1/tx.proto

package side_chainv1

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
	Msg_CreateSpace_FullMethodName       = "/iritamod.side_chain.v1.Msg/CreateSpace"
	Msg_TransferSpace_FullMethodName     = "/iritamod.side_chain.v1.Msg/TransferSpace"
	Msg_CreateBlockHeader_FullMethodName = "/iritamod.side_chain.v1.Msg/CreateBlockHeader"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateSpace defines a method for creating a space
	CreateSpace(ctx context.Context, in *MsgCreateSpace, opts ...grpc.CallOption) (*MsgCreateSpaceResponse, error)
	// TransferSpace defines a method for transferring a space
	TransferSpace(ctx context.Context, in *MsgTransferSpace, opts ...grpc.CallOption) (*MsgTransferSpaceResponse, error)
	// CreateBlockHeader defines a method for creating a record
	CreateBlockHeader(ctx context.Context, in *MsgCreateBlockHeader, opts ...grpc.CallOption) (*MsgCreateBlockHeaderResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateSpace(ctx context.Context, in *MsgCreateSpace, opts ...grpc.CallOption) (*MsgCreateSpaceResponse, error) {
	out := new(MsgCreateSpaceResponse)
	err := c.cc.Invoke(ctx, Msg_CreateSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferSpace(ctx context.Context, in *MsgTransferSpace, opts ...grpc.CallOption) (*MsgTransferSpaceResponse, error) {
	out := new(MsgTransferSpaceResponse)
	err := c.cc.Invoke(ctx, Msg_TransferSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateBlockHeader(ctx context.Context, in *MsgCreateBlockHeader, opts ...grpc.CallOption) (*MsgCreateBlockHeaderResponse, error) {
	out := new(MsgCreateBlockHeaderResponse)
	err := c.cc.Invoke(ctx, Msg_CreateBlockHeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateSpace defines a method for creating a space
	CreateSpace(context.Context, *MsgCreateSpace) (*MsgCreateSpaceResponse, error)
	// TransferSpace defines a method for transferring a space
	TransferSpace(context.Context, *MsgTransferSpace) (*MsgTransferSpaceResponse, error)
	// CreateBlockHeader defines a method for creating a record
	CreateBlockHeader(context.Context, *MsgCreateBlockHeader) (*MsgCreateBlockHeaderResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateSpace(context.Context, *MsgCreateSpace) (*MsgCreateSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (UnimplementedMsgServer) TransferSpace(context.Context, *MsgTransferSpace) (*MsgTransferSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferSpace not implemented")
}
func (UnimplementedMsgServer) CreateBlockHeader(context.Context, *MsgCreateBlockHeader) (*MsgCreateBlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlockHeader not implemented")
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

func _Msg_CreateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateSpace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateSpace(ctx, req.(*MsgCreateSpace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferSpace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferSpace(ctx, req.(*MsgTransferSpace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBlockHeader)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateBlockHeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBlockHeader(ctx, req.(*MsgCreateBlockHeader))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iritamod.side_chain.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSpace",
			Handler:    _Msg_CreateSpace_Handler,
		},
		{
			MethodName: "TransferSpace",
			Handler:    _Msg_TransferSpace_Handler,
		},
		{
			MethodName: "CreateBlockHeader",
			Handler:    _Msg_CreateBlockHeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iritamod/side-chain/v1/tx.proto",
}