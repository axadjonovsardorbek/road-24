// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: auth.proto

package auth

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
	UserService_Register_FullMethodName              = "/auth.UserService/Register"
	UserService_Login_FullMethodName                 = "/auth.UserService/Login"
	UserService_LoginDriver_FullMethodName           = "/auth.UserService/LoginDriver"
	UserService_Profile_FullMethodName               = "/auth.UserService/Profile"
	UserService_DeleteProfile_FullMethodName         = "/auth.UserService/DeleteProfile"
	UserService_GetAllUsers_FullMethodName           = "/auth.UserService/GetAllUsers"
	UserService_RefreshToken_FullMethodName          = "/auth.UserService/RefreshToken"
	UserService_RefreshTokenForDriver_FullMethodName = "/auth.UserService/RefreshTokenForDriver"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*Void, error)
	Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*TokenRes, error)
	LoginDriver(ctx context.Context, in *DriverLoginReq, opts ...grpc.CallOption) (*TokenRes, error)
	Profile(ctx context.Context, in *ById, opts ...grpc.CallOption) (*UserRes, error)
	DeleteProfile(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersReq, opts ...grpc.CallOption) (*GetAllUsersRes, error)
	RefreshToken(ctx context.Context, in *ById, opts ...grpc.CallOption) (*TokenRes, error)
	RefreshTokenForDriver(ctx context.Context, in *ById, opts ...grpc.CallOption) (*TokenRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*TokenRes, error) {
	out := new(TokenRes)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginDriver(ctx context.Context, in *DriverLoginReq, opts ...grpc.CallOption) (*TokenRes, error) {
	out := new(TokenRes)
	err := c.cc.Invoke(ctx, UserService_LoginDriver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Profile(ctx context.Context, in *ById, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, UserService_Profile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteProfile(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_DeleteProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersReq, opts ...grpc.CallOption) (*GetAllUsersRes, error) {
	out := new(GetAllUsersRes)
	err := c.cc.Invoke(ctx, UserService_GetAllUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RefreshToken(ctx context.Context, in *ById, opts ...grpc.CallOption) (*TokenRes, error) {
	out := new(TokenRes)
	err := c.cc.Invoke(ctx, UserService_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RefreshTokenForDriver(ctx context.Context, in *ById, opts ...grpc.CallOption) (*TokenRes, error) {
	out := new(TokenRes)
	err := c.cc.Invoke(ctx, UserService_RefreshTokenForDriver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *UserCreateReq) (*Void, error)
	Login(context.Context, *UserLoginReq) (*TokenRes, error)
	LoginDriver(context.Context, *DriverLoginReq) (*TokenRes, error)
	Profile(context.Context, *ById) (*UserRes, error)
	DeleteProfile(context.Context, *ById) (*Void, error)
	GetAllUsers(context.Context, *GetAllUsersReq) (*GetAllUsersRes, error)
	RefreshToken(context.Context, *ById) (*TokenRes, error)
	RefreshTokenForDriver(context.Context, *ById) (*TokenRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *UserCreateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *UserLoginReq) (*TokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) LoginDriver(context.Context, *DriverLoginReq) (*TokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginDriver not implemented")
}
func (UnimplementedUserServiceServer) Profile(context.Context, *ById) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (UnimplementedUserServiceServer) DeleteProfile(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *GetAllUsersReq) (*GetAllUsersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) RefreshToken(context.Context, *ById) (*TokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedUserServiceServer) RefreshTokenForDriver(context.Context, *ById) (*TokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshTokenForDriver not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*UserCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_LoginDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginDriver(ctx, req.(*DriverLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Profile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Profile(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteProfile(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*GetAllUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RefreshToken(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RefreshTokenForDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RefreshTokenForDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RefreshTokenForDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RefreshTokenForDriver(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "LoginDriver",
			Handler:    _UserService_LoginDriver_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _UserService_Profile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _UserService_DeleteProfile_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _UserService_RefreshToken_Handler,
		},
		{
			MethodName: "RefreshTokenForDriver",
			Handler:    _UserService_RefreshTokenForDriver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
