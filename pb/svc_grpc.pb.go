// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: svc.proto

package pb

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Call(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	Ctl(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Service_CtlClient, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Call(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/svc.Service/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ctl(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Service_CtlClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], "/svc.Service/Ctl", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceCtlClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_CtlClient interface {
	Recv() (*Cmd, error)
	grpc.ClientStream
}

type serviceCtlClient struct {
	grpc.ClientStream
}

func (x *serviceCtlClient) Recv() (*Cmd, error) {
	m := new(Cmd)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Call(context.Context, *Req) (*Resp, error)
	Ctl(*emptypb.Empty, Service_CtlServer) error
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Call(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedServiceServer) Ctl(*emptypb.Empty, Service_CtlServer) error {
	return status.Errorf(codes.Unimplemented, "method Ctl not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc.Service/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Call(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Ctl_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Ctl(m, &serviceCtlServer{stream})
}

type Service_CtlServer interface {
	Send(*Cmd) error
	grpc.ServerStream
}

type serviceCtlServer struct {
	grpc.ServerStream
}

func (x *serviceCtlServer) Send(m *Cmd) error {
	return x.ServerStream.SendMsg(m)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Service_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ctl",
			Handler:       _Service_Ctl_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "svc.proto",
}
