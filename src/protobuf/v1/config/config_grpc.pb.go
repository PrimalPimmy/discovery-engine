// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: v1/config/config.proto

package config

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

// ConfigStoreClient is the client API for ConfigStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigStoreClient interface {
	Add(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	Get(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	Update(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	Delete(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	Apply(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type configStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigStoreClient(cc grpc.ClientConnInterface) ConfigStoreClient {
	return &configStoreClient{cc}
}

func (c *configStoreClient) Add(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/v1.config.ConfigStore/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configStoreClient) Get(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/v1.config.ConfigStore/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configStoreClient) Update(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/v1.config.ConfigStore/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configStoreClient) Delete(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/v1.config.ConfigStore/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configStoreClient) Apply(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/v1.config.ConfigStore/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigStoreServer is the server API for ConfigStore service.
// All implementations must embed UnimplementedConfigStoreServer
// for forward compatibility
type ConfigStoreServer interface {
	Add(context.Context, *ConfigRequest) (*ConfigResponse, error)
	Get(context.Context, *ConfigRequest) (*ConfigResponse, error)
	Update(context.Context, *ConfigRequest) (*ConfigResponse, error)
	Delete(context.Context, *ConfigRequest) (*ConfigResponse, error)
	Apply(context.Context, *ConfigRequest) (*ConfigResponse, error)
	mustEmbedUnimplementedConfigStoreServer()
}

// UnimplementedConfigStoreServer must be embedded to have forward compatible implementations.
type UnimplementedConfigStoreServer struct {
}

func (UnimplementedConfigStoreServer) Add(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedConfigStoreServer) Get(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedConfigStoreServer) Update(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedConfigStoreServer) Delete(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedConfigStoreServer) Apply(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedConfigStoreServer) mustEmbedUnimplementedConfigStoreServer() {}

// UnsafeConfigStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigStoreServer will
// result in compilation errors.
type UnsafeConfigStoreServer interface {
	mustEmbedUnimplementedConfigStoreServer()
}

func RegisterConfigStoreServer(s grpc.ServiceRegistrar, srv ConfigStoreServer) {
	s.RegisterService(&ConfigStore_ServiceDesc, srv)
}

func _ConfigStore_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.config.ConfigStore/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).Add(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.config.ConfigStore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).Get(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigStore_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.config.ConfigStore/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).Update(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.config.ConfigStore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).Delete(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigStore_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.config.ConfigStore/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).Apply(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigStore_ServiceDesc is the grpc.ServiceDesc for ConfigStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.config.ConfigStore",
	HandlerType: (*ConfigStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ConfigStore_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ConfigStore_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ConfigStore_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ConfigStore_Delete_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _ConfigStore_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/config/config.proto",
}
