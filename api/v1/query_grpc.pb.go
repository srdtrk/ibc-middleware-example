// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmosregistry/example/v1/query.proto

package examplev1

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
	Query_CallbackCounter_FullMethodName  = "/cosmosregistry.example.v1.Query/CallbackCounter"
	Query_CallbackCounters_FullMethodName = "/cosmosregistry.example.v1.Query/CallbackCounters"
	Query_Params_FullMethodName           = "/cosmosregistry.example.v1.Query/Params"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// CallbackCounter returns the current callback counter value.
	CallbackCounter(ctx context.Context, in *QueryCallbackCounterRequest, opts ...grpc.CallOption) (*QueryCallbackCounterResponse, error)
	// CallbackCounters returns all the counter values.
	CallbackCounters(ctx context.Context, in *QueryCallbackCountersRequest, opts ...grpc.CallOption) (*QueryCallbackCountersResponse, error)
	// Params returns the module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CallbackCounter(ctx context.Context, in *QueryCallbackCounterRequest, opts ...grpc.CallOption) (*QueryCallbackCounterResponse, error) {
	out := new(QueryCallbackCounterResponse)
	err := c.cc.Invoke(ctx, Query_CallbackCounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CallbackCounters(ctx context.Context, in *QueryCallbackCountersRequest, opts ...grpc.CallOption) (*QueryCallbackCountersResponse, error) {
	out := new(QueryCallbackCountersResponse)
	err := c.cc.Invoke(ctx, Query_CallbackCounters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// CallbackCounter returns the current callback counter value.
	CallbackCounter(context.Context, *QueryCallbackCounterRequest) (*QueryCallbackCounterResponse, error)
	// CallbackCounters returns all the counter values.
	CallbackCounters(context.Context, *QueryCallbackCountersRequest) (*QueryCallbackCountersResponse, error)
	// Params returns the module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) CallbackCounter(context.Context, *QueryCallbackCounterRequest) (*QueryCallbackCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackCounter not implemented")
}
func (UnimplementedQueryServer) CallbackCounters(context.Context, *QueryCallbackCountersRequest) (*QueryCallbackCountersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackCounters not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_CallbackCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCallbackCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CallbackCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CallbackCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CallbackCounter(ctx, req.(*QueryCallbackCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CallbackCounters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCallbackCountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CallbackCounters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CallbackCounters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CallbackCounters(ctx, req.(*QueryCallbackCountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmosregistry.example.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallbackCounter",
			Handler:    _Query_CallbackCounter_Handler,
		},
		{
			MethodName: "CallbackCounters",
			Handler:    _Query_CallbackCounters_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmosregistry/example/v1/query.proto",
}
