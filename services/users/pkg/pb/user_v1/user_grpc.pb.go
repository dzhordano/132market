// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: user.proto

package user_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserServiceV1_CreateUser_FullMethodName      = "/api.UserServiceV1/CreateUser"
	UserServiceV1_UpdateUser_FullMethodName      = "/api.UserServiceV1/UpdateUser"
	UserServiceV1_FindUserById_FullMethodName    = "/api.UserServiceV1/FindUserById"
	UserServiceV1_FindUserByEmail_FullMethodName = "/api.UserServiceV1/FindUserByEmail"
	UserServiceV1_ListUsers_FullMethodName       = "/api.UserServiceV1/ListUsers"
	UserServiceV1_DeleteUser_FullMethodName      = "/api.UserServiceV1/DeleteUser"
	UserServiceV1_SearchUsers_FullMethodName     = "/api.UserServiceV1/SearchUsers"
	UserServiceV1_SetUserState_FullMethodName    = "/api.UserServiceV1/SetUserState"
	UserServiceV1_UpdateLastSeen_FullMethodName  = "/api.UserServiceV1/UpdateLastSeen"
	UserServiceV1_CheckUserExists_FullMethodName = "/api.UserServiceV1/CheckUserExists"
)

// UserServiceV1Client is the client API for UserServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceV1Client interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	FindUserById(ctx context.Context, in *FindUserByIdRequest, opts ...grpc.CallOption) (*FindUserByIdResponse, error)
	FindUserByEmail(ctx context.Context, in *FindUserByEmailRequest, opts ...grpc.CallOption) (*FindUserByEmailResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Methods offered by ai, valid? I guess useful...
	SearchUsers(ctx context.Context, in *SearchUsersRequest, opts ...grpc.CallOption) (*SearchUsersResponse, error)
	SetUserState(ctx context.Context, in *SetUserStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateLastSeen(ctx context.Context, in *UpdateLastSeenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckUserExists(ctx context.Context, in *CheckUserExistsRequest, opts ...grpc.CallOption) (*CheckUserExistsResponse, error)
}

type userServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceV1Client(cc grpc.ClientConnInterface) UserServiceV1Client {
	return &userServiceV1Client{cc}
}

func (c *userServiceV1Client) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserServiceV1_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserServiceV1_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) FindUserById(ctx context.Context, in *FindUserByIdRequest, opts ...grpc.CallOption) (*FindUserByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindUserByIdResponse)
	err := c.cc.Invoke(ctx, UserServiceV1_FindUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) FindUserByEmail(ctx context.Context, in *FindUserByEmailRequest, opts ...grpc.CallOption) (*FindUserByEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindUserByEmailResponse)
	err := c.cc.Invoke(ctx, UserServiceV1_FindUserByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, UserServiceV1_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserServiceV1_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) SearchUsers(ctx context.Context, in *SearchUsersRequest, opts ...grpc.CallOption) (*SearchUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUsersResponse)
	err := c.cc.Invoke(ctx, UserServiceV1_SearchUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) SetUserState(ctx context.Context, in *SetUserStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserServiceV1_SetUserState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) UpdateLastSeen(ctx context.Context, in *UpdateLastSeenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserServiceV1_UpdateLastSeen_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) CheckUserExists(ctx context.Context, in *CheckUserExistsRequest, opts ...grpc.CallOption) (*CheckUserExistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckUserExistsResponse)
	err := c.cc.Invoke(ctx, UserServiceV1_CheckUserExists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceV1Server is the server API for UserServiceV1 service.
// All implementations must embed UnimplementedUserServiceV1Server
// for forward compatibility.
type UserServiceV1Server interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	FindUserById(context.Context, *FindUserByIdRequest) (*FindUserByIdResponse, error)
	FindUserByEmail(context.Context, *FindUserByEmailRequest) (*FindUserByEmailResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	// Methods offered by ai, valid? I guess useful...
	SearchUsers(context.Context, *SearchUsersRequest) (*SearchUsersResponse, error)
	SetUserState(context.Context, *SetUserStateRequest) (*emptypb.Empty, error)
	UpdateLastSeen(context.Context, *UpdateLastSeenRequest) (*emptypb.Empty, error)
	CheckUserExists(context.Context, *CheckUserExistsRequest) (*CheckUserExistsResponse, error)
	mustEmbedUnimplementedUserServiceV1Server()
}

// UnimplementedUserServiceV1Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceV1Server struct{}

func (UnimplementedUserServiceV1Server) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceV1Server) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceV1Server) FindUserById(context.Context, *FindUserByIdRequest) (*FindUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserById not implemented")
}
func (UnimplementedUserServiceV1Server) FindUserByEmail(context.Context, *FindUserByEmailRequest) (*FindUserByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByEmail not implemented")
}
func (UnimplementedUserServiceV1Server) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUserServiceV1Server) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceV1Server) SearchUsers(context.Context, *SearchUsersRequest) (*SearchUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedUserServiceV1Server) SetUserState(context.Context, *SetUserStateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserState not implemented")
}
func (UnimplementedUserServiceV1Server) UpdateLastSeen(context.Context, *UpdateLastSeenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLastSeen not implemented")
}
func (UnimplementedUserServiceV1Server) CheckUserExists(context.Context, *CheckUserExistsRequest) (*CheckUserExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserExists not implemented")
}
func (UnimplementedUserServiceV1Server) mustEmbedUnimplementedUserServiceV1Server() {}
func (UnimplementedUserServiceV1Server) testEmbeddedByValue()                       {}

// UnsafeUserServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceV1Server will
// result in compilation errors.
type UnsafeUserServiceV1Server interface {
	mustEmbedUnimplementedUserServiceV1Server()
}

func RegisterUserServiceV1Server(s grpc.ServiceRegistrar, srv UserServiceV1Server) {
	// If the following call pancis, it indicates UnimplementedUserServiceV1Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserServiceV1_ServiceDesc, srv)
}

func _UserServiceV1_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_FindUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).FindUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_FindUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).FindUserById(ctx, req.(*FindUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_FindUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).FindUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_FindUserByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).FindUserByEmail(ctx, req.(*FindUserByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_SearchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).SearchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_SearchUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).SearchUsers(ctx, req.(*SearchUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_SetUserState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).SetUserState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_SetUserState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).SetUserState(ctx, req.(*SetUserStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_UpdateLastSeen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLastSeenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).UpdateLastSeen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_UpdateLastSeen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).UpdateLastSeen(ctx, req.(*UpdateLastSeenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_CheckUserExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).CheckUserExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServiceV1_CheckUserExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).CheckUserExists(ctx, req.(*CheckUserExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServiceV1_ServiceDesc is the grpc.ServiceDesc for UserServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserServiceV1",
	HandlerType: (*UserServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserServiceV1_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserServiceV1_UpdateUser_Handler,
		},
		{
			MethodName: "FindUserById",
			Handler:    _UserServiceV1_FindUserById_Handler,
		},
		{
			MethodName: "FindUserByEmail",
			Handler:    _UserServiceV1_FindUserByEmail_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserServiceV1_ListUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserServiceV1_DeleteUser_Handler,
		},
		{
			MethodName: "SearchUsers",
			Handler:    _UserServiceV1_SearchUsers_Handler,
		},
		{
			MethodName: "SetUserState",
			Handler:    _UserServiceV1_SetUserState_Handler,
		},
		{
			MethodName: "UpdateLastSeen",
			Handler:    _UserServiceV1_UpdateLastSeen_Handler,
		},
		{
			MethodName: "CheckUserExists",
			Handler:    _UserServiceV1_CheckUserExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
