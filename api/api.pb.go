// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	User
	UserId
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Utype   string `protobuf:"bytes,3,opt,name=utype" json:"utype,omitempty"`
	Balance int32  `protobuf:"varint,4,opt,name=balance" json:"balance,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetUtype() string {
	if m != nil {
		return m.Utype
	}
	return ""
}

func (m *User) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type UserId struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserId) Reset()                    { *m = UserId{} }
func (m *UserId) String() string            { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()               {}
func (*UserId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*UserId)(nil), "api.UserId")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DoUsers service

type DoUsersClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserId, error)
	SelectUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error)
	ListUsers(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type doUsersClient struct {
	cc *grpc.ClientConn
}

func NewDoUsersClient(cc *grpc.ClientConn) DoUsersClient {
	return &doUsersClient{cc}
}

func (c *doUsersClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := grpc.Invoke(ctx, "/api.DoUsers/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doUsersClient) SelectUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/api.DoUsers/SelectUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doUsersClient) ListUsers(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/api.DoUsers/ListUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doUsersClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/api.DoUsers/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DoUsers service

type DoUsersServer interface {
	CreateUser(context.Context, *User) (*UserId, error)
	SelectUser(context.Context, *UserId) (*User, error)
	ListUsers(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
}

func RegisterDoUsersServer(s *grpc.Server, srv DoUsersServer) {
	s.RegisterService(&_DoUsers_serviceDesc, srv)
}

func _DoUsers_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoUsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DoUsers/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoUsersServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoUsers_SelectUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoUsersServer).SelectUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DoUsers/SelectUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoUsersServer).SelectUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoUsers_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoUsersServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DoUsers/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoUsersServer).ListUsers(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoUsers_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoUsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DoUsers/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoUsersServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _DoUsers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DoUsers",
	HandlerType: (*DoUsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _DoUsers_CreateUser_Handler,
		},
		{
			MethodName: "SelectUser",
			Handler:    _DoUsers_SelectUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _DoUsers_ListUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _DoUsers_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x9a, 0xb6, 0x64, 0x94, 0x20, 0x43, 0x91, 0x25, 0x78, 0x28, 0x7b, 0x2a, 0x1e,
	0x1a, 0x50, 0xbc, 0x14, 0x6f, 0x7a, 0x29, 0x78, 0xaa, 0xf4, 0xe2, 0x6d, 0xdb, 0x1d, 0xca, 0x42,
	0xcc, 0x2e, 0xd9, 0xed, 0x41, 0xc4, 0x8b, 0xaf, 0xe0, 0xa3, 0xf9, 0x0a, 0x3e, 0x83, 0x67, 0xc9,
	0x04, 0x63, 0xa8, 0xbd, 0xfd, 0xb3, 0xbb, 0xdf, 0xff, 0xff, 0x3b, 0x90, 0x2a, 0x67, 0xe6, 0xae,
	0xb6, 0xc1, 0xe2, 0x40, 0x39, 0x93, 0x5f, 0xec, 0xac, 0xdd, 0x95, 0x54, 0x28, 0x67, 0x0a, 0x55,
	0x55, 0x36, 0xa8, 0x60, 0x6c, 0xe5, 0xdb, 0x27, 0xf2, 0x09, 0x92, 0xb5, 0xa7, 0x1a, 0x33, 0x88,
	0x8d, 0x16, 0xd1, 0x34, 0x9a, 0x0d, 0x57, 0xb1, 0xd1, 0x88, 0x90, 0x54, 0xea, 0x99, 0x44, 0x3c,
	0x8d, 0x66, 0xe9, 0x8a, 0x35, 0x4e, 0x60, 0xb8, 0x0f, 0x2f, 0x8e, 0xc4, 0x80, 0x0f, 0xdb, 0x01,
	0x05, 0x8c, 0x37, 0xaa, 0x54, 0xd5, 0x96, 0x44, 0xc2, 0xf8, 0xef, 0x28, 0x05, 0x8c, 0x1a, 0xef,
	0xa5, 0x3e, 0x74, 0xbf, 0xfa, 0x8e, 0x60, 0x7c, 0x6f, 0x9b, 0x4b, 0x8f, 0x0b, 0x80, 0xbb, 0x9a,
	0x54, 0x20, 0xee, 0x91, 0xce, 0x9b, 0xfa, 0x8d, 0xcc, 0x4f, 0x3a, 0xb9, 0xd4, 0x72, 0xf2, 0xfe,
	0xf9, 0xf5, 0x11, 0x67, 0x32, 0xe5, 0x7f, 0xec, 0x3d, 0xd5, 0x8b, 0xe8, 0x12, 0x6f, 0x01, 0x1e,
	0xa9, 0xa4, 0x6d, 0x60, 0xb6, 0x0f, 0xe4, 0x7f, 0x46, 0xf2, 0x9c, 0xd9, 0x33, 0xcc, 0x3a, 0xb6,
	0x78, 0x35, 0xfa, 0x0d, 0x6f, 0x20, 0x7d, 0x30, 0x3e, 0xb4, 0x35, 0x7a, 0xc1, 0x3d, 0x14, 0x19,
	0x3d, 0x45, 0xe8, 0x50, 0x2e, 0xbc, 0x76, 0xfa, 0x48, 0xe1, 0xff, 0x91, 0xf2, 0x20, 0x72, 0x33,
	0xe2, 0xad, 0x5f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x52, 0x55, 0x6b, 0xa5, 0x01, 0x00,
	0x00,
}
