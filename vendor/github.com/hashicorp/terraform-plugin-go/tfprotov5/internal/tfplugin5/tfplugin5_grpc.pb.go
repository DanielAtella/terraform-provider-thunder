// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tfplugin5

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	//////// Information about what a provider supports/expects
	GetSchema(ctx context.Context, in *GetProviderSchema_Request, opts ...grpc.CallOption) (*GetProviderSchema_Response, error)
	PrepareProviderConfig(ctx context.Context, in *PrepareProviderConfig_Request, opts ...grpc.CallOption) (*PrepareProviderConfig_Response, error)
	ValidateResourceTypeConfig(ctx context.Context, in *ValidateResourceTypeConfig_Request, opts ...grpc.CallOption) (*ValidateResourceTypeConfig_Response, error)
	ValidateDataSourceConfig(ctx context.Context, in *ValidateDataSourceConfig_Request, opts ...grpc.CallOption) (*ValidateDataSourceConfig_Response, error)
	UpgradeResourceState(ctx context.Context, in *UpgradeResourceState_Request, opts ...grpc.CallOption) (*UpgradeResourceState_Response, error)
	//////// One-time initialization, called before other functions below
	Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error)
	//////// Managed Resource Lifecycle
	ReadResource(ctx context.Context, in *ReadResource_Request, opts ...grpc.CallOption) (*ReadResource_Response, error)
	PlanResourceChange(ctx context.Context, in *PlanResourceChange_Request, opts ...grpc.CallOption) (*PlanResourceChange_Response, error)
	ApplyResourceChange(ctx context.Context, in *ApplyResourceChange_Request, opts ...grpc.CallOption) (*ApplyResourceChange_Response, error)
	ImportResourceState(ctx context.Context, in *ImportResourceState_Request, opts ...grpc.CallOption) (*ImportResourceState_Response, error)
	ReadDataSource(ctx context.Context, in *ReadDataSource_Request, opts ...grpc.CallOption) (*ReadDataSource_Response, error)
	//////// Graceful Shutdown
	Stop(ctx context.Context, in *Stop_Request, opts ...grpc.CallOption) (*Stop_Response, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) GetSchema(ctx context.Context, in *GetProviderSchema_Request, opts ...grpc.CallOption) (*GetProviderSchema_Response, error) {
	out := new(GetProviderSchema_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) PrepareProviderConfig(ctx context.Context, in *PrepareProviderConfig_Request, opts ...grpc.CallOption) (*PrepareProviderConfig_Response, error) {
	out := new(PrepareProviderConfig_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/PrepareProviderConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ValidateResourceTypeConfig(ctx context.Context, in *ValidateResourceTypeConfig_Request, opts ...grpc.CallOption) (*ValidateResourceTypeConfig_Response, error) {
	out := new(ValidateResourceTypeConfig_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/ValidateResourceTypeConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ValidateDataSourceConfig(ctx context.Context, in *ValidateDataSourceConfig_Request, opts ...grpc.CallOption) (*ValidateDataSourceConfig_Response, error) {
	out := new(ValidateDataSourceConfig_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/ValidateDataSourceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) UpgradeResourceState(ctx context.Context, in *UpgradeResourceState_Request, opts ...grpc.CallOption) (*UpgradeResourceState_Response, error) {
	out := new(UpgradeResourceState_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/UpgradeResourceState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error) {
	out := new(Configure_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ReadResource(ctx context.Context, in *ReadResource_Request, opts ...grpc.CallOption) (*ReadResource_Response, error) {
	out := new(ReadResource_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/ReadResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) PlanResourceChange(ctx context.Context, in *PlanResourceChange_Request, opts ...grpc.CallOption) (*PlanResourceChange_Response, error) {
	out := new(PlanResourceChange_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/PlanResourceChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ApplyResourceChange(ctx context.Context, in *ApplyResourceChange_Request, opts ...grpc.CallOption) (*ApplyResourceChange_Response, error) {
	out := new(ApplyResourceChange_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/ApplyResourceChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ImportResourceState(ctx context.Context, in *ImportResourceState_Request, opts ...grpc.CallOption) (*ImportResourceState_Response, error) {
	out := new(ImportResourceState_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/ImportResourceState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ReadDataSource(ctx context.Context, in *ReadDataSource_Request, opts ...grpc.CallOption) (*ReadDataSource_Response, error) {
	out := new(ReadDataSource_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/ReadDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Stop(ctx context.Context, in *Stop_Request, opts ...grpc.CallOption) (*Stop_Response, error) {
	out := new(Stop_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provider/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	//////// Information about what a provider supports/expects
	GetSchema(context.Context, *GetProviderSchema_Request) (*GetProviderSchema_Response, error)
	PrepareProviderConfig(context.Context, *PrepareProviderConfig_Request) (*PrepareProviderConfig_Response, error)
	ValidateResourceTypeConfig(context.Context, *ValidateResourceTypeConfig_Request) (*ValidateResourceTypeConfig_Response, error)
	ValidateDataSourceConfig(context.Context, *ValidateDataSourceConfig_Request) (*ValidateDataSourceConfig_Response, error)
	UpgradeResourceState(context.Context, *UpgradeResourceState_Request) (*UpgradeResourceState_Response, error)
	//////// One-time initialization, called before other functions below
	Configure(context.Context, *Configure_Request) (*Configure_Response, error)
	//////// Managed Resource Lifecycle
	ReadResource(context.Context, *ReadResource_Request) (*ReadResource_Response, error)
	PlanResourceChange(context.Context, *PlanResourceChange_Request) (*PlanResourceChange_Response, error)
	ApplyResourceChange(context.Context, *ApplyResourceChange_Request) (*ApplyResourceChange_Response, error)
	ImportResourceState(context.Context, *ImportResourceState_Request) (*ImportResourceState_Response, error)
	ReadDataSource(context.Context, *ReadDataSource_Request) (*ReadDataSource_Response, error)
	//////// Graceful Shutdown
	Stop(context.Context, *Stop_Request) (*Stop_Response, error)
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) GetSchema(context.Context, *GetProviderSchema_Request) (*GetProviderSchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedProviderServer) PrepareProviderConfig(context.Context, *PrepareProviderConfig_Request) (*PrepareProviderConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareProviderConfig not implemented")
}
func (UnimplementedProviderServer) ValidateResourceTypeConfig(context.Context, *ValidateResourceTypeConfig_Request) (*ValidateResourceTypeConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateResourceTypeConfig not implemented")
}
func (UnimplementedProviderServer) ValidateDataSourceConfig(context.Context, *ValidateDataSourceConfig_Request) (*ValidateDataSourceConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateDataSourceConfig not implemented")
}
func (UnimplementedProviderServer) UpgradeResourceState(context.Context, *UpgradeResourceState_Request) (*UpgradeResourceState_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeResourceState not implemented")
}
func (UnimplementedProviderServer) Configure(context.Context, *Configure_Request) (*Configure_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedProviderServer) ReadResource(context.Context, *ReadResource_Request) (*ReadResource_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadResource not implemented")
}
func (UnimplementedProviderServer) PlanResourceChange(context.Context, *PlanResourceChange_Request) (*PlanResourceChange_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlanResourceChange not implemented")
}
func (UnimplementedProviderServer) ApplyResourceChange(context.Context, *ApplyResourceChange_Request) (*ApplyResourceChange_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyResourceChange not implemented")
}
func (UnimplementedProviderServer) ImportResourceState(context.Context, *ImportResourceState_Request) (*ImportResourceState_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportResourceState not implemented")
}
func (UnimplementedProviderServer) ReadDataSource(context.Context, *ReadDataSource_Request) (*ReadDataSource_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadDataSource not implemented")
}
func (UnimplementedProviderServer) Stop(context.Context, *Stop_Request) (*Stop_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s *grpc.Server, srv ProviderServer) {
	s.RegisterService(&_Provider_serviceDesc, srv)
}

func _Provider_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProviderSchema_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).GetSchema(ctx, req.(*GetProviderSchema_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_PrepareProviderConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareProviderConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).PrepareProviderConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/PrepareProviderConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).PrepareProviderConfig(ctx, req.(*PrepareProviderConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ValidateResourceTypeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateResourceTypeConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ValidateResourceTypeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/ValidateResourceTypeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ValidateResourceTypeConfig(ctx, req.(*ValidateResourceTypeConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ValidateDataSourceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateDataSourceConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ValidateDataSourceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/ValidateDataSourceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ValidateDataSourceConfig(ctx, req.(*ValidateDataSourceConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_UpgradeResourceState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeResourceState_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).UpgradeResourceState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/UpgradeResourceState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).UpgradeResourceState(ctx, req.(*UpgradeResourceState_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configure_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Configure(ctx, req.(*Configure_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ReadResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadResource_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ReadResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/ReadResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ReadResource(ctx, req.(*ReadResource_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_PlanResourceChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanResourceChange_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).PlanResourceChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/PlanResourceChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).PlanResourceChange(ctx, req.(*PlanResourceChange_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ApplyResourceChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyResourceChange_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ApplyResourceChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/ApplyResourceChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ApplyResourceChange(ctx, req.(*ApplyResourceChange_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ImportResourceState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportResourceState_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ImportResourceState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/ImportResourceState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ImportResourceState(ctx, req.(*ImportResourceState_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ReadDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadDataSource_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ReadDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/ReadDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ReadDataSource(ctx, req.(*ReadDataSource_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stop_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provider/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Stop(ctx, req.(*Stop_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Provider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tfplugin5.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchema",
			Handler:    _Provider_GetSchema_Handler,
		},
		{
			MethodName: "PrepareProviderConfig",
			Handler:    _Provider_PrepareProviderConfig_Handler,
		},
		{
			MethodName: "ValidateResourceTypeConfig",
			Handler:    _Provider_ValidateResourceTypeConfig_Handler,
		},
		{
			MethodName: "ValidateDataSourceConfig",
			Handler:    _Provider_ValidateDataSourceConfig_Handler,
		},
		{
			MethodName: "UpgradeResourceState",
			Handler:    _Provider_UpgradeResourceState_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Provider_Configure_Handler,
		},
		{
			MethodName: "ReadResource",
			Handler:    _Provider_ReadResource_Handler,
		},
		{
			MethodName: "PlanResourceChange",
			Handler:    _Provider_PlanResourceChange_Handler,
		},
		{
			MethodName: "ApplyResourceChange",
			Handler:    _Provider_ApplyResourceChange_Handler,
		},
		{
			MethodName: "ImportResourceState",
			Handler:    _Provider_ImportResourceState_Handler,
		},
		{
			MethodName: "ReadDataSource",
			Handler:    _Provider_ReadDataSource_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Provider_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tfplugin5.proto",
}

// ProvisionerClient is the client API for Provisioner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvisionerClient interface {
	GetSchema(ctx context.Context, in *GetProvisionerSchema_Request, opts ...grpc.CallOption) (*GetProvisionerSchema_Response, error)
	ValidateProvisionerConfig(ctx context.Context, in *ValidateProvisionerConfig_Request, opts ...grpc.CallOption) (*ValidateProvisionerConfig_Response, error)
	ProvisionResource(ctx context.Context, in *ProvisionResource_Request, opts ...grpc.CallOption) (Provisioner_ProvisionResourceClient, error)
	Stop(ctx context.Context, in *Stop_Request, opts ...grpc.CallOption) (*Stop_Response, error)
}

type provisionerClient struct {
	cc grpc.ClientConnInterface
}

func NewProvisionerClient(cc grpc.ClientConnInterface) ProvisionerClient {
	return &provisionerClient{cc}
}

func (c *provisionerClient) GetSchema(ctx context.Context, in *GetProvisionerSchema_Request, opts ...grpc.CallOption) (*GetProvisionerSchema_Response, error) {
	out := new(GetProvisionerSchema_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provisioner/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionerClient) ValidateProvisionerConfig(ctx context.Context, in *ValidateProvisionerConfig_Request, opts ...grpc.CallOption) (*ValidateProvisionerConfig_Response, error) {
	out := new(ValidateProvisionerConfig_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provisioner/ValidateProvisionerConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionerClient) ProvisionResource(ctx context.Context, in *ProvisionResource_Request, opts ...grpc.CallOption) (Provisioner_ProvisionResourceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Provisioner_serviceDesc.Streams[0], "/tfplugin5.Provisioner/ProvisionResource", opts...)
	if err != nil {
		return nil, err
	}
	x := &provisionerProvisionResourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Provisioner_ProvisionResourceClient interface {
	Recv() (*ProvisionResource_Response, error)
	grpc.ClientStream
}

type provisionerProvisionResourceClient struct {
	grpc.ClientStream
}

func (x *provisionerProvisionResourceClient) Recv() (*ProvisionResource_Response, error) {
	m := new(ProvisionResource_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *provisionerClient) Stop(ctx context.Context, in *Stop_Request, opts ...grpc.CallOption) (*Stop_Response, error) {
	out := new(Stop_Response)
	err := c.cc.Invoke(ctx, "/tfplugin5.Provisioner/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionerServer is the server API for Provisioner service.
// All implementations must embed UnimplementedProvisionerServer
// for forward compatibility
type ProvisionerServer interface {
	GetSchema(context.Context, *GetProvisionerSchema_Request) (*GetProvisionerSchema_Response, error)
	ValidateProvisionerConfig(context.Context, *ValidateProvisionerConfig_Request) (*ValidateProvisionerConfig_Response, error)
	ProvisionResource(*ProvisionResource_Request, Provisioner_ProvisionResourceServer) error
	Stop(context.Context, *Stop_Request) (*Stop_Response, error)
	mustEmbedUnimplementedProvisionerServer()
}

// UnimplementedProvisionerServer must be embedded to have forward compatible implementations.
type UnimplementedProvisionerServer struct {
}

func (UnimplementedProvisionerServer) GetSchema(context.Context, *GetProvisionerSchema_Request) (*GetProvisionerSchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedProvisionerServer) ValidateProvisionerConfig(context.Context, *ValidateProvisionerConfig_Request) (*ValidateProvisionerConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateProvisionerConfig not implemented")
}
func (UnimplementedProvisionerServer) ProvisionResource(*ProvisionResource_Request, Provisioner_ProvisionResourceServer) error {
	return status.Errorf(codes.Unimplemented, "method ProvisionResource not implemented")
}
func (UnimplementedProvisionerServer) Stop(context.Context, *Stop_Request) (*Stop_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedProvisionerServer) mustEmbedUnimplementedProvisionerServer() {}

// UnsafeProvisionerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvisionerServer will
// result in compilation errors.
type UnsafeProvisionerServer interface {
	mustEmbedUnimplementedProvisionerServer()
}

func RegisterProvisionerServer(s *grpc.Server, srv ProvisionerServer) {
	s.RegisterService(&_Provisioner_serviceDesc, srv)
}

func _Provisioner_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProvisionerSchema_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provisioner/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).GetSchema(ctx, req.(*GetProvisionerSchema_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioner_ValidateProvisionerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateProvisionerConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).ValidateProvisionerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provisioner/ValidateProvisionerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).ValidateProvisionerConfig(ctx, req.(*ValidateProvisionerConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioner_ProvisionResource_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProvisionResource_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProvisionerServer).ProvisionResource(m, &provisionerProvisionResourceServer{stream})
}

type Provisioner_ProvisionResourceServer interface {
	Send(*ProvisionResource_Response) error
	grpc.ServerStream
}

type provisionerProvisionResourceServer struct {
	grpc.ServerStream
}

func (x *provisionerProvisionResourceServer) Send(m *ProvisionResource_Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Provisioner_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stop_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfplugin5.Provisioner/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).Stop(ctx, req.(*Stop_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Provisioner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tfplugin5.Provisioner",
	HandlerType: (*ProvisionerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchema",
			Handler:    _Provisioner_GetSchema_Handler,
		},
		{
			MethodName: "ValidateProvisionerConfig",
			Handler:    _Provisioner_ValidateProvisionerConfig_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Provisioner_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProvisionResource",
			Handler:       _Provisioner_ProvisionResource_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tfplugin5.proto",
}
