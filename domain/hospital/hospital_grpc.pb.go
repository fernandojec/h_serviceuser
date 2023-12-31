// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: domain/hospital/proto/hospital.proto

package hospital

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

// HospitalServiceClient is the client API for HospitalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HospitalServiceClient interface {
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HospitalListProto, error)
	Get(ctx context.Context, in *HospitalGetProto, opts ...grpc.CallOption) (*HospitalProto, error)
	Add(ctx context.Context, in *HospitalAddProto, opts ...grpc.CallOption) (*HospitalProto, error)
}

type hospitalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalServiceClient(cc grpc.ClientConnInterface) HospitalServiceClient {
	return &hospitalServiceClient{cc}
}

func (c *hospitalServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HospitalListProto, error) {
	out := new(HospitalListProto)
	err := c.cc.Invoke(ctx, "/hospital.HospitalService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalServiceClient) Get(ctx context.Context, in *HospitalGetProto, opts ...grpc.CallOption) (*HospitalProto, error) {
	out := new(HospitalProto)
	err := c.cc.Invoke(ctx, "/hospital.HospitalService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalServiceClient) Add(ctx context.Context, in *HospitalAddProto, opts ...grpc.CallOption) (*HospitalProto, error) {
	out := new(HospitalProto)
	err := c.cc.Invoke(ctx, "/hospital.HospitalService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServiceServer is the server API for HospitalService service.
// All implementations must embed UnimplementedHospitalServiceServer
// for forward compatibility
type HospitalServiceServer interface {
	List(context.Context, *emptypb.Empty) (*HospitalListProto, error)
	Get(context.Context, *HospitalGetProto) (*HospitalProto, error)
	Add(context.Context, *HospitalAddProto) (*HospitalProto, error)
	mustEmbedUnimplementedHospitalServiceServer()
}

// UnimplementedHospitalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHospitalServiceServer struct {
}

func (UnimplementedHospitalServiceServer) List(context.Context, *emptypb.Empty) (*HospitalListProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedHospitalServiceServer) Get(context.Context, *HospitalGetProto) (*HospitalProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHospitalServiceServer) Add(context.Context, *HospitalAddProto) (*HospitalProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedHospitalServiceServer) mustEmbedUnimplementedHospitalServiceServer() {}

// UnsafeHospitalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HospitalServiceServer will
// result in compilation errors.
type UnsafeHospitalServiceServer interface {
	mustEmbedUnimplementedHospitalServiceServer()
}

func RegisterHospitalServiceServer(s grpc.ServiceRegistrar, srv HospitalServiceServer) {
	s.RegisterService(&HospitalService_ServiceDesc, srv)
}

func _HospitalService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.HospitalService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HospitalService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HospitalGetProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.HospitalService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).Get(ctx, req.(*HospitalGetProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _HospitalService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HospitalAddProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.HospitalService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).Add(ctx, req.(*HospitalAddProto))
	}
	return interceptor(ctx, in, info, handler)
}

// HospitalService_ServiceDesc is the grpc.ServiceDesc for HospitalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HospitalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hospital.HospitalService",
	HandlerType: (*HospitalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _HospitalService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _HospitalService_Get_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _HospitalService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain/hospital/proto/hospital.proto",
}
