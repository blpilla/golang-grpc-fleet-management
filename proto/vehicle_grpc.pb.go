// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: vehicle.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VehicleService_CreateVehicle_FullMethodName  = "/proto.VehicleService/CreateVehicle"
	VehicleService_GetVehicleByID_FullMethodName = "/proto.VehicleService/GetVehicleByID"
	VehicleService_GetAllVehicles_FullMethodName = "/proto.VehicleService/GetAllVehicles"
	VehicleService_UpdateVehicle_FullMethodName  = "/proto.VehicleService/UpdateVehicle"
	VehicleService_DeleteVehicle_FullMethodName  = "/proto.VehicleService/DeleteVehicle"
)

// VehicleServiceClient is the client API for VehicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Serviço de Vehicle
type VehicleServiceClient interface {
	CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...grpc.CallOption) (*VehicleResponse, error)
	GetVehicleByID(ctx context.Context, in *GetVehicleByIDRequest, opts ...grpc.CallOption) (*VehicleResponse, error)
	GetAllVehicles(ctx context.Context, in *GetAllVehiclesRequest, opts ...grpc.CallOption) (*GetAllVehiclesResponse, error)
	UpdateVehicle(ctx context.Context, in *UpdateVehicleRequest, opts ...grpc.CallOption) (*VehicleResponse, error)
	DeleteVehicle(ctx context.Context, in *DeleteVehicleRequest, opts ...grpc.CallOption) (*DeleteVehicleResponse, error)
}

type vehicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleServiceClient(cc grpc.ClientConnInterface) VehicleServiceClient {
	return &vehicleServiceClient{cc}
}

func (c *vehicleServiceClient) CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...grpc.CallOption) (*VehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VehicleResponse)
	err := c.cc.Invoke(ctx, VehicleService_CreateVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) GetVehicleByID(ctx context.Context, in *GetVehicleByIDRequest, opts ...grpc.CallOption) (*VehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VehicleResponse)
	err := c.cc.Invoke(ctx, VehicleService_GetVehicleByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) GetAllVehicles(ctx context.Context, in *GetAllVehiclesRequest, opts ...grpc.CallOption) (*GetAllVehiclesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllVehiclesResponse)
	err := c.cc.Invoke(ctx, VehicleService_GetAllVehicles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) UpdateVehicle(ctx context.Context, in *UpdateVehicleRequest, opts ...grpc.CallOption) (*VehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VehicleResponse)
	err := c.cc.Invoke(ctx, VehicleService_UpdateVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) DeleteVehicle(ctx context.Context, in *DeleteVehicleRequest, opts ...grpc.CallOption) (*DeleteVehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVehicleResponse)
	err := c.cc.Invoke(ctx, VehicleService_DeleteVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServiceServer is the server API for VehicleService service.
// All implementations must embed UnimplementedVehicleServiceServer
// for forward compatibility.
//
// Serviço de Vehicle
type VehicleServiceServer interface {
	CreateVehicle(context.Context, *CreateVehicleRequest) (*VehicleResponse, error)
	GetVehicleByID(context.Context, *GetVehicleByIDRequest) (*VehicleResponse, error)
	GetAllVehicles(context.Context, *GetAllVehiclesRequest) (*GetAllVehiclesResponse, error)
	UpdateVehicle(context.Context, *UpdateVehicleRequest) (*VehicleResponse, error)
	DeleteVehicle(context.Context, *DeleteVehicleRequest) (*DeleteVehicleResponse, error)
	mustEmbedUnimplementedVehicleServiceServer()
}

// UnimplementedVehicleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVehicleServiceServer struct{}

func (UnimplementedVehicleServiceServer) CreateVehicle(context.Context, *CreateVehicleRequest) (*VehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) GetVehicleByID(context.Context, *GetVehicleByIDRequest) (*VehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicleByID not implemented")
}
func (UnimplementedVehicleServiceServer) GetAllVehicles(context.Context, *GetAllVehiclesRequest) (*GetAllVehiclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVehicles not implemented")
}
func (UnimplementedVehicleServiceServer) UpdateVehicle(context.Context, *UpdateVehicleRequest) (*VehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) DeleteVehicle(context.Context, *DeleteVehicleRequest) (*DeleteVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) mustEmbedUnimplementedVehicleServiceServer() {}
func (UnimplementedVehicleServiceServer) testEmbeddedByValue()                        {}

// UnsafeVehicleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleServiceServer will
// result in compilation errors.
type UnsafeVehicleServiceServer interface {
	mustEmbedUnimplementedVehicleServiceServer()
}

func RegisterVehicleServiceServer(s grpc.ServiceRegistrar, srv VehicleServiceServer) {
	// If the following call pancis, it indicates UnimplementedVehicleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VehicleService_ServiceDesc, srv)
}

func _VehicleService_CreateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).CreateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_CreateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).CreateVehicle(ctx, req.(*CreateVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_GetVehicleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).GetVehicleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_GetVehicleByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).GetVehicleByID(ctx, req.(*GetVehicleByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_GetAllVehicles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllVehiclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).GetAllVehicles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_GetAllVehicles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).GetAllVehicles(ctx, req.(*GetAllVehiclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_UpdateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).UpdateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_UpdateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).UpdateVehicle(ctx, req.(*UpdateVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_DeleteVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).DeleteVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_DeleteVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).DeleteVehicle(ctx, req.(*DeleteVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VehicleService_ServiceDesc is the grpc.ServiceDesc for VehicleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VehicleService",
	HandlerType: (*VehicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVehicle",
			Handler:    _VehicleService_CreateVehicle_Handler,
		},
		{
			MethodName: "GetVehicleByID",
			Handler:    _VehicleService_GetVehicleByID_Handler,
		},
		{
			MethodName: "GetAllVehicles",
			Handler:    _VehicleService_GetAllVehicles_Handler,
		},
		{
			MethodName: "UpdateVehicle",
			Handler:    _VehicleService_UpdateVehicle_Handler,
		},
		{
			MethodName: "DeleteVehicle",
			Handler:    _VehicleService_DeleteVehicle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vehicle.proto",
}
