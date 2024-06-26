// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: dashboards.proto

package dpgen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DashboardsClient is the client API for Dashboards service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardsClient interface {
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetExchanges(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ExchangesResponse, error)
	GetMarkets(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*MarketsResponse, error)
	GetFunding(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error)
	GetOhlcv(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*OhlcvResponse, error)
	GetOpenInterest(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error)
	GetLiquidations(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error)
	GetOhlcvDiff(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*OhlcvDiffResponse, error)
	GetOpenInterestDiff(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error)
}

type dashboardsClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardsClient(cc grpc.ClientConnInterface) DashboardsClient {
	return &dashboardsClient{cc}
}

func (c *dashboardsClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/public.Dashboards/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsClient) GetExchanges(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ExchangesResponse, error) {
	out := new(ExchangesResponse)
	err := c.cc.Invoke(ctx, "/public.Dashboards/GetExchanges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsClient) GetMarkets(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*MarketsResponse, error) {
	out := new(MarketsResponse)
	err := c.cc.Invoke(ctx, "/public.Dashboards/GetMarkets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsClient) GetFunding(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/public.Dashboards/GetFunding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsClient) GetOhlcv(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*OhlcvResponse, error) {
	out := new(OhlcvResponse)
	err := c.cc.Invoke(ctx, "/public.Dashboards/GetOhlcv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsClient) GetOpenInterest(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/public.Dashboards/GetOpenInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsClient) GetLiquidations(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/public.Dashboards/GetLiquidations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsClient) GetOhlcvDiff(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*OhlcvDiffResponse, error) {
	out := new(OhlcvDiffResponse)
	err := c.cc.Invoke(ctx, "/public.Dashboards/GetOhlcvDiff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsClient) GetOpenInterestDiff(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/public.Dashboards/GetOpenInterestDiff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardsServer is the server API for Dashboards service.
// All implementations must embed UnimplementedDashboardsServer
// for forward compatibility
type DashboardsServer interface {
	Status(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetExchanges(context.Context, *emptypb.Empty) (*ExchangesResponse, error)
	GetMarkets(context.Context, *BasicRequest) (*MarketsResponse, error)
	GetFunding(context.Context, *BasicRequest) (*BasicResponse, error)
	GetOhlcv(context.Context, *BasicRequest) (*OhlcvResponse, error)
	GetOpenInterest(context.Context, *BasicRequest) (*BasicResponse, error)
	GetLiquidations(context.Context, *BasicRequest) (*BasicResponse, error)
	GetOhlcvDiff(context.Context, *BasicRequest) (*OhlcvDiffResponse, error)
	GetOpenInterestDiff(context.Context, *BasicRequest) (*BasicResponse, error)
	mustEmbedUnimplementedDashboardsServer()
}

// UnimplementedDashboardsServer must be embedded to have forward compatible implementations.
type UnimplementedDashboardsServer struct {
}

func (UnimplementedDashboardsServer) Status(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedDashboardsServer) GetExchanges(context.Context, *emptypb.Empty) (*ExchangesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchanges not implemented")
}
func (UnimplementedDashboardsServer) GetMarkets(context.Context, *BasicRequest) (*MarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarkets not implemented")
}
func (UnimplementedDashboardsServer) GetFunding(context.Context, *BasicRequest) (*BasicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFunding not implemented")
}
func (UnimplementedDashboardsServer) GetOhlcv(context.Context, *BasicRequest) (*OhlcvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOhlcv not implemented")
}
func (UnimplementedDashboardsServer) GetOpenInterest(context.Context, *BasicRequest) (*BasicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenInterest not implemented")
}
func (UnimplementedDashboardsServer) GetLiquidations(context.Context, *BasicRequest) (*BasicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiquidations not implemented")
}
func (UnimplementedDashboardsServer) GetOhlcvDiff(context.Context, *BasicRequest) (*OhlcvDiffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOhlcvDiff not implemented")
}
func (UnimplementedDashboardsServer) GetOpenInterestDiff(context.Context, *BasicRequest) (*BasicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenInterestDiff not implemented")
}
func (UnimplementedDashboardsServer) mustEmbedUnimplementedDashboardsServer() {}

// UnsafeDashboardsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardsServer will
// result in compilation errors.
type UnsafeDashboardsServer interface {
	mustEmbedUnimplementedDashboardsServer()
}

func RegisterDashboardsServer(s grpc.ServiceRegistrar, srv DashboardsServer) {
	s.RegisterService(&Dashboards_ServiceDesc, srv)
}

func _Dashboards_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboards_GetExchanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).GetExchanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/GetExchanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).GetExchanges(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboards_GetMarkets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).GetMarkets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/GetMarkets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).GetMarkets(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboards_GetFunding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).GetFunding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/GetFunding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).GetFunding(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboards_GetOhlcv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).GetOhlcv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/GetOhlcv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).GetOhlcv(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboards_GetOpenInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).GetOpenInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/GetOpenInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).GetOpenInterest(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboards_GetLiquidations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).GetLiquidations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/GetLiquidations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).GetLiquidations(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboards_GetOhlcvDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).GetOhlcvDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/GetOhlcvDiff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).GetOhlcvDiff(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboards_GetOpenInterestDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServer).GetOpenInterestDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.Dashboards/GetOpenInterestDiff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServer).GetOpenInterestDiff(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dashboards_ServiceDesc is the grpc.ServiceDesc for Dashboards service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dashboards_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.Dashboards",
	HandlerType: (*DashboardsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Dashboards_Status_Handler,
		},
		{
			MethodName: "GetExchanges",
			Handler:    _Dashboards_GetExchanges_Handler,
		},
		{
			MethodName: "GetMarkets",
			Handler:    _Dashboards_GetMarkets_Handler,
		},
		{
			MethodName: "GetFunding",
			Handler:    _Dashboards_GetFunding_Handler,
		},
		{
			MethodName: "GetOhlcv",
			Handler:    _Dashboards_GetOhlcv_Handler,
		},
		{
			MethodName: "GetOpenInterest",
			Handler:    _Dashboards_GetOpenInterest_Handler,
		},
		{
			MethodName: "GetLiquidations",
			Handler:    _Dashboards_GetLiquidations_Handler,
		},
		{
			MethodName: "GetOhlcvDiff",
			Handler:    _Dashboards_GetOhlcvDiff_Handler,
		},
		{
			MethodName: "GetOpenInterestDiff",
			Handler:    _Dashboards_GetOpenInterestDiff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dashboards.proto",
}
