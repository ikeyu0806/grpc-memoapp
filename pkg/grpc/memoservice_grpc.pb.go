// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// MemoAPIClient is the client API for MemoAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemoAPIClient interface {
	GetMemo(ctx context.Context, in *GetMemoRequest, opts ...grpc.CallOption) (*GetMemoResponse, error)
	ListMemos(ctx context.Context, in *ListMemosRequest, opts ...grpc.CallOption) (*ListMemosResponse, error)
	CreateMemo(ctx context.Context, in *CreateMemoRequest, opts ...grpc.CallOption) (*CreateMemoResponse, error)
	UpdateMemo(ctx context.Context, in *UpdateMemoRequest, opts ...grpc.CallOption) (*UpdateMemoResponse, error)
	DeleteMemo(ctx context.Context, in *DeleteMemoRequest, opts ...grpc.CallOption) (*DeleteMemoResponse, error)
}

type memoAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoAPIClient(cc grpc.ClientConnInterface) MemoAPIClient {
	return &memoAPIClient{cc}
}

func (c *memoAPIClient) GetMemo(ctx context.Context, in *GetMemoRequest, opts ...grpc.CallOption) (*GetMemoResponse, error) {
	out := new(GetMemoResponse)
	err := c.cc.Invoke(ctx, "/service.MemoAPI/GetMemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoAPIClient) ListMemos(ctx context.Context, in *ListMemosRequest, opts ...grpc.CallOption) (*ListMemosResponse, error) {
	out := new(ListMemosResponse)
	err := c.cc.Invoke(ctx, "/service.MemoAPI/ListMemos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoAPIClient) CreateMemo(ctx context.Context, in *CreateMemoRequest, opts ...grpc.CallOption) (*CreateMemoResponse, error) {
	out := new(CreateMemoResponse)
	err := c.cc.Invoke(ctx, "/service.MemoAPI/CreateMemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoAPIClient) UpdateMemo(ctx context.Context, in *UpdateMemoRequest, opts ...grpc.CallOption) (*UpdateMemoResponse, error) {
	out := new(UpdateMemoResponse)
	err := c.cc.Invoke(ctx, "/service.MemoAPI/UpdateMemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoAPIClient) DeleteMemo(ctx context.Context, in *DeleteMemoRequest, opts ...grpc.CallOption) (*DeleteMemoResponse, error) {
	out := new(DeleteMemoResponse)
	err := c.cc.Invoke(ctx, "/service.MemoAPI/DeleteMemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemoAPIServer is the server API for MemoAPI service.
// All implementations must embed UnimplementedMemoAPIServer
// for forward compatibility
type MemoAPIServer interface {
	GetMemo(context.Context, *GetMemoRequest) (*GetMemoResponse, error)
	ListMemos(context.Context, *ListMemosRequest) (*ListMemosResponse, error)
	CreateMemo(context.Context, *CreateMemoRequest) (*CreateMemoResponse, error)
	UpdateMemo(context.Context, *UpdateMemoRequest) (*UpdateMemoResponse, error)
	DeleteMemo(context.Context, *DeleteMemoRequest) (*DeleteMemoResponse, error)
	mustEmbedUnimplementedMemoAPIServer()
}

// UnimplementedMemoAPIServer must be embedded to have forward compatible implementations.
type UnimplementedMemoAPIServer struct {
}

func (UnimplementedMemoAPIServer) GetMemo(context.Context, *GetMemoRequest) (*GetMemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemo not implemented")
}
func (UnimplementedMemoAPIServer) ListMemos(context.Context, *ListMemosRequest) (*ListMemosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMemos not implemented")
}
func (UnimplementedMemoAPIServer) CreateMemo(context.Context, *CreateMemoRequest) (*CreateMemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMemo not implemented")
}
func (UnimplementedMemoAPIServer) UpdateMemo(context.Context, *UpdateMemoRequest) (*UpdateMemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMemo not implemented")
}
func (UnimplementedMemoAPIServer) DeleteMemo(context.Context, *DeleteMemoRequest) (*DeleteMemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMemo not implemented")
}
func (UnimplementedMemoAPIServer) mustEmbedUnimplementedMemoAPIServer() {}

// UnsafeMemoAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemoAPIServer will
// result in compilation errors.
type UnsafeMemoAPIServer interface {
	mustEmbedUnimplementedMemoAPIServer()
}

func RegisterMemoAPIServer(s grpc.ServiceRegistrar, srv MemoAPIServer) {
	s.RegisterService(&MemoAPI_ServiceDesc, srv)
}

func _MemoAPI_GetMemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoAPIServer).GetMemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MemoAPI/GetMemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoAPIServer).GetMemo(ctx, req.(*GetMemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoAPI_ListMemos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMemosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoAPIServer).ListMemos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MemoAPI/ListMemos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoAPIServer).ListMemos(ctx, req.(*ListMemosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoAPI_CreateMemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoAPIServer).CreateMemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MemoAPI/CreateMemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoAPIServer).CreateMemo(ctx, req.(*CreateMemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoAPI_UpdateMemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoAPIServer).UpdateMemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MemoAPI/UpdateMemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoAPIServer).UpdateMemo(ctx, req.(*UpdateMemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoAPI_DeleteMemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoAPIServer).DeleteMemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MemoAPI/DeleteMemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoAPIServer).DeleteMemo(ctx, req.(*DeleteMemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemoAPI_ServiceDesc is the grpc.ServiceDesc for MemoAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemoAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MemoAPI",
	HandlerType: (*MemoAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMemo",
			Handler:    _MemoAPI_GetMemo_Handler,
		},
		{
			MethodName: "ListMemos",
			Handler:    _MemoAPI_ListMemos_Handler,
		},
		{
			MethodName: "CreateMemo",
			Handler:    _MemoAPI_CreateMemo_Handler,
		},
		{
			MethodName: "UpdateMemo",
			Handler:    _MemoAPI_UpdateMemo_Handler,
		},
		{
			MethodName: "DeleteMemo",
			Handler:    _MemoAPI_DeleteMemo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "memoservice.proto",
}
