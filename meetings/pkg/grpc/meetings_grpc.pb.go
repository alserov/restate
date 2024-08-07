// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: pkg/grpc/meetings.proto

package meetings

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

// MeetingsServiceClient is the client API for MeetingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeetingsServiceClient interface {
	ArrangeMeeting(ctx context.Context, in *Meeting, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelMeeting(ctx context.Context, in *CancelMeetingParameter, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAvailableTimeForMeeting(ctx context.Context, in *GetAvailableTimeForMeetingParameter, opts ...grpc.CallOption) (*AvailableTimeList, error)
	GetMeetingsByEstateID(ctx context.Context, in *GetMeetingsByEstateIDParameter, opts ...grpc.CallOption) (*Meetings, error)
	GetMeetingsByPhoneNumber(ctx context.Context, in *GetMeetingsByPhoneNumberParameter, opts ...grpc.CallOption) (*Meetings, error)
}

type meetingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeetingsServiceClient(cc grpc.ClientConnInterface) MeetingsServiceClient {
	return &meetingsServiceClient{cc}
}

func (c *meetingsServiceClient) ArrangeMeeting(ctx context.Context, in *Meeting, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/meetings.MeetingsService/ArrangeMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsServiceClient) CancelMeeting(ctx context.Context, in *CancelMeetingParameter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/meetings.MeetingsService/CancelMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsServiceClient) GetAvailableTimeForMeeting(ctx context.Context, in *GetAvailableTimeForMeetingParameter, opts ...grpc.CallOption) (*AvailableTimeList, error) {
	out := new(AvailableTimeList)
	err := c.cc.Invoke(ctx, "/meetings.MeetingsService/GetAvailableTimeForMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsServiceClient) GetMeetingsByEstateID(ctx context.Context, in *GetMeetingsByEstateIDParameter, opts ...grpc.CallOption) (*Meetings, error) {
	out := new(Meetings)
	err := c.cc.Invoke(ctx, "/meetings.MeetingsService/GetMeetingsByEstateID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsServiceClient) GetMeetingsByPhoneNumber(ctx context.Context, in *GetMeetingsByPhoneNumberParameter, opts ...grpc.CallOption) (*Meetings, error) {
	out := new(Meetings)
	err := c.cc.Invoke(ctx, "/meetings.MeetingsService/GetMeetingsByPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeetingsServiceServer is the server API for MeetingsService service.
// All implementations must embed UnimplementedMeetingsServiceServer
// for forward compatibility
type MeetingsServiceServer interface {
	ArrangeMeeting(context.Context, *Meeting) (*emptypb.Empty, error)
	CancelMeeting(context.Context, *CancelMeetingParameter) (*emptypb.Empty, error)
	GetAvailableTimeForMeeting(context.Context, *GetAvailableTimeForMeetingParameter) (*AvailableTimeList, error)
	GetMeetingsByEstateID(context.Context, *GetMeetingsByEstateIDParameter) (*Meetings, error)
	GetMeetingsByPhoneNumber(context.Context, *GetMeetingsByPhoneNumberParameter) (*Meetings, error)
	mustEmbedUnimplementedMeetingsServiceServer()
}

// UnimplementedMeetingsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeetingsServiceServer struct {
}

func (UnimplementedMeetingsServiceServer) ArrangeMeeting(context.Context, *Meeting) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArrangeMeeting not implemented")
}
func (UnimplementedMeetingsServiceServer) CancelMeeting(context.Context, *CancelMeetingParameter) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelMeeting not implemented")
}
func (UnimplementedMeetingsServiceServer) GetAvailableTimeForMeeting(context.Context, *GetAvailableTimeForMeetingParameter) (*AvailableTimeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableTimeForMeeting not implemented")
}
func (UnimplementedMeetingsServiceServer) GetMeetingsByEstateID(context.Context, *GetMeetingsByEstateIDParameter) (*Meetings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingsByEstateID not implemented")
}
func (UnimplementedMeetingsServiceServer) GetMeetingsByPhoneNumber(context.Context, *GetMeetingsByPhoneNumberParameter) (*Meetings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingsByPhoneNumber not implemented")
}
func (UnimplementedMeetingsServiceServer) mustEmbedUnimplementedMeetingsServiceServer() {}

// UnsafeMeetingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeetingsServiceServer will
// result in compilation errors.
type UnsafeMeetingsServiceServer interface {
	mustEmbedUnimplementedMeetingsServiceServer()
}

func RegisterMeetingsServiceServer(s grpc.ServiceRegistrar, srv MeetingsServiceServer) {
	s.RegisterService(&MeetingsService_ServiceDesc, srv)
}

func _MeetingsService_ArrangeMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meeting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServiceServer).ArrangeMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetings.MeetingsService/ArrangeMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServiceServer).ArrangeMeeting(ctx, req.(*Meeting))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingsService_CancelMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelMeetingParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServiceServer).CancelMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetings.MeetingsService/CancelMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServiceServer).CancelMeeting(ctx, req.(*CancelMeetingParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingsService_GetAvailableTimeForMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableTimeForMeetingParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServiceServer).GetAvailableTimeForMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetings.MeetingsService/GetAvailableTimeForMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServiceServer).GetAvailableTimeForMeeting(ctx, req.(*GetAvailableTimeForMeetingParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingsService_GetMeetingsByEstateID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingsByEstateIDParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServiceServer).GetMeetingsByEstateID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetings.MeetingsService/GetMeetingsByEstateID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServiceServer).GetMeetingsByEstateID(ctx, req.(*GetMeetingsByEstateIDParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingsService_GetMeetingsByPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingsByPhoneNumberParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServiceServer).GetMeetingsByPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetings.MeetingsService/GetMeetingsByPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServiceServer).GetMeetingsByPhoneNumber(ctx, req.(*GetMeetingsByPhoneNumberParameter))
	}
	return interceptor(ctx, in, info, handler)
}

// MeetingsService_ServiceDesc is the grpc.ServiceDesc for MeetingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeetingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meetings.MeetingsService",
	HandlerType: (*MeetingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ArrangeMeeting",
			Handler:    _MeetingsService_ArrangeMeeting_Handler,
		},
		{
			MethodName: "CancelMeeting",
			Handler:    _MeetingsService_CancelMeeting_Handler,
		},
		{
			MethodName: "GetAvailableTimeForMeeting",
			Handler:    _MeetingsService_GetAvailableTimeForMeeting_Handler,
		},
		{
			MethodName: "GetMeetingsByEstateID",
			Handler:    _MeetingsService_GetMeetingsByEstateID_Handler,
		},
		{
			MethodName: "GetMeetingsByPhoneNumber",
			Handler:    _MeetingsService_GetMeetingsByPhoneNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/meetings.proto",
}
