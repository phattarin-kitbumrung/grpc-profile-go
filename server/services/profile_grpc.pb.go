// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: profile.proto

package services

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

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	Profile(ctx context.Context, opts ...grpc.CallOption) (Profile_ProfileClient, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) Profile(ctx context.Context, opts ...grpc.CallOption) (Profile_ProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Profile_ServiceDesc.Streams[0], "/services.Profile/Profile", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileProfileClient{stream}
	return x, nil
}

type Profile_ProfileClient interface {
	Send(*ProfileRequest) error
	Recv() (*ProfileResponse, error)
	grpc.ClientStream
}

type profileProfileClient struct {
	grpc.ClientStream
}

func (x *profileProfileClient) Send(m *ProfileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profileProfileClient) Recv() (*ProfileResponse, error) {
	m := new(ProfileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	Profile(Profile_ProfileServer) error
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) Profile(Profile_ProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_Profile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfileServer).Profile(&profileProfileServer{stream})
}

type Profile_ProfileServer interface {
	Send(*ProfileResponse) error
	Recv() (*ProfileRequest, error)
	grpc.ServerStream
}

type profileProfileServer struct {
	grpc.ServerStream
}

func (x *profileProfileServer) Send(m *ProfileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profileProfileServer) Recv() (*ProfileRequest, error) {
	m := new(ProfileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Profile",
			Handler:       _Profile_Profile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "profile.proto",
}