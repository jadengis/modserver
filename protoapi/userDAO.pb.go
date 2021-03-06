// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userDAO.proto

/*
Package protoapi is a generated protocol buffer package.

It is generated from these files:
	userDAO.proto

It has these top-level messages:
	Id
	User
	SaveResponse
	UpdateResponse
	UserResponse
	DeleteResponse
*/
package protoapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type Id struct {
	Value uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Id) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type User struct {
	Id        *Id                        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	DeletedAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt" json:"deleted_at,omitempty"`
	Age       int64                      `protobuf:"varint,5,opt,name=age" json:"age,omitempty"`
	Name      string                     `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	Email     string                     `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetId() *Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *User) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetDeletedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SaveResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *SaveResponse) Reset()                    { *m = SaveResponse{} }
func (m *SaveResponse) String() string            { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()               {}
func (*SaveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SaveResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Id)(nil), "protoapi.Id")
	proto.RegisterType((*User)(nil), "protoapi.User")
	proto.RegisterType((*SaveResponse)(nil), "protoapi.SaveResponse")
	proto.RegisterType((*UpdateResponse)(nil), "protoapi.UpdateResponse")
	proto.RegisterType((*UserResponse)(nil), "protoapi.UserResponse")
	proto.RegisterType((*DeleteResponse)(nil), "protoapi.DeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserDAO service

type UserDAOClient interface {
	SaveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*SaveResponse, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type userDAOClient struct {
	cc *grpc.ClientConn
}

func NewUserDAOClient(cc *grpc.ClientConn) UserDAOClient {
	return &userDAOClient{cc}
}

func (c *userDAOClient) SaveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := grpc.Invoke(ctx, "/protoapi.UserDAO/SaveUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDAOClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/protoapi.UserDAO/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDAOClient) GetUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/protoapi.UserDAO/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDAOClient) DeleteUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/protoapi.UserDAO/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserDAO service

type UserDAOServer interface {
	SaveUser(context.Context, *User) (*SaveResponse, error)
	UpdateUser(context.Context, *User) (*UpdateResponse, error)
	GetUser(context.Context, *Id) (*UserResponse, error)
	DeleteUser(context.Context, *Id) (*DeleteResponse, error)
}

func RegisterUserDAOServer(s *grpc.Server, srv UserDAOServer) {
	s.RegisterService(&_UserDAO_serviceDesc, srv)
}

func _UserDAO_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDAOServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoapi.UserDAO/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDAOServer).SaveUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDAO_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDAOServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoapi.UserDAO/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDAOServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDAO_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDAOServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoapi.UserDAO/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDAOServer).GetUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDAO_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDAOServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoapi.UserDAO/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDAOServer).DeleteUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserDAO_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoapi.UserDAO",
	HandlerType: (*UserDAOServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _UserDAO_SaveUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserDAO_UpdateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserDAO_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserDAO_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userDAO.proto",
}

func init() { proto.RegisterFile("userDAO.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0x6a, 0xea, 0x40,
	0x10, 0xc6, 0x93, 0x18, 0xff, 0xcd, 0xf1, 0x88, 0x2c, 0x87, 0x43, 0x08, 0x85, 0xca, 0x5e, 0x79,
	0x15, 0x69, 0x2a, 0xa5, 0xbd, 0x14, 0x84, 0xe2, 0x55, 0x61, 0x5b, 0xaf, 0xcb, 0xda, 0x9d, 0x4a,
	0x20, 0x31, 0x21, 0xd9, 0xf8, 0x1e, 0x7d, 0xc7, 0x3e, 0x48, 0xc9, 0x6c, 0x82, 0xb1, 0x08, 0xe2,
	0x95, 0x3b, 0x3b, 0xf3, 0xf3, 0x9b, 0xef, 0xcb, 0xc2, 0xdf, 0xb2, 0xc0, 0x7c, 0xb5, 0x7c, 0x09,
	0xb2, 0x3c, 0xd5, 0x29, 0x1b, 0xd0, 0x8f, 0xcc, 0x22, 0xff, 0x76, 0x97, 0xa6, 0xbb, 0x18, 0xe7,
	0x74, 0xb1, 0x2d, 0x3f, 0xe7, 0x3a, 0x4a, 0xb0, 0xd0, 0x32, 0xc9, 0xcc, 0x28, 0xf7, 0xc1, 0x59,
	0x2b, 0xf6, 0x0f, 0xba, 0x07, 0x19, 0x97, 0xe8, 0xd9, 0x53, 0x7b, 0xe6, 0x0a, 0x53, 0xf0, 0x2f,
	0x07, 0xdc, 0x4d, 0x81, 0x39, 0xbb, 0x01, 0x27, 0x52, 0xd4, 0xfb, 0x13, 0x8e, 0x82, 0xe6, 0xcf,
	0x83, 0xb5, 0x12, 0x4e, 0xa4, 0xd8, 0x13, 0xc0, 0x47, 0x8e, 0x52, 0xa3, 0x7a, 0x97, 0xda, 0x73,
	0x68, 0xca, 0x0f, 0x8c, 0x70, 0xd0, 0x08, 0x07, 0x6f, 0x8d, 0xb0, 0x18, 0xd6, 0xd3, 0x4b, 0x5d,
	0xa1, 0x65, 0xa6, 0x1a, 0xb4, 0x73, 0x19, 0xad, 0xa7, 0x0d, 0xaa, 0x30, 0xc6, 0x1a, 0x75, 0x2f,
	0xa3, 0xf5, 0xf4, 0x52, 0xb3, 0x09, 0x74, 0xe4, 0x0e, 0xbd, 0xee, 0xd4, 0x9e, 0x75, 0x44, 0x75,
	0x64, 0x0c, 0xdc, 0xbd, 0x4c, 0xd0, 0xeb, 0x4d, 0xed, 0xd9, 0x50, 0xd0, 0xb9, 0xca, 0x04, 0x13,
	0x19, 0xc5, 0x5e, 0x9f, 0x2e, 0x4d, 0xc1, 0x43, 0x18, 0xbd, 0xca, 0x03, 0x0a, 0x2c, 0xb2, 0x74,
	0x5f, 0x20, 0xe3, 0xe0, 0x56, 0xd9, 0xd7, 0xe1, 0x8c, 0x8f, 0xe1, 0x54, 0xc1, 0x09, 0xea, 0xf1,
	0x05, 0x8c, 0x37, 0xb4, 0xf7, 0x55, 0x54, 0x08, 0x23, 0xaa, 0xae, 0x61, 0x26, 0x30, 0x5e, 0x91,
	0xcd, 0x86, 0x0a, 0xbf, 0x6d, 0xe8, 0x6f, 0xcc, 0xe3, 0x60, 0x0b, 0x18, 0x54, 0xbb, 0xd3, 0x27,
	0xfd, 0xc5, 0xfb, 0xff, 0x8f, 0x75, 0xdb, 0x1f, 0xb7, 0xd8, 0x23, 0x80, 0xd9, 0xfe, 0x2c, 0xe7,
	0xb5, 0xea, 0x13, 0x8f, 0xdc, 0x62, 0x77, 0xd0, 0x7f, 0x46, 0x4d, 0xd8, 0xc9, 0xab, 0x69, 0x8b,
	0xb5, 0x2d, 0x72, 0x8b, 0x3d, 0x00, 0x18, 0x03, 0x67, 0xa8, 0x96, 0xd4, 0xa9, 0x49, 0x6e, 0x6d,
	0x7b, 0xd4, 0xba, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x91, 0x90, 0x6d, 0x09, 0x03, 0x00,
	0x00,
}
